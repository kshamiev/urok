//// +build ignore go1.7 // запрет компиляции с указанием версии go
// package routing // import "application/routing" // указывает как импортировать пакет и где он должен находится
// точка входа в программу

package main

import (
	"fmt"
	"urok/webserver"
)

// Эта функция выполняется в момент запуска программы, перед входом в функцию main
// Порядок выполнения таких функций аналогичен стеку defer
// (последний вложенный первым и первый вложенный последним)
func init() {
	fmt.Println("main")
}

func main() {

	webserver.Sample2()

	// urok.SamplechanelBlock()

	//	gorun.SamplechanelBlock()
	//	gorun.SampleWork()

	//	control.Control()
	//	if err := control.PanicTest(); err != nil {
	//		fmt.Println("Error: ", err)
	//	}

	//	function.Closure()
	//	function.Function()

	//	structandface.SampleStruct()
	//	structandface.SampleFace()

	//	variable.Rune()
	//	variable.Sample()

	// init(), пространство имен - алиасы, подключение без использования
	// urok.Test()

}
