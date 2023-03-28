package main

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/kshamiev/urok/debug"
)

func main() {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.
		Select("*").
		From("users").
		Join("emails USING (email_id)").
		Where(sq.Eq{"deleted_at": nil}).
		Where(sq.Like{"name": "%QWERTY%"}).
		Where(sq.Eq{"username": []string{"moe", "larry", "curly", "shemp"}}).
		Where(sq.Eq{"age": 45}).
		Where(sq.Or{
			sq.Eq{"col1": 1, "col2": 2},
			sq.Eq{"col1": 3, "col2": 4},
		}).
		Where(sq.And{
			sq.Eq{"col1": 1, "col2": 2},
			sq.Eq{"col1": 3, "col2": 4},
		}).
		Limit(3).
		ToSql()
	debug.Dumper(sql, args, err)

	sql, args, err = sq.
		Insert("users").
		Columns("name", "age").
		Values("moe", 13).
		Values("larry", sq.Expr("? + 5", 12)).
		Values("larry", sq.Expr("FROM_UNIXTIME(?)", 12)).
		Suffix(`RETURNING "id"`).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	// "INSERT INTO users (name,age) VALUES (?,?),(?,? + 5)"
	debug.Dumper(sql, args, err)
}
