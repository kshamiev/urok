package webserver

import (
	"fmt"
	"net/http"
	"time"
	"net"
)

type serverHTTP struct {
	// Тип сервера
	Type string
	// Режим работы приложения, каким способом приложение слушает запросы. Возможные значения: socket, tcp
	Mode string
	// Адрес занимаемый приложением, можно указывать как доменное имя так и ip адрес
	Host string
	// Порт занимаемый приложение в режиме работы Mode=tcp
	Port int64
	// Unix:Socket занимаемый приложением, применимо только для *nix и mac os x
	Socket string
	// Домены (через запятую) обслуживаемые данным сервером
	Domain string
	// Защита от атак. Время ожидания web запроса в секундах, по истечении которого соединение сбрасывается
	ReadTimeout int64
	// Защита от атак. Время ожидания окончания передачи ответа в секундах, по истечении которого соединение сбрасывается
	WriteTimeout int64
	// Защита от атак. Максимальный размер заголовка получаемого от браузера клиента
	MaxHeaderBytes int64
	// Ллогирование работы сервер (зарезервировано, не используеться)
	//Logs bool
}

func Sample1() {
	fmt.Println("start")
	s := newServer()
	http.ListenAndServe("localhost:8081", s)
	fmt.Println("stop")
}

func Sample2() {
	fmt.Println("start")

	var store net.Listener
	var err error

	Server := &http.Server{
		Addr: "localhost:8081",
		//Addr:           "195.161.115.49:8081",
		Handler:        newServer(),
		ReadTimeout:    time.Second * time.Duration(300),
		WriteTimeout:   time.Second * time.Duration(300),
		MaxHeaderBytes: 1048576,
	}
	for i := 5; i > 0; i-- {
		store, err = net.Listen("tcp", Server.Addr)
		time.Sleep(time.Millisecond * 100)
		if err == nil {
			break
		}
		fmt.Println("connected retry ", Server.Addr)
	}
	if err == nil && store != nil {
		fmt.Println("connect YES")
		go Server.Serve(store)
		//time.Sleep(time.Second * 30)
		fmt.Scanln()
		store.Close()
	} else {
		fmt.Println("connect NOT")
	}

	fmt.Println("stop")
}

// Создание УП. Создается пакетом роутера.
func newServer() *serverHTTP {
	var self = new(serverHTTP)
	return self
}

// ServeHTTP Точка входа запроса (в приложение).
func (self *serverHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response(w)
}

func response(Writer http.ResponseWriter) {
	// Тип и Кодировка документа
	Writer.Header().Set("Content-Type", "text/html; charset=utf-8")

	content := []byte("popcorn")
	var err error
	var loc *time.Location

	if loc, err = time.LoadLocation(`Europe/Moscow`); err != nil {
		loc = time.UTC
	}
	t := time.Now().In(loc)
	d := t.Format(time.RFC1123)

	// запрет кеширования
	Writer.Header().Set("Cache-Control", "no-cache, must-revalidate")
	Writer.Header().Set("Pragma", "no-cache")
	Writer.Header().Set("Date", d)
	Writer.Header().Set("Last-Modified", d)
	// размер контента
	Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
	// Статус ответа
	Writer.WriteHeader(200)
	// Тело документа
	Writer.Write(content)
	//
	return
}
