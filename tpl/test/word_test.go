package test

import (
	"testing"

	"gitlab.tn.ru/golang/kit/tpl"
)

func TestWord(t *testing.T) {
	err := tpl.WordReplace("document.docx", "document_new.docx",
		map[string]string{
			"{Tortik}": "Фикус губоцветный",
			"{Tazik}":  "Дон Кихот Ламанчский",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
}
