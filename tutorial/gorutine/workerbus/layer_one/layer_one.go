package layer_one

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	"github.com/kshamiev/urok/tutorial/gorutine/workerbus"
	"github.com/kshamiev/urok/tutorial/gorutine/workerbus/typs"
)

func ActionSend() {
	// отправитель
	go func() {
		var i int
		for {
			i++
			workerbus.Gist().SendData(&typs.General{
				Name:   "send_layer_one_type_General" + strconv.Itoa(i),
				Amount: 1,
			})
			time.Sleep(time.Millisecond)
		}
	}()
	go func() {
		var i int
		for {
			i++
			workerbus.Gist().SendData(&typs.LayerOne{
				Name:   "send_layer_one_type_LayerOne" + strconv.Itoa(i),
				Amount: 1,
			})
			time.Sleep(time.Microsecond)
		}
	}()
}

func ActionConsumer() {
	// подписчик
	go func() {
		for {
			ch := workerbus.Gist().Subscribe(&typs.General{})
			go consumerGeneral(ch, genInt(100000))
			ch = workerbus.Gist().Subscribe(&typs.LayerTwo{})
			go consumerLayerTwo(ch, genInt(100000))
			time.Sleep(time.Second * 3)
		}
	}()
}

func consumerGeneral(ch chan interface{}, limit int) {
	var i int
	var ok bool
	defer func() {
		if rvr := recover(); rvr != nil {
			// log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = workerbus.Gist().Subscribe(&typs.General{})
			go consumerGeneral(ch, limit)
		}
	}()
	for obj := range ch {
		_, ok = obj.(*typs.General)
		if !ok || limit == i {
			close(ch)
			break
		}
		// It`s Work
		time.Sleep(time.Microsecond)
		// ...
		ch <- true
		i++
	}
	// log.Println("full count: ", i)
}

func consumerLayerTwo(ch chan interface{}, limit int) {
	var i int
	var ok bool
	defer func() {
		if rvr := recover(); rvr != nil {
			// log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = workerbus.Gist().Subscribe(&typs.LayerTwo{})
			go consumerLayerTwo(ch, limit)
		}
	}()
	for obj := range ch {
		_, ok = obj.(*typs.LayerTwo)
		if !ok || limit == i {
			close(ch)
			break
		}
		// It`s Work
		time.Sleep(time.Microsecond)
		// ...
		ch <- true
		i++
	}
	// log.Println("full count: ", i)
}

func genInt(x int64) int {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(x))
	if err != nil {
		panic(err)
	}
	return int(safeNum.Int64())
}
