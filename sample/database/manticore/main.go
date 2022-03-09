// install from ubuntu
//
// wget https://repo.manticoresearch.com/manticore-repo.noarch.deb
// sudo dpkg -i manticore-repo.noarch.deb
// sudo apt update
// sudo apt install manticore manticore-columnar-lib
//
// sudo indexer --all --rotate
//
// systemctl status manticore
// systemctl restart manticore
//
// sudo journalctl --unit manticore
// sudo journalctl -xe
//
// Дмитрий Свиридов @dimuska139
package main

import (
	"fmt"

	"github.com/manticoresoftware/go-sdk/manticore"
)

func main() {
	cl := manticore.NewClient()
	cl.SetServer("127.0.0.1", 9312)
	if _, err := cl.Open(); err != nil {
		fmt.Println(err)
		return
	}

	res, err := cl.Sphinxql(`replace into usersidxrt values(1, 'my subject', 'my content')`)
	fmt.Println(res, err)
	res, err = cl.Sphinxql(`replace into usersidxrt values(2,'another subject', 'more content')`)
	fmt.Println(res, err)
	res, err = cl.Sphinxql(`replace into usersidxrt values(5,'again subject', 'one more content')`)
	fmt.Println(res, err)
	res2, err2 := cl.Query("more|another", "usersidxrt")
	fmt.Println(res2, err2)

	res2, err2 = cl.Query("YAqcEZM8", "usersidx")
	fmt.Println(res2, err2)
}
