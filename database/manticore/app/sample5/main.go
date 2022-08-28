package main

import (
	"log"

	"github.com/kshamiev/urok/database/manticore/manti"
)

func main() {

	if err := manti.Init(); err != nil {
		log.Fatalln(err)
	}

}
