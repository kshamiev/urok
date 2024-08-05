package main

import (
	"net/http"

	"github.com/kshamiev/urok/sample/plugin/handlers"
)

// статьи
// https://habr.com/ru/post/318896/
// https://pkg.go.dev/plugin@master

// Сначала нужно подготовить плагин
// go build -buildmode=plugin

// curl 'http://localhost:9999/example?name=Yuriy'
// curl 'http://localhost:9999/reload'

func main() {
	http.HandleFunc("/example", handlers.Example)
	http.HandleFunc("/reload", handlers.Reload)
	http.ListenAndServe(":9999", nil)
}
