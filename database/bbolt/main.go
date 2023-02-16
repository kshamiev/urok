package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"

	bb "github.com/kshamiev/urok/database/bbolt/bbolt_sample"
)

var dbPath = "/home/konstantin/work/urok/database/bbolt/data.db"

func main() {
	db, err := bolt.Open(dbPath, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = bb.TransactionUpdate(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")
}
