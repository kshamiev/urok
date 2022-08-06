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
	// cl.SetServer("192.168.0.101", 9308)
	cl.SetServer("127.0.0.1", 9308)
	if _, err := cl.Open(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nREAL TIME INDEX\n")

	res, err := cl.Sphinxql(`replace into usersidxrt values(1005, 'my subject', 'my content')`)
	fmt.Println(res, err)
	res, err = cl.Sphinxql(`replace into usersidxrt values(1002,'another subject', 'more content')`)
	fmt.Println(res, err)
	res, err = cl.Sphinxql(`replace into usersidxrt values(1003,'again subject', 'one more content')`)
	fmt.Println(res, err)
	res, err = cl.Sphinxql(`replace into usersidxrt values(1004,'TEST RIDGEMONT', 'OPIS RIDGEMONT ONE')`)
	fmt.Println(res, err)
	res2, err2 := cl.Query("more|another", "usersidxrt")
	fmt.Println(res2, err2)

	fmt.Println("\nPLAIN INDEX\n")

	res2, err2 = cl.Query("RIDGEMONT", "usersidxrt usersidx")
	fmt.Println(res2, err2)
}
