package core

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"strconv"

	"github.com/xuri/excelize/v2"
)

const TplList = "template"

type Component struct {
	fp           *excelize.File
	StHeaderMain int
	StHeader     int
	StHeaderSub  int
	StData       int
	StFooter     int
	sheetName    string
	row          int
}

func NewComponent(fp *excelize.File) (*Component, error) {
	comp := &Component{fp: fp, row: 1}
	var err error

	if comp.StHeaderMain, err = fp.GetCellStyle(TplList, "B1"); err != nil {
		return nil, err
	}
	if comp.StHeader, err = fp.GetCellStyle(TplList, "B2"); err != nil {
		return nil, err
	}
	if comp.StHeaderSub, err = fp.GetCellStyle(TplList, "B3"); err != nil {
		return nil, err
	}
	if comp.StData, err = fp.GetCellStyle(TplList, "B34"); err != nil {
		return nil, err
	}
	if comp.StFooter, err = fp.GetCellStyle(TplList, "B20"); err != nil {
		return nil, err
	}

	return comp, nil
}

func (comp *Component) NewSheet(name string) int {
	comp.sheetName = name
	comp.row = 1
	return comp.fp.NewSheet(name)
}

func (comp *Component) SetRow(row int) {
	comp.row = row
}

func (comp *Component) SetRowHeight(h float64) error {
	return comp.fp.SetRowHeight(comp.sheetName, comp.row, h)
}

func (comp *Component) SetRowMove(row int) {
	comp.row += row
}

func (comp *Component) SetColWidth(beg, end string, h float64) error {
	return comp.fp.SetColWidth(comp.sheetName, beg, end, h)
}

func (comp *Component) SetCellStHeaderMain(beg, end string, value interface{}) error {
	_ = comp.SetRowHeight(21.0)
	return comp.SetCellSt(beg, end, value, comp.StHeaderMain)
}
func (comp *Component) SetCellStHeader(beg, end string, value interface{}) error {
	_ = comp.SetRowHeight(18.0)
	return comp.SetCellSt(beg, end, value, comp.StHeader)
}
func (comp *Component) SetCellStHeaderSub(beg, end string, value interface{}) error {
	_ = comp.SetRowHeight(16.0)
	return comp.SetCellSt(beg, end, value, comp.StHeaderSub)
}
func (comp *Component) SetCellStData(beg, end string, value interface{}) error {
	_ = comp.SetRowHeight(16.0)
	return comp.SetCellSt(beg, end, value, comp.StData)
}
func (comp *Component) SetCellStFooter(beg, end string, value interface{}) error {
	_ = comp.SetRowHeight(16.0)
	return comp.SetCellSt(beg, end, value, comp.StFooter)
}

func (comp *Component) SetCellSt(beg, end string, value interface{}, style int) error {
	beg = beg + strconv.Itoa(comp.row)
	end = end + strconv.Itoa(comp.row)
	if beg != end {
		if err := comp.fp.MergeCell(comp.sheetName, beg, end); err != nil {
			return err
		}
	}
	if err := comp.fp.SetCellStyle(comp.sheetName, beg, end, style); err != nil {
		return err
	}
	if err := comp.fp.SetCellValue(comp.sheetName, beg, value); err != nil {
		return err
	}
	return nil
}

// Dump all variables to STDOUT
// From local debug
func Dumper(idl ...interface{}) string {
	ret := dump(idl...)
	fmt.Print(ret.String())

	return ret.String()
}

// dump all variables to bytes.Buffer
func dump(idl ...interface{}) bytes.Buffer {
	var buf bytes.Buffer

	var wr = io.MultiWriter(&buf)

	for _, field := range idl {
		fset := token.NewFileSet()
		_ = ast.Fprint(wr, fset, field, ast.NotNilFilter)
	}

	return buf
}
