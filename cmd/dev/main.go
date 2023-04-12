package main

// https://stackoverflow.com/questions/71418671/restart-or-shutdown-golang-apps-programmatically

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
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
	qu1 := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	q, arg, err := qu1.Select("*").
		From("articles").
		Where(sq.Eq{
			"avtor_id":    args1,
			"redaktor_id": args2,
		}).
		ToSql()
	fmt.Println(q, arg, err)
	// SELECT * FROM articles WHERE avtor_id IN ($1,$2,$3,$4,$5) AND redaktor_id IN ($6,$7,$8,$9,$10)
	// [1 2 3 4 5 11 12 13 14 15]
	// <nil>

}
