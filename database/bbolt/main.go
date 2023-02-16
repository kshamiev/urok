package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shopspring/decimal"
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

	it := &bb.Item{
		Name:     "Фикус губоцветный",
		Price:    decimal.NewFromFloat(34.76),
		CreateAt: time.Now(),
	}
	err = bb.CreateItem(db, it)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")
}
