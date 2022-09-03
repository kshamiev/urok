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

	for i := range data {
		b := bu.NewSheet(data[i].Number)
		b.Row += 2
		b.CellRow("C", b.Row, "D", b.Row+2).Width(25).Value("Name")
		b.CellRow("E", b.Row, "F", b.Row+2).Value("Length")
		b.CellRow("G", b.Row, "H", b.Row+2).Value("Width")
		b.CellRow("I", b.Row, "J", b.Row+2).Value("Height")
		b.CellRow("K", b.Row, "L", b.Row+2).Value("Amount")
		b.CellRow("M", b.Row, "N", b.Row+2).Value("Summ")
		b.Row += 3
		for _, elm := range data[i].Cargos {
			b.CellRow("C", b.Row, "D", b.Row+1).Value(elm.Name)
			b.CellRow("E", b.Row, "F", b.Row+1).Value(elm.Length)
			b.CellRow("G", b.Row, "H", b.Row+1).Value(elm.Width)
			b.CellRow("I", b.Row, "J", b.Row+1).Value(elm.Height)
			b.CellRow("K", b.Row, "L", b.Row+1).Value(elm.Amount)
			b.CellRow("M", b.Row, "N", b.Row+1).Value(elm.Summ)
			b.Row += 2
		}

	}

	return bu.GetFp(), nil
}
