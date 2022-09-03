package main

import (
	"log"

	"github.com/kshamiev/urok/sample/excel/docexcel"
	"github.com/kshamiev/urok/sample/excel/typs"
)

func main() {
	data := []typs.InvoiceTC{
		typs.NewInvoiceTC("2345"),
		typs.NewInvoiceTC("6578"),
		typs.NewInvoiceTC("83464"),
	}
	doc := docexcel.NewSample()
	fp, err := doc.Compile(data)
	if err != nil {
		log.Fatal(err)
	}
	if err := fp.SaveAs("sample/excel/app/sample2/test.xlsx"); err != nil {
		log.Fatal(err)
	}
}
