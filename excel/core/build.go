package core

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Build struct {
	fp        *excelize.File
	style     int
	sheetName string
	rowBeg    int
	rowEnd    int
	colBeg    string
	colEnd    string
	Err       error
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
	return b.fp.SetCellFormula(b.sheetName, b.colBeg+strconv.Itoa(b.rowBeg), f)
}

func (b *Build) cell() error {
	beg := b.colBeg + strconv.Itoa(b.rowBeg)
	end := b.colEnd + strconv.Itoa(b.rowEnd)
	if beg != end {
		if err := b.fp.MergeCell(b.sheetName, beg, end); err != nil {
			return err
		}
	}
	if 0 < b.style {
		return b.fp.SetCellStyle(b.sheetName, beg, end, b.style)
	}
	return nil
}
