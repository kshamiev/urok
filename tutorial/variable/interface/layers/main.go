package main

import (
	"github.com/kshamiev/urok/tutorial/variable/interface/layers/database"
	"github.com/kshamiev/urok/tutorial/variable/interface/layers/face"
	"github.com/kshamiev/urok/tutorial/variable/interface/layers/redis"
)

func main() {

}

type Model struct {
	RD face.RD
	DB face.DB
}

func NewModel() *Model {
	return &Model{
		RD: &redis.Redis{},
		DB: &database.Database{},
	}
}

func (rec *Model) Test() {
	rec.RD.Name()
}
