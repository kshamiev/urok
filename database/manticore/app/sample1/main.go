package main

import (
	"fmt"

	"github.com/manticoresoftware/go-sdk/manticore"
)

func main() {
	cl := manticore.NewClient()
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
