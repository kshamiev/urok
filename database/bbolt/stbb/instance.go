package stbb

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path"
	"time"

	"go.etcd.io/bbolt"
)

func (self *Instance) DeleteRelation(obj Modeler, objSlice Modelers) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		bRel := tx.Bucket([]byte(bucketRelation))
		if bRel == nil {
			return nil
		}
		ids := objSlice.GetIds()
		if len(ids) == 0 {
			return nil
		}
		var err error

		// связь от родителя
		var i int
		idx := obj.GetIndex() + ":" + objSlice.GetIndex() + ":" + string(obj.GetBID()) + ":"
		for i = range ids {
			err = bRel.Delete([]byte(idx + string(ids[i])))
			if err != nil {
				return err
			}
		}

		// связь от потомка
		idx = objSlice.GetIndex() + ":" + obj.GetIndex() + ":"
		id := string(obj.GetBID())
		for i = range ids {
			err = bRel.Delete([]byte(idx + string(ids[i]) + ":" + id))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (self *Instance) LoadRelation(obj Modeler, objSlice Modelers) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		bRel := tx.Bucket([]byte(bucketRelation))
		if bRel == nil {
			return nil
		}
		bObj := tx.Bucket([]byte(objSlice.GetIndex()))
		if bObj == nil {
			return nil
		}

		// связь от родителя
		var res, k, v []byte
		idx := obj.GetIndex() + ":" + objSlice.GetIndex() + ":" + string(obj.GetBID()) + ":"
		c := bRel.Cursor()
		for k, v = c.Seek([]byte(idx)); k != nil && bytes.HasPrefix(k, []byte(idx)); k, v = c.Next() {
			res = bObj.Get(v)
			if res == nil {
				return errors.New(errNil + string(k))
			}
			if len(res) == 0 {
				return errors.New(errEmpty + objSlice.GetIndex() + "/" + string(v))
			}
			objSlice.ParseObject(-1, res)
		}

		return nil
	})
}

func (self *Instance) SaveRelation(obj Modeler, objSlice Modelers) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		bRel, err := tx.CreateBucketIfNotExists([]byte(bucketRelation))
		if err != nil {
			return err
		}
		ids := objSlice.GetIds()
		if len(ids) == 0 {
			return nil
		}

		// связь от родителя
		var i int
		idx := obj.GetIndex() + ":" + objSlice.GetIndex() + ":" + string(obj.GetBID()) + ":"
		for i = range ids {
			err = bRel.Put([]byte(idx+string(ids[i])), ids[i])
			if err != nil {
				return err
			}
		}

		// связь от потомка
		idx = objSlice.GetIndex() + ":" + obj.GetIndex() + ":"
		id := string(obj.GetBID())
		for i = range ids {
			err = bRel.Put([]byte(idx+string(ids[i])+":"+id), obj.GetBID())
			if err != nil {
				return err
			}
		}

		// связь с хранилищем
		err = bRel.Put([]byte(obj.GetIndex()+":"+objSlice.GetIndex()), []byte(objSlice.GetIndex()))
		if err != nil {
			return err
		}

		return nil
	})
}

// SelectRange Поиск и получение значений по диапазону ключей
// Sample:
// min := "1990-01-01T00:00:00Z"
// max := "2000-01-01T00:00:00Z"
func (self *Instance) SelectRange(objSlice Modelers, min, max string) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		bObj := tx.Bucket([]byte(objSlice.GetIndex()))
		if bObj == nil {
			return ErrNotFound
		}

		var k, v []byte
		c := bObj.Cursor()
		for k, v = c.Seek([]byte(min)); k != nil && bytes.Compare(k, []byte(max)) < 0; k, v = c.Next() {
			objSlice.ParseObject(-1, v)
		}

		return nil
	})
}

