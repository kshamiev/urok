package main

import (
	"log"

	"github.com/kshamiev/urok/sample/excel/assembly"
)

func main() {
	data := []assembly.InvoiceTC{
		assembly.NewInvoiceTCSample("2345"),
		assembly.NewInvoiceTCSample("6578"),
		assembly.NewInvoiceTCSample("83464"),
	}
	fp, err := assembly.InvoiceTCTrucking("sample/excel/combined.xlsx", data)
	if err != nil {
		log.Fatal(err)
	}
	if err := fp.SaveAs("sample/excel/test.xlsx"); err != nil {
		log.Fatal(err)
	}
}