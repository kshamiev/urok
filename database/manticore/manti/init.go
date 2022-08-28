package manti

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

var instance *Storage

func Init() error {
	var err error
	driverName := "mysql"
	sdb, err := sql.Open(driverName, "tcp(127.0.0.1:9306)/")
	if err != nil {
		return err
	}
	db := sqlx.NewDb(sdb, driverName)
	if err := db.Ping(); err != nil {
		return err
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	instance = &Storage{
		db: db,
	}
	return nil
}
