package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

var dbPath = "/home/konstantin/work/urok/database/bbolt/data.db"

func main() {
	db, err := bolt.Open(dbPath, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = transactionUpdate(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")
}

const null = "nil"

// транзакция на чтение и запись
func transactionUpdate(db *bolt.DB) error {
	var v, v1 []byte

	err := db.Update(func(tx *bolt.Tx) error {

		// b, err := tx.CreateBucket([]byte("MyBucket"))
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return err
		}
		defer tx.DeleteBucket([]byte("MyBucket"))

		err = b.Put([]byte("answer"), []byte("42"))
		if err != nil {
			return err
		}
		err = b.Put([]byte("answerZeroValue"), []byte(null))
		if err != nil {
			return err
		}
		v = b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		v1 = b.Get([]byte("answerZeroValue"))
		fmt.Printf("The answerZeroValue is: %s\n", v1)

		return nil
	})

	fmt.Println(string(v), string(v1))

	return err
}

// транзакция на чтение
func transactionView(db *bolt.DB) error {
	err := db.View(func(tx *bolt.Tx) error {
		//
		return nil
	})
	return err
}

// транзакция на конкурентную запись
func transactionBatch(db *bolt.DB) error {
	var id uint64
	// Пакетная обработка полезна только тогда, когда ее вызывает несколько горутин.
	// Подходит для генерации идентификатора
	err := db.Batch(func(tx *bolt.Tx) error {
		// Find last key in bucket, decode as bigendian uint64, increment
		// by one, encode back to []byte, and add new key.
		// id = newValue
		return nil
	})
	if err != nil {
		return err
	}
	fmt.Println("Allocated ID %d", id)
	return nil
}