// SelectPrefix Поиск и получение значений по префиксу ключа
// Sample:
// prefix := []byte("1234")
func (self *Instance) SelectPrefix(objSlice Modelers, prefix string) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		bObj := tx.Bucket([]byte(objSlice.GetIndex()))
		if bObj == nil {
			return ErrNotFound
		}

		var k, v []byte
		c := bObj.Cursor()
		for k, v = c.Seek([]byte(prefix)); k != nil && bytes.HasPrefix(k, []byte(prefix)); k, v = c.Next() {
			objSlice.ParseObject(-1, v)
		}

		return nil
	})
}

// Select Получение всех элементов
func (self *Instance) Select(objSlice Modelers) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		bObj := tx.Bucket([]byte(objSlice.GetIndex()))
		if bObj == nil {
			return nil
		}

		var k, v []byte
		c := bObj.Cursor()
		for k, v = c.First(); k != nil; k, v = c.Next() {
			objSlice.ParseObject(-1, v)
		}

		return nil
	})
}

func (self *Instance) Delete(obj Modeler) error {
	return self.DeleteByID(obj, string(obj.GetBID()))
}

func (self *Instance) DeleteByID(obj Modeler, id string) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		bObj := tx.Bucket([]byte(obj.GetIndex()))
		if bObj == nil {
			return nil
		}
		err := bObj.Delete([]byte(id))
		if err != nil {
			return err
		}

		bRel := tx.Bucket([]byte(bucketRelation))
		if bRel == nil {
			return nil
		}

		var k, k1, v, v1 []byte
		var idxP, idxC string
		idx := obj.GetIndex() + ":"
		c := bRel.Cursor()
		cc := bRel.Cursor()
		for k, v = c.Seek([]byte(idx)); k != nil && bytes.HasPrefix(k, []byte(idx)); k, v = c.Next() {
			//
			idxP = obj.GetIndex() + ":" + string(v) + ":" + id + ":"
			idxC = string(v) + ":" + obj.GetIndex() + ":"
			for k1, v1 = cc.Seek([]byte(idxP)); k1 != nil && bytes.HasPrefix(k1, []byte(idxP)); k1, v1 = c.Next() {
				// связь от родителя
				err = bRel.Delete(k1)
				if err != nil {
					return err
				}
				// связь от потомка
				err = bRel.Delete([]byte(idxC + string(v1) + ":" + id))
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (self *Instance) Load(obj Modeler) error {
	return self.LoadByID(obj, string(obj.GetBID()))
}

func (self *Instance) LoadByID(obj Modeler, id string) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		bObj := tx.Bucket([]byte(obj.GetIndex()))
		if bObj == nil {
			return ErrNotFound
		}
		res := bObj.Get([]byte(id))
		if res == nil {
			return ErrNotFound
		}
		if len(res) == 0 {
			return errors.New(errEmpty + obj.GetIndex() + "/" + id)
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
		bObj, err := tx.CreateBucketIfNotExists([]byte(obj.GetIndex()))
		if err != nil {
			return err
		}
		if bytes.Equal(Itob(0), obj.GetBID()) {
			id, _ := bObj.NextSequence()
			obj.SetBID(Itob(id))
		}
		buf, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		err = bObj.Put(obj.GetBID(), buf)
		if err != nil {
			return err
		}
		return nil
	})
}

func (self *Instance) SaveByID(obj Modeler, id string) error {
	return self.db.Update(func(tx *bbolt.Tx) error {
		bObj, err := tx.CreateBucketIfNotExists([]byte(obj.GetIndex()))
		if err != nil {
			return err
		}
		buf, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		err = bObj.Put([]byte(id), buf)
		if err != nil {
			return err
		}
		return nil
	})
}

func (self *Instance) Stats() {
	// Grab the initial stats.
	prev := self.db.Stats()
	for {
		// Wait for 10s.
		time.Sleep(10 * time.Second)

		// Grab the current stats and diff them.
		stats := self.db.Stats()
		diff := stats.Sub(&prev)

		// Encode stats to JSON and print to STDERR.
		json.NewEncoder(os.Stderr).Encode(diff)

		// Save stats for the next loop.
		prev = stats
	}
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
