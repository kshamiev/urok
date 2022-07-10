package layer_one

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kshamiev/urok/tutorial/gorutine/workerbus"
	"github.com/kshamiev/urok/tutorial/gorutine/workerbus/typs"
)

func ActionSend() {
	// отправитель
	go func() {
		for i := 0; i < 1000000; i++ {
			workerbus.Gist().SendData(&typs.General{
				Name:   "additional_layer_one_General" + strconv.Itoa(i),
				Amount: 1,
			})
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			workerbus.Gist().SendData(&typs.LayerOne{
				Name:   "additional_layer_one_LayerOne" + strconv.Itoa(i),
				Amount: 1,
			})
			time.Sleep(time.Second)
		}
	}()
}

func ActionConsumer() {
	// подписчик
	ch := workerbus.Gist().Subscribe(&typs.General{})
	go consumerGeneral(ch)
	ch = workerbus.Gist().Subscribe(&typs.LayerTwo{})
	go consumerLayerTwo(ch)
}

func consumerGeneral(ch chan interface{}) {
	i := 0
	defer func() {
		if rvr := recover(); rvr != nil {
			// log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = workerbus.Gist().Subscribe(&typs.General{})
			go consumerGeneral(ch)
		}
	}()
	for obj := range ch {
		o, ok := obj.(*typs.General)
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

func consumerLayerTwo(ch chan interface{}) {
	i := 0
	defer func() {
		if rvr := recover(); rvr != nil {
			// log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = workerbus.Gist().Subscribe(&typs.LayerTwo{})
			go consumerLayerTwo(ch)
		}
	}()
	for obj := range ch {
		o, ok := obj.(*typs.LayerTwo)
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
