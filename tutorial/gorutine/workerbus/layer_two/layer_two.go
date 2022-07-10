package layer_two

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
				Name:   "send_layer_two_type_General" + strconv.Itoa(i),
				Amount: 1,
			})
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			workerbus.Gist().SendData(&typs.LayerTwo{
				Name:   "send_layer_two_type_LayerTwo" + strconv.Itoa(i),
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
	ch = workerbus.Gist().Subscribe(&typs.LayerOne{})
	go consumerLayerOne(ch)
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

func consumerLayerOne(ch chan interface{}) {
	i := 0
	defer func() {
		if rvr := recover(); rvr != nil {
			// log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = workerbus.Gist().Subscribe(&typs.LayerOne{})
			go consumerLayerOne(ch)
		}
	}()
	for obj := range ch {
		o, ok := obj.(*typs.LayerOne)
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
