package redis

import (
	"fmt"
)

type Redis struct {
	Comment string
}

func (rec *Redis) Name() {
	fmt.Println("Database.Name: " + rec.Comment)
}

func (rec *Redis) IsIndex(index string) bool {
	fmt.Println("Redis.IsIndex")
	return index == "index"
}

func (rec *Redis) Members(index string, value string) bool {
	fmt.Println("Redis.Members")
	return index == value
}

func (rec *Redis) Set(index string, value interface{}) error {
	fmt.Println("Redis.Set")
	return nil
}
