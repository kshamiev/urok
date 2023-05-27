package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

// http://localhost:8010/create
// {"id":"0","username":"Вася Пупкин","email":"Мыло","age":20}

// http://localhost:8010/get/23
// {"id":"23","username":"Вася Пупкин","email":"Мыло","friends":null,"age":20}

type Service struct {
	store []string
	mu    sync.Mutex
	pos   int
}

func NewService() *Service {
	return &Service{
		store: []string{
			"http://localhost:8010",
			"http://localhost:8020",
			"http://localhost:8030",
		},
	}
}

func (s *Service) iterationServer() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pos++
	if s.pos == len(s.store) {
		s.pos = 0
	}
	// здесь можно реализовать проверку, что сервер в рабочем состоянии
	return s.store[s.pos]
}

func (s *Service) Proxy(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(r.Method, s.iterationServer()+r.URL.String(), r.Body)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	for key := range resp.Header {
		w.Header().Set(key, resp.Header.Get(key))
	}
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	s, err := NewHTTPServer(&HttpServerConfig{
		Proto:          "http",
		Host:           "localhost",
		Port:           8080,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		IdleTimeout:    time.Minute,
		MaxHeaderBytes: 1048576,
	}, Routing())
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	Unlock(ch)
	// }()
	Lock(ch)

	err = s.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// ROUTING

func Routing() *chi.Mux {
	router := chi.NewRouter()
	service := NewService()

	router.HandleFunc("/*", service.Proxy)

	return router
}

// HTTP SERVER

type HttpServerConfig struct {
	Domain         string        `yaml:"domain"`         // External Host:Port (from swagger)
	Proto          string        `yaml:"proto"`          // Server Proto
	Host           string        `yaml:"host"`           // Server Host
	Port           int           `yaml:"port"`           // Server Port
	ReadTimeout    time.Duration `yaml:"readTimeout"`    // Время ожидания web запроса в секундах
	WriteTimeout   time.Duration `yaml:"writeTimeout"`   // Время ожидания окончания передачи ответа в секундах
	RequestTimeout time.Duration `yaml:"requestTimeout"` // Время ожидания окончания выполнения запроса
	IdleTimeout    time.Duration `yaml:"idleTimeout"`    // Время ожидания следующего запроса
	MaxHeaderBytes int           `yaml:"maxHeaderBytes"` // Максимальный размер заголовка получаемого от браузера клиента в байтах
	PKeyFile       string        `yaml:"pKeyFile"`       // Файл содержащий приватный ключ
	CertFile       string        `yaml:"certFile"`       // Файл содержащий сертификат
}

type HttpServer struct {
	server    *http.Server  // Сервер HTTP
	chControl chan struct{} // Управление ожиданием завершения работы сервера
	lis       net.Listener
}

func NewHTTPServer(cfg *HttpServerConfig, mux http.Handler) (comp *HttpServer, err error) {
	comp = &HttpServer{
		server: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:        mux,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			IdleTimeout:    cfg.IdleTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
		chControl: make(chan struct{}),
	}

	if comp.lis, err = net.Listen("tcp", comp.server.Addr); err != nil {
		return
	}

	go func() {
		if cfg.Proto == "https" && cfg.CertFile != "" && cfg.PKeyFile != "" {
			_ = comp.server.ServeTLS(comp.lis, cfg.CertFile, cfg.PKeyFile)
		} else {
			_ = comp.server.Serve(comp.lis)
		}
		close(comp.chControl)
	}()

	return comp, nil
}

func (comp *HttpServer) Close() error {
	if comp.lis == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := comp.server.Shutdown(ctx); err != nil {
		if err := comp.lis.Close(); err != nil {
			return err
		}
	}

	<-comp.chControl
	return nil
}

// APP LOCK AND UNLOCK

// Lock run application
func Lock(ch chan os.Signal) {
	defer func() {
		ch <- os.Interrupt
	}()
	// The correctness of the application is closed by a signal
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-ch
}

// Unlock run application
func Unlock(ch chan os.Signal) {
	ch <- os.Interrupt
	<-ch
}
