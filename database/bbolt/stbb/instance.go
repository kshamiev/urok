package stbb

import (
	"encoding/json"
	"os"
	"path"

	"go.etcd.io/bbolt"
)

func (self *Instance) Select(objSlice Modelers) error {
	return self.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(objSlice.GetIndex()))
		if b == nil {
			return ErrNotFound
		}

		var err error
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			err = objSlice.ParseByte(k, v)
			if err != nil {
				return err
			}
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
		err := b.Delete(Itob(obj.GetID()))
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
		res := b.Get(Itob(obj.GetID()))
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
		obj.SetID(id)

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
