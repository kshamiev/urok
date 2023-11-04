package database

import (
	"fmt"
)

type Database struct {
	Comment string
}

func (rec *Database) Name() {
	fmt.Println("Database.Name: " + rec.Comment)
}

func (rec *Database) Get() string {
	fmt.Println("Database.Get")
	return "sample get"

}

func (rec *Database) Select() []string {
	fmt.Println("Database.Select")
	return []string{"one", "two"}
}

func (rec *Database) Insert(s string) {
	fmt.Println("Database.Insert")
}

func (rec *Database) Update(id int64, s string) {
	fmt.Println("Database.Update")
}
