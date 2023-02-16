package main

import (
	"bytes"
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
	return db.Update(func(tx *bolt.Tx) error {

		// Создание и удаление "таблицы"
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		// b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}
		// defer tx.DeleteBucket([]byte("MyBucket"))
		b = tx.Bucket([]byte("MyBucket"))

		// Полный перебор всех значений по порядку
		// variant 1
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%d, value=%s\n", btoi(k), v)
		}
		// variant 2
		err = b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%d, value=%s\n", btoi(k), v)
			return nil
		})
		if err != nil {
			return err
		}

		// Поиск и получение значений по префиксу ключа
		prefix := []byte("1234")
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		// Поиск и получение значений по диапазону ключей
		min := []byte("1990-01-01T00:00:00Z")
		max := []byte("2000-01-01T00:00:00Z")
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}

		return nil
	})
}

// btoi returns an 8-byte big endian representation of v.
func btoi(v []byte) uint64 {
	return binary.BigEndian.Uint64(v)
}
