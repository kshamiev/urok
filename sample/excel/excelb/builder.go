package excelb

import (
	"github.com/xuri/excelize/v2"
)

const TplStyle = "style"

type Builder struct {
	fp *excelize.File

	header1  int
	header2  int
	header3  int
	header4  int
	header5  int
	header6  int
	content1 int
	content2 int
	content3 int
	formula1 int
	formula2 int
	formula3 int
	footer1  int
	footer2  int
	footer3  int

	sheetName string

	Row int
}

func NewBuilder(fp *excelize.File) (*Builder, error) {
	b := &Builder{fp: fp, Row: 1}
	var err error

	if b.header1, err = fp.GetCellStyle(TplStyle, "C2"); err != nil {
		return nil, err
	}
	if b.header2, err = fp.GetCellStyle(TplStyle, "C4"); err != nil {
		return nil, err
	}
	if b.header3, err = fp.GetCellStyle(TplStyle, "C6"); err != nil {
		return nil, err
	}
	if b.header4, err = fp.GetCellStyle(TplStyle, "C8"); err != nil {
		return nil, err
	}
	if b.header5, err = fp.GetCellStyle(TplStyle, "C10"); err != nil {
		return nil, err
	}
	if b.header6, err = fp.GetCellStyle(TplStyle, "C12"); err != nil {
		return nil, err
	}
	if b.content1, err = fp.GetCellStyle(TplStyle, "C15"); err != nil {
		return nil, err
	}
	if b.content2, err = fp.GetCellStyle(TplStyle, "C17"); err != nil {
		return nil, err
	}
	if b.content3, err = fp.GetCellStyle(TplStyle, "C19"); err != nil {
		return nil, err
	}
	if b.formula1, err = fp.GetCellStyle(TplStyle, "C22"); err != nil {
		return nil, err
	}
	if b.formula2, err = fp.GetCellStyle(TplStyle, "C24"); err != nil {
		return nil, err
	}
	if b.formula3, err = fp.GetCellStyle(TplStyle, "C26"); err != nil {
		return nil, err
	}
	if b.footer1, err = fp.GetCellStyle(TplStyle, "C29"); err != nil {
		return nil, err
	}
	if b.footer2, err = fp.GetCellStyle(TplStyle, "C31"); err != nil {
		return nil, err
	}
	if b.footer3, err = fp.GetCellStyle(TplStyle, "C33"); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Builder) NewSheet(name string) int {
	b.sheetName = name
	b.Row = 1
	return b.fp.NewSheet(name)
}

func (b *Builder) Header1(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.header1)
}
func (b *Builder) Header2(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.header2)
}
func (b *Builder) Header3(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.header3)
}
func (b *Builder) Header4(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.header4)
}
func (b *Builder) Header5(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.header5)
}
func (b *Builder) Header6(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.header6)
}

func (b *Builder) Content1(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.content1)
}
func (b *Builder) Content2(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.content2)
}
func (b *Builder) Content3(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.content3)
}

func (b *Builder) Formula1(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.formula1)
}
func (b *Builder) Formula2(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.formula2)
}
func (b *Builder) Formula3(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.formula3)
}

func (b *Builder) Footer1(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.footer1)
}
func (b *Builder) Footer2(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.footer2)
}
func (b *Builder) Footer3(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return b.Cell(colBeg, colEnd, rowBeg, rowEnd, b.footer3)
}

func (b *Builder) Cell(colBeg, colEnd string, rowBeg, rowEnd, style int) *Build {
	return &Build{
		fp:        b.fp,
		style:     style,
		sheetName: b.sheetName,
		rowBeg:    rowBeg,
		rowEnd:    rowEnd,
		colBeg:    colBeg,
		colEnd:    colEnd,
		Err:       nil,
	}
}
