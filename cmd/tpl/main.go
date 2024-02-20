package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"gitlab.tn.ru/golang/kit/tpl"
)

// install
// go install github.com/kshamiev/urok/cmd/tpl@v1.0.4
// go install github.com/kshamiev/urok/cmd/tpl@latest

// use:
// Если пути относительные то отправной точкой считается директория из которой запущена программа
// tpl -h
// tpl -json data.json -in page_in.html -out page_out.html

func main() {
	f := map[string]interface{}{
		"TplTest": func(name string) string {
			return "<H1>" + name + "</H1>"
		},
	}

	fileJson := flag.String("json", "data.json", "данные в формате json")
	fileIn := flag.String("in", "page_in.html", "исходный html шаблон")
	fileOut := flag.String("out", "page_out.html", "результирующая html страница с данными")
	flag.Parse()

	data, err := os.ReadFile(*fileJson)
	if err != nil {
		fmt.Println("ошибка json файла, данные ошибочны либо файл не найден")
		log.Fatalln(err)
	}
	variable := map[string]interface{}{}
	err = json.Unmarshal(data, &variable)
	if err != nil {
		fmt.Println("ошибка распаковки данных в переменную")
		log.Fatalln(err)
	}

	res, err := tpl.ExecuteFile(*fileIn, f, variable)
	if err != nil {
		fmt.Println("ошибка компиляции шаблона, возможно файл не найден")
		log.Fatalln(err)
	}

	err = os.WriteFile(*fileOut, res.Bytes(), 0o600)
	if err != nil {
		fmt.Println("ошибка сохранения результирующей страницы")
		log.Fatalln(err)
	}
}
