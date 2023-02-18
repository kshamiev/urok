package stbb

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path"

	"go.etcd.io/bbolt"
)

func (self *Instance) DeleteRelation(obj Modeler, indexRel string, ids [][]byte) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketRelation))
		if b == nil {
			return nil
		}
		if len(ids) == 0 {
			return nil
		}
		var err error

		// связь от родителя
		var i int
		index := obj.GetIndex() + ":" + indexRel + ":" + string(obj.GetID()) + ":"
		for i = range ids {
			err = b.Delete([]byte(index + string(ids[i])))
			if err != nil {
				return err
			}
		}

		// связь от потомка
		index = indexRel + ":" + obj.GetIndex() + ":"
		id := string(obj.GetID())
		for i = range ids {
			err = b.Delete([]byte(index + string(ids[i]) + ":" + id))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (self *Instance) LoadRelation(obj Modeler, objSlice Modelers, isLoad bool) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		bRel := tx.Bucket([]byte(bucketRelation))
		if bRel == nil {
			return nil
		}
		bObj := tx.Bucket([]byte(objSlice.GetIndex()))
		if bObj == nil {
			return nil
		}
		var err error

		// связь от родителя
		index := obj.GetIndex() + ":" + objSlice.GetIndex() + ":" + string(obj.GetID()) + ":"
		c := bRel.Cursor()
		for k, v := c.Seek([]byte(index)); k != nil && bytes.HasPrefix(k, []byte(index)); k, v = c.Next() {
			res := bObj.Get(v)
			if res == nil { // Связанный объект был удалён. Удаляем устаревшую связь
				err = bRel.Delete(k)
				if err != nil {
					return err
				}
			}
			if len(res) == 0 {
				return errors.New(errEmpty + objSlice.GetIndex() + "/" + string(v))
			}
			if isLoad {
				objSlice.ParseObject(v, res)
			} else {
				objSlice.ParseIds(v)
			}
		}
		return nil
	})
}

func (self *Instance) SaveRelation(obj Modeler, indexRel string, ids [][]byte) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketRelation))
		if err != nil {
			return err
		}
		if len(ids) == 0 {
			return nil
		}

		// связь от родителя
		var i int
		index := obj.GetIndex() + ":" + indexRel + ":" + string(obj.GetID()) + ":"
		for i = range ids {
			err = b.Put([]byte(index+string(ids[i])), ids[i])
			if err != nil {
				return err
			}
		}

		// связь от потомка
		for i = range ids {
			index = indexRel + ":" + obj.GetIndex() + ":" + string(ids[i]) + ":" + string(obj.GetID())
			err = b.Put([]byte(index), obj.GetID())
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// SelectRange Поиск и получение значений по диапазону ключей
// Sample:
// min := "1990-01-01T00:00:00Z"
// max := "2000-01-01T00:00:00Z"
func (self *Instance) SelectRange(objSlice Modelers, min, max string, isLoad bool) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(objSlice.GetIndex()))
		if b == nil {
			return ErrNotFound
		}

		var k, v []byte
		c := b.Cursor()
		if isLoad {
			for k, v = c.Seek([]byte(min)); k != nil && bytes.Compare(k, []byte(max)) < 0; k, v = c.Next() {
				objSlice.ParseObject(k, v)
			}
		} else {
			for k, _ = c.Seek([]byte(min)); k != nil && bytes.Compare(k, []byte(max)) < 0; k, _ = c.Next() {
				objSlice.ParseIds(k)
			}
		}

		return nil
	})
}

// SelectPrefix Поиск и получение значений по префиксу ключа
// Sample:
// prefix := []byte("1234")
func (self *Instance) SelectPrefix(objSlice Modelers, prefix string, isLoad bool) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(objSlice.GetIndex()))
		if b == nil {
			return ErrNotFound
		}

		var k, v []byte
		c := b.Cursor()
		if isLoad {
			for k, v = c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, []byte(prefix)); k, v = c.Next() {
				objSlice.ParseObject(k, v)
			}
		} else {
			for k, _ = c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, []byte(prefix)); k, _ = c.Next() {
				objSlice.ParseIds(k)
			}
		}

		return nil
	})
}

// Select Получение всех элементов
func (self *Instance) Select(objSlice Modelers, isLoad bool) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(objSlice.GetIndex()))
		if b == nil {
			return ErrNotFound
		}

		var k, v []byte
		c := b.Cursor()
		if isLoad {
			for k, v = c.First(); k != nil; k, v = c.Next() {
				objSlice.ParseObject(k, v)
			}
		} else {
			for k, _ = c.First(); k != nil; k, _ = c.Next() {
				objSlice.ParseIds(k)
			}
		}

		return nil
	})
}

func (self *Instance) Delete(obj Modeler) error {
	return self.DeleteByIndex(obj, string(obj.GetID()))
}

func (self *Instance) DeleteByIndex(obj Modeler, index string) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(obj.GetIndex()))
		if b == nil {
			return nil
		}
		err := b.Delete([]byte(index))
		if err != nil {
			return err
		}
		return nil
	})
}

func (self *Instance) Load(obj Modeler) error {
	return self.LoadByIndex(obj, string(obj.GetID()))
}

func (self *Instance) LoadByIndex(obj Modeler, index string) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(obj.GetIndex()))
		if b == nil {
			return ErrNotFound
		}
		res := b.Get([]byte(index))
		if res == nil {
			return ErrNotFound
		}
		if len(res) == 0 {
			return errors.New(errEmpty + obj.GetIndex() + "/" + index)
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
		if bytes.Equal(Itob(0), obj.GetID()) {
			id, _ := b.NextSequence()
			obj.SetID(Itob(id))
		}
		buf, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		err = b.Put(obj.GetID(), buf)
		if err != nil {
			return err
		}
		return nil
	})
}

func (self *Instance) SaveByIndex(obj Modeler, index string) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(obj.GetIndex()))
		if err != nil {
			return err
		}
		buf, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		err = b.Put([]byte(index), buf)
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
