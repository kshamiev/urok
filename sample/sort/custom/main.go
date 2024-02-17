package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"

	"github.com/kshamiev/urok/debug"
)

func main() {
	data := generate()
	debug.Dumper(data)
	fmt.Println("\n\n===================\n\n")
	data = sort(data)
	debug.Dumper(data)
}

func sort(data []Good) []Good {
	if len(data) == 0 {
		return data
	}

	dataNew := make([]Good, 0, len(data))
	id := data[0].ID
	cid := data[0].CID
	temp := make([]Good, 0, 100)
	var i int

	for i = range data {
		// nex block
		if id != data[i].ID || cid != data[i].CID {
			id = data[i].ID
			cid = data[i].CID
			dataNew = append(dataNew, temp...)
			temp = make([]Good, 0, 100)
		}
		if data[i].Number1 == data[i].Number2 { // valid
			dataNew = append(dataNew, data[i])
		} else { // not valid
			temp = append(temp, data[i])
		}
	}
	if len(temp) > 0 { // оставшийся хвост
		dataNew = append(dataNew, temp...)
	}
	return dataNew
}

type Good struct {
	ID      int
	CID     int
	Number1 int
	Number2 int
}

func generate() []Good {
	res := make([]Good, 0, 100)
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for n := 0; n < 10; n++ {
				res = append(res, Good{
					ID:      i,
					CID:     j,
					Number1: genInt(1, 9),
					Number2: genInt(1, 9),
				})
			}
		}
	}
	return res
}

func genInt(min, max int) int {
	safeNum, err := crand.Int(crand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}
	return min + int(safeNum.Int64())
}
