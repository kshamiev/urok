package excel

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

const styleSheet = "style"

type StyleXY string

type Build struct {
	fp        *excelize.File
	style     int
	sheetName string
	rowBeg    int
	rowEnd    int
	colBeg    string
	colEnd    string

	styleStore map[StyleXY]int

	Row int
	Err error
}

func NewBuild(fp *excelize.File) *Build {
	return &Build{
		fp: fp,
	}
}

func (b *Build) NewSheet(name string) *Build {
	b.sheetName = name
	b.Row = 1
	b.fp.NewSheet(name)
	return b
}

func (b *Build) Style(position StyleXY) *Build {
	if string(position) == "" {
		b.style = 0
		return b
	}
	if _, ok := b.styleStore[position]; !ok {
		b.styleStore[position], b.Err = b.fp.GetCellStyle(styleSheet, string(position))
	}
	b.style = b.styleStore[position]
	return b
}

func (b *Build) Cell(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	b.rowBeg = rowBeg
	b.rowEnd = rowEnd
	b.colBeg = colBeg
	b.colEnd = colEnd
	return b
}

func (b *Build) Height(h float64) *Build {
	if b.Err != nil {
		return b
	}
	b.Err = b.fp.SetRowHeight(b.sheetName, b.rowBeg, h)
	return b
}

func (b *Build) Width(h float64) *Build {
	if b.Err != nil {
		return b
	}
	b.Err = b.fp.SetColWidth(b.sheetName, b.colBeg, b.colEnd, h)
	return b
}

func (b *Build) Value(value interface{}) error {
	if err := b.cell(); err != nil {
		return err
	}
	return b.fp.SetCellValue(b.sheetName, b.colBeg+strconv.Itoa(b.rowBeg), value)
}

func (b *Build) Formula(f string) error {
	if err := b.cell(); err != nil {
		return err
	}
	if err := b.fp.SetCellFormula(b.sheetName, b.colBeg+strconv.Itoa(b.rowBeg), f); err != nil {
		return err
	}
	return nil
}

func (b *Build) cell() error {
	beg := b.colBeg + strconv.Itoa(b.rowBeg)
	end := b.colEnd + strconv.Itoa(b.rowEnd)
	if beg != end {
		if err := b.fp.MergeCell(b.sheetName, beg, end); err != nil {
			return err
		}
	}
	if b.style > 0 {
		return b.fp.SetCellStyle(b.sheetName, beg, end, b.style)
	}
	return nil
}
