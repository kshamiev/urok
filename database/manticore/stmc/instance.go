package stmc

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
)

type Instance struct {
	DB   *sql.DB
	Conf *Config
}

func NewInstance(cfg *Config) (*Instance, error) {
	var instance = &Instance{}

	cfgM, err := mysql.ParseDSN(cfg.Dsn)
	if err != nil {
		return nil, err
	}
	cfgM.InterpolateParams = true
	drvM, err := mysql.NewConnector(cfgM)
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(drvM)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	if cfg.DecimalDot > 0 {
		cfg.DotDecimal = decimal.NewFromInt(10).Pow(decimal.NewFromInt(int64(cfg.DecimalDot)))
	} else {
		cfg.DotDecimal = decimal.NewFromInt(100)
	}

	instance.DB = db
	instance.Conf = cfg

	return instance, nil
}
