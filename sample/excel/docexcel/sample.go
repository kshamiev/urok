package docexcel

import (
	"github.com/xuri/excelize/v2"

	"github.com/kshamiev/urok/sample/excel/excel"
)

type Sample struct {
	templatePath string
	styleH1      excel.StyleXY
	styleH2      excel.StyleXY
	styleH3      excel.StyleXY
	styleH4      excel.StyleXY
	styleC1      excel.StyleXY
	styleFR1     excel.StyleXY
	styleFT1     excel.StyleXY
}

func NewSample() Sample {
	return Sample{
		templatePath: "sample/excel/app/sample2/combined.xlsx",
		styleH1:      "C2",
		styleH2:      "C4",
		styleH3:      "C6",
		styleH4:      "C8",
		styleC1:      "C15",
		styleFR1:     "C22",
		styleFT1:     "C29",
	}
}

func (doc Sample) Compile(data interface{}) (*excelize.File, error) {
	bu, err := excel.NewBuilderFile(doc.templatePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		bu.DeleteStartSheet()
	}()

	b := bu.NewSheet("test")
	b.Style(doc.styleH3).Cell("B", 2, "E", 2).Value("Заголовок")

	return bu.GetFp(), nil
}
