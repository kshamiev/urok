package main

import (
	"log"

	"github.com/kshamiev/urok/sample/excel/docexcel"
)

func main() {
	var data interface{}
	doc := docexcel.NewSample()
	fp, err := doc.Compile(data)
	if err != nil {
		log.Fatal(err)
	}
	if err := fp.SaveAs("sample/excel/app/sample2/test.xlsx"); err != nil {
		log.Fatal(err)
	}
}
