package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s := &HttpServer{}
	s.start()
	lock(nil)
	s.stop()
}

type HttpServer struct {
	server    *http.Server  // Сервер HTTP
	chControl chan struct{} // Управление ожиданием завершения работы сервера
	lis       net.Listener
}

func (rec *HttpServer) start() {
	var err error
	mux := http.NewServeMux()
	mux.HandleFunc("GET /task/{id}/", home)
	rec.server = &http.Server{
		Addr:           "localhost:8080",
		Handler:        mux,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		IdleTimeout:    time.Minute,
		MaxHeaderBytes: 1048576,
	}
	if rec.lis, err = net.Listen("tcp", rec.server.Addr); err != nil {
		log.Fatal(err)
	}
	rec.chControl = make(chan struct{})

	go func() {
		_ = rec.server.Serve(rec.lis)
		close(rec.chControl)
	}()
}

func (rec *HttpServer) stop() {
	if rec.lis == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := rec.server.Shutdown(ctx); err != nil {
		if err = rec.lis.Close(); err != nil {
			return
		}
	}
	<-rec.chControl
	return
}

func home(w http.ResponseWriter, r *http.Request) {
	str := "hello world: " + r.PathValue("id")
	fmt.Println(str)
	_, _ = w.Write([]byte(str))
	w.WriteHeader(http.StatusOK)
}

// lock run application
func lock(ch chan os.Signal) {
	if ch == nil {
		ch = make(chan os.Signal, 1)
	}
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
