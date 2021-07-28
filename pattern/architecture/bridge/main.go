// Мост
// Это тип который реализует часть интерфейса (часть его методов)
// И далее он встраивается в нужные типы.
// Таким образом дополняя их до нужной реализации и объединяя в нужный интерфейс
//
package main

import (
	"fmt"
	"strings"
)

func main() {
	textMaker := fillTextBuilder(&TextBuilder{})
	text := textMaker.getText()
	// test: line 1
	//      line 2

	htmlMaker := fillTextBuilder(&HTMLBuilder{})
	html := htmlMaker.getText()
	// html: <span>line 1</span><br/>
	//      <span>line 2</span><br/>

	fmt.Println(text)
	fmt.Println(html)
}

// AText is Abstraction
type AText interface {
	getText() string
	addLine(value string)
}

// TextMaker RefinedAbstraction
type TextMaker struct {
	textImp ITextImp
}

func (tm TextMaker) getText() string {
	return tm.textImp.getString()
}

func (tm TextMaker) addLine(value string) {
	tm.textImp.appendLine(value)
}

// ITextImp is abstract Implementator
type ITextImp interface {
	getString() string
	appendLine(value string)
}

// TextImp is Implementator
type TextImp struct {
	rows []string
}

func (ti TextImp) getString() string {
	return strings.Join(ti.rows, "\n")
}

// ////

// TextBuilder is ConcreteImplementator1
type TextBuilder struct {
	TextImp
}

func (tb *TextBuilder) appendLine(value string) {
	tb.rows = append(tb.rows, value)
}

// HTMLBuilder is ConcreteImplementator2
type HTMLBuilder struct {
	TextImp
}

func (hb *HTMLBuilder) appendLine(value string) {
	hb.rows = append(hb.rows, "<span>"+value+"</span><br/>")
}

func fillTextBuilder(textImp ITextImp) AText {
	textMaker := TextMaker{textImp}
	textMaker.addLine("line 1")
	textMaker.addLine("line 2")
	return textMaker
}
