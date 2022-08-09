package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/manticoresoftware/go-sdk/manticore"
)

func main() {
	cl := manticore.NewClient()
	cl.SetServer("localhost", 9312)
	if _, err := cl.Open(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nREAL TIME INDEX\n")
	tmunix := int(time.Now().Unix())
	fmt.Println(tmunix)
	tm := strconv.Itoa(tmunix)
	index := "forumrt"
	// tm := time.Now().String()
	// tm = "2006-01-02T15:04:05Z07:00"
	q := `replace into ` + index + ` values(1, 'my subject', 'my content', 9, 9, 1000, '` + tm + `', 1, 34.67, 'str', '{}')`
	res, err := cl.Sphinxql(q)
	fmt.Println(res, err)
	q = `replace into ` + index + ` values(1, 'my subject', 'my content', 8, 8, 1000, '` + tm + `', 1, 34.67, 'str', '{}')`
	res, err = cl.Sphinxql(q)
	fmt.Println(res, err)
	q = `replace into ` + index + ` values(2,'another subject', 'more content', 8, 8, 1000, '` + tm + `', 1, 34.67, 'str', '{}')`
	res, err = cl.Sphinxql(q)
	fmt.Println(res, err)
	q = `replace into ` + index + ` values(3,'again subject', 'one more content', 8, 8, 1000, '` + tm + `', 1, 34.67, 'str', '{}')`
	res, err = cl.Sphinxql(q)
	fmt.Println(res, err)
	q = `replace into ` + index + ` values(4,'TEST RIDGEMONT', 'OPIS RIDGEMONT ONE', 8, 8, 1000, '` + tm + `', 1, 34.67, 'str', '{}')`
	res, err = cl.Sphinxql(q)
	fmt.Println(res, err)

	res2, err2 := cl.Query("more", index)
	fmt.Println(res2, err2)

	fmt.Println("\nPLAIN INDEX\n")

	res3, err3 := cl.Query("PRIDE")
	fmt.Println(res3, err3)

	qq := manticore.NewSearch("PRIDE", "testidx", "")
	res4, err4 := cl.RunQuery(qq)
	fmt.Println(res4, err4)

}
