package main

import (
	"fmt"

	"github.com/kshamiev/urok/tutorial/gorutine/dtbus"
	"github.com/kshamiev/urok/tutorial/gorutine/dtbus/sample/layer_one"
	"github.com/kshamiev/urok/tutorial/gorutine/dtbus/sample/layer_two"
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
