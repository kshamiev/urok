// Композиция
// Это когда типы реализующие различный функционал связанный бизнес логикой
// включаются как свойства в один тип как собирательный образ
package main

import (
	"fmt"
	"strings"
)

func main() {
	adapter := NewTextAdapter()
	text := adapter.getText()
	fmt.Println(text)
	fmt.Println(adapter.RowList.getString())
}

// StringList is Adaptee
type StringList struct {
	rows []string
}

func (sl StringList) getString() string {
	return strings.Join(sl.rows, "\n")
}

func (sl *StringList) add(value string) {
	sl.rows = append(sl.rows, value)
}

// TextAdapter is Adapter
type TextAdapter struct {
	RowList StringList
}

func (ta TextAdapter) getText() string {
	return ta.RowList.getString()
}

func NewTextAdapter() TextAdapter {
	rowList := StringList{}
	rowList.add("line 1")
	rowList.add("line 2")
	return TextAdapter{rowList}
}
