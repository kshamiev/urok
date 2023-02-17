package stbb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"go.etcd.io/bbolt"
)

// SelectRange Поиск и получение значений по диапазону ключей
// Sample:
// min := []byte("1990-01-01T00:00:00Z")
// max := []byte("2000-01-01T00:00:00Z")
func (self *Instance) SelectRange(objSlice Modelers, min, max []byte) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(objSlice.GetIndex()))
		if b == nil {
			return ErrNotFound
		}

		c := b.Cursor()
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) < 0; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}

		return nil
	})
}

// SelectPrefix Поиск и получение значений по префиксу ключа
// Sample:
// prefix := []byte("1234")
func (self *Instance) SelectPrefix(objSlice Modelers, prefix []byte) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(objSlice.GetIndex()))
		if b == nil {
			return ErrNotFound
		}

		c := b.Cursor()
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			objSlice.ParseByte(k, v)
		}

		return nil
	})
}

// Select Получение всех элементов
func (self *Instance) Select(objSlice Modelers) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(objSlice.GetIndex()))
		if b == nil {
			return ErrNotFound
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			objSlice.ParseByte(k, v)
		}

		return nil
	})
}

func (self *Instance) Delete(obj Modeler) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(obj.GetIndex()))
		if b == nil {
			return nil
		}
		err := b.Delete(obj.GetID())
		if err != nil {
			return err
		}
		return nil
	})
}

func (self *Instance) Load(obj Modeler) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(obj.GetIndex()))
		if b == nil {
			return ErrNotFound
		}
		res := b.Get(obj.GetID())
		if string(res) == emptyValue {
			return ErrNotFound
		}
		err := json.Unmarshal(res, obj)
		if err != nil {
			return err
		}
		return nil
	})
}

func (self *Instance) Save(obj Modeler) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(obj.GetIndex()))
		if err != nil {
			return err
		}

		id, _ := b.NextSequence()
		obj.SetID(Itob(id))

		buf, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		err = b.Put(Itob(id), buf)
		if err != nil {
			return err
		}
		return nil
	})
}

func (self *Instance) Close() {
	_ = self.db.Close()
}

func (self *Instance) DB() *bbolt.DB {
	return self.db
}

func NewInstance(cfg *Config) (*Instance, error) {
	err := os.MkdirAll(path.Dir(cfg.PathDB), 0o777)
	if err != nil {
		return nil, err
	}
	inst := &Instance{}
	inst.db, err = bbolt.Open(cfg.PathDB, cfg.FileDBMode, &bbolt.Options{
		Timeout:         cfg.Timeout,
		NoGrowSync:      cfg.NoGrowSync,
		NoFreelistSync:  cfg.NoFreelistSync,
		PreLoadFreelist: cfg.PreLoadFreelist,
		FreelistType:    cfg.FreelistType,
		ReadOnly:        cfg.ReadOnly,
		MmapFlags:       cfg.MmapFlags,
		InitialMmapSize: cfg.InitialMmapSize,
		PageSize:        cfg.PageSize,
		NoSync:          cfg.NoSync,
		Mlock:           cfg.Mlock,
	})
	if err != nil {
		return nil, err
	}
	return inst, nil
}
