package main

import (
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

	err = TransactionUpdate(db)
	if err != nil {
		log.Fatal(err)
	}
}

// TransactionUpdate транзакция на чтение и запись
func TransactionUpdate(db *bolt.DB) error {
	var v, v1 []byte
	err := db.Update(func(tx *bolt.Tx) error {

		// Создание и удаление "таблицы"
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		// b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}
		// defer tx.DeleteBucket([]byte("MyBucket"))
		b = tx.Bucket([]byte("MyBucket"))

		// Сохранение данных
		err = b.Put([]byte("answer"), []byte("42"))
		if err != nil {
			return err
		}
		err = b.Put([]byte("answerZeroValue"), []byte(null))
		if err != nil {
			return err
		}

		// Получение данных
		v = b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		v1 = b.Get([]byte("answerZeroValue"))
		fmt.Printf("The answerZeroValue is: %s\n", v1)

		// Удаление данных
		err = b.Delete([]byte("answer"))
		if err != nil {
			return err
		}

		if b.Get([]byte("answer")) == nil {
			fmt.Printf("The answer is delete\n")
		}

		return nil
	})
	fmt.Println(string(v), string(v1))
	return err
}

// TransactionView транзакция на чтение
func TransactionView(db *bolt.DB) error {
	err := db.View(func(tx *bolt.Tx) error {
		//
		return nil
	})
	return err
}

// TransactionBatch транзакция на конкурентную запись
func TransactionBatch(db *bolt.DB) error {
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
