package handlers

import (
	"fmt"
	"net/http"
	"plugin"

	"bou.ke/monkey"
)

func Example(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintf(rw, "Hello, %s!", req.Form.Get("name"))
}

func Reload(rw http.ResponseWriter, req *http.Request) {
	p, _ := plugin.Open("handlers.so")
	sym, _ := p.Lookup("Example")
	monkey.Patch(Example, sym)
}
