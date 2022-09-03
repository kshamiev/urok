package docexcel

import (
	"github.com/xuri/excelize/v2"

	"github.com/kshamiev/urok/sample/excel/excel"
	"github.com/kshamiev/urok/sample/excel/excel/sample2/typs"
)

type Sample struct{}

func NewSample() Sample {
	return Sample{}
}

func (doc Sample) Compile(data []typs.InvoiceTC) (*excelize.File, error) {
	bu := excel.NewBuilder()
	defer func() {
		bu.DeleteStartSheet()
	}()

	b := bu.NewSheet("test")
	b.CellRow("B", 2, "E", 2).Value("Заголовок")

	return bu.GetFp(), nil
}
