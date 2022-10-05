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
	res, err := cl.Sphinxql(`replace into testrt values(1,'my subject', 'my content', 15)`)
	fmt.Println(res, err)
	res, err = cl.Sphinxql(`replace into testrt values(2,'another subject', 'more content', 15)`)
	fmt.Println(res, err)
	res, err = cl.Sphinxql(`replace into testrt values(5,'again subject', 'one more content', 10)`)
	fmt.Println(res, err)

	q := manticore.NewSearch("content", "testrt", "")
	res2, err2 := cl.RunQuery(q)
	fmt.Println(res2, err2)

	attrValues := []int64{10}
	q.AddFilter("gid", attrValues, true)
	res2, err2 = cl.RunQuery(q)
	fmt.Println(res2, err2)
}
