package test

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"gitlab.tn.ru/golang/kit/tpl"
)

func TestStorage(t *testing.T) {
	f := map[string]interface{}{
		"TplTest": func(name string) string {
			return "<H1>" + name + "</H1>"
		},
	}
	tt, err := tpl.NewTplGenerator("www", f)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)
	t.Log(tt.GetStorageIndex())
}

func TestTplStorage(t *testing.T) {
	f := map[string]interface{}{
		"TplTest": func(name string) string {
			return "<H1>" + name + "</H1>"
		},
	}
	tt, err := tpl.NewTplGenerator("www", f)
	if err != nil {
		t.Fatal(err)
	}

	goods := Goods{
		{ID: 37, Name: "Item 10", Price: decimal.NewFromFloat(23.76)},
		{ID: 49, Name: "Item 2", Price: decimal.NewFromFloat(87.42)},
		{ID: 54, Name: "Item 30", Price: decimal.NewFromFloat(38.23)},
	}
	variable := map[string]interface{}{
		"Title": "TestTplStorage",
		"Goods": goods,
	}

	for _, i := range tt.GetStorageIndex() {
		t.Log(i)
		if _, err := tt.ExecuteStorage(i, variable); err != nil {
			t.Fatal(err)
		}
	}
}

func TestFile(t *testing.T) {
	f := map[string]interface{}{
		"TplTest": func(name string) string {
			return "<H1>" + name + "</H1>"
		},
	}

	goods := Goods{
		{ID: 23, Name: "Item 1", Price: decimal.NewFromFloat(45.76)},
		{ID: 34, Name: "Item 2", Price: decimal.NewFromFloat(12.42)},
		{ID: 45, Name: "Item 3", Price: decimal.NewFromFloat(74.23)},
	}
	variable := map[string]interface{}{
		"Title": "TestFile",
		"Goods": goods,
	}

	_, err := tpl.ExecuteFile("www/index.html", f, variable)
	if err != nil {
		t.Fatal(err)
	}
}

func TestString(t *testing.T) {
	f := map[string]interface{}{
		"TplTest": func(name string) string {
			return "<H1>" + name + "</H1>"
		},
	}

	goods := Goods{
		{ID: 23, Name: "Item 1", Price: decimal.NewFromFloat(45.76)},
		{ID: 34, Name: "Item 2", Price: decimal.NewFromFloat(12.42)},
		{ID: 45, Name: "Item 3", Price: decimal.NewFromFloat(74.23)},
	}
	variable := map[string]interface{}{
		"Title": "TestString",
		"Goods": goods,
	}

	_, err := tpl.ExecuteString(testTpl, f, variable)
	if err != nil {
		t.Fatal(err)
	}
}
