package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/kshamiev/urok/sample/excel/sample1/docexcel"
	"github.com/kshamiev/urok/sample/excel/sample1/typs"
)

func main() {
	_, filePath, _, _ := runtime.Caller(0)
	filePath = filepath.Dir(filePath)
	data := []typs.InvoiceTC{
		typs.NewInvoiceTC("2345"),
		typs.NewInvoiceTC("6578"),
		typs.NewInvoiceTC("83464"),
	}
	doc := docexcel.NewSample(filePath + "/combined.xlsx")
	fp, err := doc.Compile(data)
	if err != nil {
		log.Fatal(err)
	}
	if err := fp.SaveAs(filePath + "/test.xlsx"); err != nil {
		log.Fatal(err)
	}
}
