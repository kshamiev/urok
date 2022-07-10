package layer_two

import (
	"fmt"
	"log"

	"github.com/kshamiev/urok/sample/excel/typs"
	"github.com/kshamiev/urok/tutorial/gorutine/workerbus"
)

func Action() {
	// подписчик
	ch := workerbus.Gist().Subscribe(&typs.Cargo{})
	go consumer(ch)
}

func consumer(ch chan interface{}) {
	i := 0
	defer func() {
		if rvr := recover(); rvr != nil {
			// log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = workerbus.Gist().Subscribe(&typs.Cargo{})
			go consumer(ch)
		}
	}()
	for obj := range ch {
		o, ok := obj.(*typs.Cargo)
		if !ok {
			close(ch)
			break
		}
		// It`s Work
		fmt.Println(o.Name)
		// ...
		ch <- true
		i++
	}
	log.Println("full count: ", i)
}
