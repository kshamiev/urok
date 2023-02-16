package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
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
	err = CreateItem(db, it)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateItem saves u to the store. The new item ID is set on u once the data is persisted.
func CreateItem(db *bolt.DB, item1 *Item) error {
	return db.Update(func(tx *bolt.Tx) error {

		// Создание и удаление "таблицы"
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		// b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}
		// defer tx.DeleteBucket([]byte("MyBucket"))
		b = tx.Bucket([]byte("MyBucket"))

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		item1.ID = int64(id)

		buf, err := json.Marshal(item1)
		if err != nil {
			return err
		}
		err = b.Put(itob(item1.ID), buf)
		if err != nil {
			return err
		}

		res := b.Get(itob(item1.ID))
		us := &Item{}
		err = json.Unmarshal(res, us)
		if err != nil {
			return err
		}
		fmt.Println(us)

		return nil
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

type Item struct {
	ID       int64
	Name     string
	Price    decimal.Decimal
	CreateAt time.Time
}
