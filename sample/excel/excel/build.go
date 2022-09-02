package excel

import (
	"strconv"
)

type Build struct {
	builder   *Builder
	sheetName string
	style     int
	rowBeg    int
	rowEnd    int
	colBeg    string
	colEnd    string
	Row       int
	Err       error
}

func (b *Build) Style(position StyleXY) *Build {
	if string(position) == "" {
		b.style = 0
		return b
	}
	if _, ok := b.builder.styleStore[position]; !ok {
		b.builder.styleStore[position], b.Err = b.builder.fp.GetCellStyle(b.builder.sheetList[0], string(position))
	}
	b.style = b.builder.styleStore[position]
	return b
}

func (b *Build) Cell(colBeg string, rowBeg int, colEnd string, rowEnd int) *Build {
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
	b.Err = b.builder.fp.SetRowHeight(b.sheetName, b.rowBeg, h)
	return b
}

func (b *Build) Width(h float64) *Build {
	if b.Err != nil {
		return b
	}
	b.Err = b.builder.fp.SetColWidth(b.sheetName, b.colBeg, b.colEnd, h)
	return b
}

func (b *Build) Value(value interface{}) {
	if err := b.cell(); err != nil {
		b.Err = err
		return
	}
	b.Err = b.builder.fp.SetCellValue(b.sheetName, b.colBeg+strconv.Itoa(b.rowBeg), value)
}

func (b *Build) ValueFormula(f string) {
	if err := b.cell(); err != nil {
		b.Err = err
		return
	}
	if err := b.builder.fp.SetCellFormula(b.sheetName, b.colBeg+strconv.Itoa(b.rowBeg), f); err != nil {
		b.Err = err
		return
	}
}

func (b *Build) cell() error {
	beg := b.colBeg + strconv.Itoa(b.rowBeg)
	end := b.colEnd + strconv.Itoa(b.rowEnd)
	if beg != end {
		if err := b.builder.fp.MergeCell(b.sheetName, beg, end); err != nil {
			return err
		}
	}
	if b.style > 0 {
		return b.builder.fp.SetCellStyle(b.sheetName, beg, end, b.style)
	}
	return nil
}
