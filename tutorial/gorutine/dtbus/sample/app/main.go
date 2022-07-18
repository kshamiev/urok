package main

import (
	"fmt"

	"gitlab.tn.ru/golang/app/dtbus"
	"gitlab.tn.ru/golang/app/dtbus/sample/layer_one"
	"gitlab.tn.ru/golang/app/dtbus/sample/layer_two"
)

func main() {
	dtbus.Init(100000)

	layer_one.ActionSend()
	layer_two.ActionSend()

	layer_one.ActionConsumer()
	layer_two.ActionConsumer()

	var tt int
	fmt.Scanln(&tt)

	dtbus.Gist().Wait()
}
