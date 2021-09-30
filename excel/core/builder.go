package core

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io"

	"github.com/xuri/excelize/v2"
)

const TplList = "template"

type Builder struct {
	fp           *excelize.File
	stHeaderMain int
	stHeader     int
	stHeaderSub  int
	stData       int
	stFooter     int
	sheetName    string
	Row          int
}

func NewBuilder(fp *excelize.File) (*Builder, error) {
	comp := &Builder{fp: fp, Row: 1}
	var err error

	if comp.stHeaderMain, err = fp.GetCellStyle(TplList, "B1"); err != nil {
		return nil, err
	}
	if comp.stHeader, err = fp.GetCellStyle(TplList, "B2"); err != nil {
		return nil, err
	}
	if comp.stHeaderSub, err = fp.GetCellStyle(TplList, "B3"); err != nil {
		return nil, err
	}
	if comp.stData, err = fp.GetCellStyle(TplList, "B34"); err != nil {
		return nil, err
	}
	if comp.stFooter, err = fp.GetCellStyle(TplList, "B20"); err != nil {
		return nil, err
	}

	return comp, nil
}

func (comp *Builder) NewSheet(name string) int {
	comp.sheetName = name
	return comp.fp.NewSheet(name)
}

func (comp *Builder) HeaderMain(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return &Build{
		fp:        comp.fp,
		style:     comp.stHeaderMain,
		sheetName: comp.sheetName,
		rowBeg:    rowBeg,
		rowEnd:    rowEnd,
		colBeg:    colBeg,
		colEnd:    colEnd,
		Err:       nil,
	}
	// _ = comp.SetRowHeight(21.0)
}
func (comp *Builder) Header(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return &Build{
		fp:        comp.fp,
		style:     comp.stHeader,
		sheetName: comp.sheetName,
		rowBeg:    rowBeg,
		rowEnd:    rowEnd,
		colBeg:    colBeg,
		colEnd:    colEnd,
		Err:       nil,
	}
	// _ = comp.SetRowHeight(18.0)
}
func (comp *Builder) HeaderSub(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return &Build{
		fp:        comp.fp,
		style:     comp.stHeaderSub,
		sheetName: comp.sheetName,
		rowBeg:    rowBeg,
		rowEnd:    rowEnd,
		colBeg:    colBeg,
		colEnd:    colEnd,
		Err:       nil,
	}
	// _ = comp.SetRowHeight(16.0)
}
func (comp *Builder) Data(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return &Build{
		fp:        comp.fp,
		style:     comp.stData,
		sheetName: comp.sheetName,
		rowBeg:    rowBeg,
		rowEnd:    rowEnd,
		colBeg:    colBeg,
		colEnd:    colEnd,
		Err:       nil,
	}
	// _ = comp.SetRowHeight(16.0)
}
func (comp *Builder) Footer(colBeg, colEnd string, rowBeg, rowEnd int) *Build {
	return &Build{
		fp:        comp.fp,
		style:     comp.stFooter,
		sheetName: comp.sheetName,
		rowBeg:    rowBeg,
		rowEnd:    rowEnd,
		colBeg:    colBeg,
		colEnd:    colEnd,
		Err:       nil,
	}
	// _ = comp.SetRowHeight(16.0)
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
