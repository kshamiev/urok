package main

// https://stackoverflow.com/questions/71418671/restart-or-shutdown-golang-apps-programmatically

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func main() {
	qu := "select * from articles where avtor_id IN (?) and avtor_id IN (?) and status = true"
	args1 := []int64{1, 2, 3, 4, 5}
	args2 := []int64{11, 12, 13, 14, 15}
	fmt.Println(sqlx.In(qu, args1, args2))
	// output:
	// select * from articles where avtor_id IN (?, ?, ?, ?, ?) and avtor_id IN (?, ?, ?, ?, ?) and status = true
	// [1 2 3 4 5 11 12 13 14 15] <nil>
	// <nil>
}
