package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

// http://localhost:8010/create
// {"id":"0","username":"Вася Пупкин","email":"Мыло","age":20}

// http://localhost:8010/get/23
// {"id":"23","username":"Вася Пупкин","email":"Мыло","friends":null,"age":20}

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Friends  []string `json:"friends"`
	Age      int      `json:"age"`
}

type Service struct {
	address string
	store   map[string]User
	mutex   sync.Mutex
}

func NewService(address string) *Service {
	return &Service{
		address: address,
		store:   make(map[string]User),
	}
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("[" + r.Method + "] " + s.address + r.URL.String())
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	id := strconv.Itoa(len(s.store) + 1)
	user.ID = id
	s.store[id] = user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	enc := json.NewEncoder(w)
	if err := enc.Encode(map[string]string{"id": id}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("[" + r.Method + "] " + s.address + r.URL.String())
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	user, ok := s.store[strconv.Itoa(id)]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonUser)
}

func (s *Service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	log.Println("[" + r.Method + "] " + s.address + r.URL.String())
	var data struct {
		SourceID string `json:"source_id"`
		TargetID string `json:"target_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	source, ok := s.store[data.SourceID]
	if !ok {
		http.Error(w, "source user not found", http.StatusNotFound)
		return
	}
	target, ok := s.store[data.TargetID]
	if !ok {
		http.Error(w, "target user not found", http.StatusNotFound)
		return
	}

	source.Friends = append(source.Friends, target.ID)
	target.Friends = append(target.Friends, source.ID)

	_, _ = fmt.Fprintf(w, "%s and %s are now friends", source.Username, target.Username)
}

func (s *Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("[" + r.Method + "] " + s.address + r.URL.String())
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	user, ok := s.store[strconv.Itoa(id)]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(s.store, strconv.Itoa(id))
	for _, friend := range user.Friends {
		friendUser, ok := s.store[friend]
		if ok {
			for i, f := range friendUser.Friends {
				if f == user.Username {
					friendUser.Friends = append(friendUser.Friends[:i], friendUser.Friends[i+1:]...)
					break
				}
			}
			s.store[friend] = friendUser
		}
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("Deleted user %s", user.Username)))
}

func (s *Service) GetFriends(w http.ResponseWriter, r *http.Request) {
	log.Println("[" + r.Method + "] " + s.address + r.URL.String())
	s.mutex.Lock()
	defer s.mutex.Unlock()
	userID := chi.URLParam(r, "id")
	user, ok := s.store[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	friends := user.Friends

	jsonData, err := json.MarshalIndent(friends, "", "  ")
	if err != nil {
		http.Error(w, "Error converting data to JSON", http.StatusInternalServerError)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}

func (s *Service) UpdateAge(w http.ResponseWriter, r *http.Request) {
	log.Println("[" + r.Method + "] " + s.address + r.URL.String())
	s.mutex.Lock()
	defer s.mutex.Unlock()
	userID := chi.URLParam(r, "id")
	user, ok := s.store[userID]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	newAge, ok := data["new age"]
	if !ok {
		http.Error(w, "New age not found in request body", http.StatusBadRequest)
		return
	}
	age, err := strconv.Atoi(newAge)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	user.Age = age

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "The user's age has been updated")
}

func main() {
	host := "localhost"
	port := 8010
	s, err := NewHTTPServer(&HttpServerConfig{
		Proto:          "http",
		Host:           host,
		Port:           port,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		IdleTimeout:    time.Minute,
		MaxHeaderBytes: 1048576,
	}, Routing(host+":"+strconv.Itoa(port)))
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

func Routing(address string) *chi.Mux {
	router := chi.NewRouter()
	service := NewService(address)

	router.Post("/create", service.Create)
	router.Get("/get/{id}", service.Get)
	router.Post("/make/friends", service.MakeFriends)
	router.Delete("/delete/{id}", service.DeleteUser)
	router.Get("/friends/{id}", service.GetFriends)
	router.Put("/update/{id}", service.UpdateAge)

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
