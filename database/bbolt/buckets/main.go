package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
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

	it := &Item{
		Name:     "Фикус губоцветный",
		Price:    decimal.NewFromFloat(34.76),
		CreateAt: time.Now(),
	}
	err = Buckets(db, 230, it)
	if err != nil {
		log.Fatal(err)
	}
}

// Buckets Работа с вложенными (зависимыми) buckets "таблицами"
func Buckets(db *bolt.DB, accountID uint64, it *Item) error {
	return db.Update(func(tx *bolt.Tx) error {

		bName := strconv.FormatUint(accountID, 10)
		// Создание и удаление "таблицы"
		bRoot, err := tx.CreateBucketIfNotExists([]byte(bName))
		// bRoot, err := tx.CreateBucket([]byte(bName))
		if err != nil {
			return err
		}
		// defer tx.DeleteBucket([]byte(bName))
		bRoot = tx.Bucket([]byte(bName))

		bkt, err := bRoot.CreateBucketIfNotExists([]byte("ITEMS"))
		if err != nil {
			return err
		}

		it.ID, err = bkt.NextSequence()
		if err != nil {
			return err
		}

		if buf, err := json.Marshal(it); err != nil {
			return err
		} else if err := bkt.Put(itob(it.ID), buf); err != nil {
			return err
		}

		err = bRoot.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

type Item struct {
	ID       uint64
	Name     string
	Price    decimal.Decimal
	CreateAt time.Time
}
