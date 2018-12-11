package main

import (
	"encoding/json"
	"fmt" // пакет для форматированного ввода вывода
	"io/ioutil"
	"net"
	"net/http" // пакет для поддержки HTTP протокола
	"os"

	"urok/webserver/uplaodfile/uploader"

	"gopkg.in/sungora/app.v1/lg"
)

type content struct {
	Code    int
	Message string
	Error   bool
	Data    interface{} `json:"Data,omitempty"`
}

var (
	chanelAppControl = make(chan os.Signal, 1) // Канал управления и остановкой приложения
	dir              string
)

func main() {
	defer func() {
		chanelAppControl <- os.Interrupt
	}()
	var (
		err   error
		store net.Listener
	)
	dir, _ = os.Getwd()
	dir += "/webserver/uplaodfile"

	// Модуль загрузки файлов и получение их по идентификатору
	if err = uploader.Init(dir, 30); err != nil {
		fmt.Println(err.Error())
		return
	}

	http.HandleFunc("/", RouterHandler)
	store, err = net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go http.Serve(store, nil)
	defer store.Close()

	fmt.Println("SERVER START http://localhost:9090")
	<-chanelAppControl
	return

}

// Wait an application
func Wait() {
	chanelAppControl <- os.Interrupt
	<-chanelAppControl
}

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/upload" {
		UploadRouterHandler(w, r)
	} else {
		HomeRouterHandler(w, r)
	}
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	var data, _ = ioutil.ReadFile(dir + string(os.PathSeparator) + "test.html")
	// размер и тип контента
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	// Статус ответа
	w.WriteHeader(200)
	// Тело документа
	w.Write(data)
}

func UploadRouterHandler(w http.ResponseWriter, r *http.Request) {
	key, err := uploader.Upload(r, `files[]`)
	if err != nil {
		key = err.Error()
	}
	res := new(content)
	res.Code = 345
	res.Message = "Все хорошо"
	res.Error = false
	res.Data = key
	data, err := json.Marshal(res)
	if err != nil {
		lg.Error(err.Error())
		return
	}
	// размер и тип контента
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	// Статус ответа
	w.WriteHeader(200)
	// Тело документа
	w.Write(data)
	fmt.Println("key: "+key, r.Method, r.URL)
}
