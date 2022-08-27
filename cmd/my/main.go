package stpg

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var errConnect = errors.New("config is empty or connect is not init")

type Storage struct {
	db *sqlx.DB
}

var instance *Storage

func InitConnect() error {
	var err error
	driverName := "mysql"

	sdb, err := sql.Open(driverName, stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return err
	}
	db := sqlx.NewDb(sdb, driverName)
	if err := db.Ping(); err != nil {
		return err
	}
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	} else {
		db.SetMaxIdleConns(100)
	}
	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	} else {
		db.SetMaxOpenConns(100)
	}

	instance = &Storage{
		db: db,
	}
	return nil
}
