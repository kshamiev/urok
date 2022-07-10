package main

import (
	"fmt"

	"github.com/kshamiev/urok/tutorial/gorutine/workerbus"
	"github.com/kshamiev/urok/tutorial/gorutine/workerbus/layer_one"
	"github.com/kshamiev/urok/tutorial/gorutine/workerbus/layer_two"
)

func main() {
	workerbus.Init(100000, 3)

	layer_one.Action()
	layer_two.Action()

	var tt int
	fmt.Scanln(&tt)

	workerbus.Gist().Wait()
}
