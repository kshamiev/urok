package layer_one

import (
	"strconv"
	"time"

	"github.com/kshamiev/urok/tutorial/gorutine/workerbus"
	"github.com/kshamiev/urok/tutorial/gorutine/workerbus/typs"
)

func Action() {
	// отправитель
	go func() {
		for i := 0; i < 1000000; i++ {
			workerbus.Gist().SendData(&typs.Cargo{Name: "additional_" + strconv.Itoa(i), Amount: 1})
			time.Sleep(time.Second)
		}
	}()
}
