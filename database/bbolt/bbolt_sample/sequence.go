package bbolt_sample

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	bolt "go.etcd.io/bbolt"
)

// CreateItem saves u to the store. The new item ID is set on u once the data is persisted.
func CreateItem(db *bolt.DB, item1 *Item) error {
	return db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return err
		}
		defer tx.DeleteBucket([]byte("MyBucket"))

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
