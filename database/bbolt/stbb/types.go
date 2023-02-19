package stbb

import (
	"errors"
	"os"
	"time"

	"go.etcd.io/bbolt"
)

const (
	bucketRelation = "_RELATION"
	errEmpty       = "Мы нашли ведро либо антиматерию: "
	errNil         = "Целостность данных нарушена: "
)

var ErrNotFound = errors.New("Не найдено")

type Modeler interface {
	GetBID() []byte   // get object ID
	SetBID([]byte)    // set object ID
	GetIndex() string // index (manticore), bucket (bolt, minio), table (postgres, mysql ...)
}

type Modelers interface {
	ParseObject(i int, value []byte) // Обработка одного элемента
	GetIndex() string                // index (manticore), bucket (bolt, minio), table (postgres, mysql ...)
	GetIds() [][]byte                // Получение идентификаторов (для работы со связями в bbolt)
}

// ////

type Config struct {
	PathDB string `yaml:"pathDB"`

	FileDBMode os.FileMode `yaml:"fileDBMode"`

	// Timeout is the amount of time to wait to obtain a file lock.
	// When set to zero it will wait indefinitely. This option is only
	// available on Darwin and Linux.
	Timeout time.Duration `yaml:"timeout"`

	// Sets the DB.NoGrowSync flag before memory mapping the file.
	NoGrowSync bool `yaml:"noGrowSync"`

	// Do not sync freelist to disk. This improves the database write performance
	// under normal operation, but requires a full database re-sync during recovery.
	NoFreelistSync bool `yaml:"noFreelistSync"`

	// PreLoadFreelist sets whether to load the free pages when opening
	// the db file. Note when opening db in write mode, bbolt will always
	// load the free pages.
	PreLoadFreelist bool `yaml:"preLoadFreelist"`

	// FreelistType sets the backend freelist type. There are two options. Array which is simple but endures
	// dramatic performance degradation if database is large and fragmentation in freelist is common.
	// The alternative one is using hashmap, it is faster in almost all circumstances
	// but it doesn't guarantee that it offers the smallest page id available. In normal case it is safe.
	// The default type is array
	FreelistType bbolt.FreelistType `yaml:"freelistType"`

	// Open database in read-only mode. Uses flock(..., LOCK_SH |LOCK_NB) to
	// grab a shared lock (UNIX).
	ReadOnly bool `yaml:"readOnly"`

	// Sets the DB.MmapFlags flag before memory mapping the file.
	MmapFlags int `yaml:"mmapFlags"`

	// InitialMmapSize is the initial mmap size of the database
	// in bytes. Read transactions won't block write transaction
	// if the InitialMmapSize is large enough to hold database mmap
	// size. (See DB.Begin for more information)
	//
	// If <=0, the initial map size is 0.
	// If initialMmapSize is smaller than the previous database size,
	// it takes no effect.
	InitialMmapSize int `yaml:"initialMmapSize"`

	// PageSize overrides the default OS page size.
	PageSize int `yaml:"pageSize"`

	// NoSync sets the initial value of DB.NoSync. Normally this can just be
	// set directly on the DB itself when returned from Open(), but this option
	// is useful in APIs which expose Options but not the underlying DB.
	NoSync bool `yaml:"noSync"`

	// Mlock locks database file in memory when set to true.
	// It prevents potential page faults, however
	// used memory can't be reclaimed. (UNIX only)
	Mlock bool `yaml:"mlock"`
}

type Instance struct {
	db *bbolt.DB
}
