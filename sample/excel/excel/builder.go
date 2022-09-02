package excel

import (
	"io"

	"github.com/xuri/excelize/v2"
)

type StyleXY string

type Builder struct {
	fp         *excelize.File
	sheetList  []string
	styleStore map[StyleXY]int
}

func NewBuilder() *Builder {
	fp := excelize.NewFile()
	return &Builder{
		fp:         fp,
		sheetList:  fp.GetSheetList(),
		styleStore: make(map[StyleXY]int),
	}
}

func NewBuilderFile(templatePath string) (*Builder, error) {
	fp, err := excelize.OpenFile(templatePath)
	if err != nil {
		return nil, err
	}
	return &Builder{
		fp:         fp,
		sheetList:  fp.GetSheetList(),
		styleStore: make(map[StyleXY]int),
	}, nil
}

func NewBuilderReader(r io.Reader) (*Builder, error) {
	fp, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	return &Builder{
		fp:         fp,
		sheetList:  fp.GetSheetList(),
		styleStore: make(map[StyleXY]int),
	}, nil
}

func (b *Builder) GetFp() *excelize.File {
	return b.fp
}

func (b *Builder) DeleteStartSheet() {
	for i := range b.sheetList {
		b.fp.DeleteSheet(b.sheetList[i])
	}
}

func (b *Builder) NewSheet(name string) *Build {
	b.fp.NewSheet(name)
	return &Build{
		builder:   b,
		sheetName: name,
		Row:       1,
	}
}
