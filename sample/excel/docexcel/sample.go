package docexcel

import (
	"github.com/xuri/excelize/v2"

	"github.com/kshamiev/urok/sample/excel/excel"
	"github.com/kshamiev/urok/sample/excel/typs"
)

type Sample struct{}

func NewSample() Sample {
	return Sample{}
}

func (doc Sample) Compile(data []typs.InvoiceTC) (*excelize.File, error) {
	bu, err := excel.NewBuilderFile("sample/excel/app/sample2/combined.xlsx")
	if err != nil {
		return nil, err
	}
	defer func() {
		bu.DeleteStartSheet()
	}()

	// b := bu.NewSheet("test")
	// b.Style(doc.styleH3).Cell("B", 2, "E", 2).Value("Заголовок")

	for _, inv := range data {
		b := bu.NewSheet(inv.Number)
		if err := doc.HeaderMain(b, inv); err != nil {
			return nil, err
		}
		if err := doc.InitiatorA(b, inv); err != nil {
			return nil, err
		}
	}
	return bu.GetFp(), nil
}

func (doc Sample) HeaderMain(b *excel.Build, inv typs.InvoiceTC) (err error) {
	doc.Head1(b).Cell("B", b.Row, "S", b.Row).Height(21).
		Value("Заявка на организацию транспортно-экспедиционного обслуживания № " + inv.Number)
	b.Row++
	return b.Err
}

func (doc Sample) InitiatorA(b *excel.Build, inv typs.InvoiceTC) (err error) {
	doc.Head3(b).Cell("B", b.Row, "S", b.Row).Height(18).Value("Инициатор")
	b.Row++
	doc.Head4(b).Cell("B", b.Row, "B", b.Row).Height(16).Value("ФИО")
	doc.Body1(b).Cell("C", b.Row, "H", b.Row).Value("${fullname}")
	doc.Head4(b).Cell("I", b.Row, "I", b.Row).Value("Тел")
	doc.Body1(b).Cell("J", b.Row, "M", b.Row).Value("${phone}")
	doc.Head4(b).Cell("N", b.Row, "N", b.Row).Value("E-mail")
	doc.Body1(b).Cell("O", b.Row, "S", b.Row).Value("${email}")
	b.Row += 2
	return b.Err
}

func (doc Sample) Head1(b *excel.Build) *excel.Build {
	return b.Style("C2")
}
func (doc Sample) Head2(b *excel.Build) *excel.Build {
	return b.Style("C4")
}
func (doc Sample) Head3(b *excel.Build) *excel.Build {
	return b.Style("C6")
}
func (doc Sample) Head4(b *excel.Build) *excel.Build {
	return b.Style("C8")
}
func (doc Sample) Body1(b *excel.Build) *excel.Build {
	return b.Style("C15")
}
func (doc Sample) Footer1(b *excel.Build) *excel.Build {
	return b.Style("C29")
}
func (doc Sample) Formula1(b *excel.Build) *excel.Build {
	return b.Style("C22")
}
