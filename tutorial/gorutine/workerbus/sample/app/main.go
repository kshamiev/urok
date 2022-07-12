package main

import (
	"fmt"

	"github.com/kshamiev/urok/tutorial/gorutine/workerbus"
	"github.com/kshamiev/urok/tutorial/gorutine/workerbus/sample/layer_one"
	"github.com/kshamiev/urok/tutorial/gorutine/workerbus/sample/layer_two"
)

func main() {
	workerbus.Init(100000, 3, true)

	layer_one.ActionSend()
	layer_two.ActionSend()

	layer_one.ActionConsumer()
	layer_two.ActionConsumer()

	var tt int
	fmt.Scanln(&tt)

	workerbus.Gist().Wait()
}
