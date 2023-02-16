package main

import (
	"encoding/binary"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

// Используется для определения пустых значений
// Так как несуществующие и пустые значения вернут одно и тоже 0 байтов
const null = "nil"

func main() {
	db, err := bolt.Open("database/bbolt/data.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = Cursor(db)
	if err != nil {
		log.Fatal(err)
	}
}

// Cursor Итерация данных с помощью курсора
func Cursor(db *bolt.DB) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%d, value=%s\n", btoi(k), v)
		}

		err := b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%d, value=%s\n", btoi(k), v)
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})
}

// btoi returns an 8-byte big endian representation of v.
func btoi(v []byte) int64 {
	return int64(binary.BigEndian.Uint64(v))
}
