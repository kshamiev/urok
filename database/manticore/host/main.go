package main

import (
	"fmt"

	"github.com/manticoresoftware/go-sdk/manticore"
)

func main() {
	cl := manticore.NewClient()
	cl.SetServer("localhost", 9312)
	if _, err := cl.Open(); err != nil {
		fmt.Println(err)
		return
	}

	index := "rt"

	q := `replace into ` + index + ` (id, title) values(1, 'my subject fake')`
	res, err := cl.Sphinxql(q)
	fmt.Println(res, err)
	q = `replace into ` + index + ` (title) values('тузик')`
	res, err = cl.Sphinxql(q)
	fmt.Println(res, err)
	q = `replace into ` + index + ` (title) values('грелка')`
	res, err = cl.Sphinxql(q)
	fmt.Println(res, err)

	res2, err2 := cl.Query("грелка")
	fmt.Println(res2, err2)

	// res3, err3 := cl.Query("PRIDE")
	// fmt.Println(res3, err3)
	// res3, err3 = cl.Query("Пенка")
	// fmt.Println(res3, err3)
}
