// Композиция (через встраивание)
// Это когда типы реализующие различный функционал связанный бизнес логикой
// включаются как свойства в один тип как собирательный образ
package main

import (
	"fmt"
	"strings"
)

func main() {
	adapter := getTextAdapter()
	text := adapter.getText()
	fmt.Println(text)
	fmt.Println(adapter.getString())
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
	StringList
}

func (ta TextAdapter) getText() string {
	return ta.getString()
}

func getTextAdapter() TextAdapter {
	adapter := TextAdapter{}
	adapter.add("line 1")
	adapter.add("line 2")
	return adapter
}
