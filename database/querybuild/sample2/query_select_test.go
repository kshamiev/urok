package main

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	qu := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, args, err := qu.
		Select("col1", "col2", "count(*)").
		From("users").
		Join("emails USING (email_id)").
		// полное объединение по условию
		InnerJoin("emails ON emails.id = users.email_id").
		// объединение с левой таблицей
		// получение users с ИД 34 для которых нет записей в emails
		LeftJoin("emails ON users.id = emails.user_id AND users.id = ?", 34).
		// объединение с правой таблицей
		// получение emails с ИД 57 для которых нет записей в users
		RightJoin("emails ON users.email_id = emails.id AND emails.id = ?", 57).
		Where(sq.Eq{"deleted_at": nil}).
		Where(sq.Like{"name": "%QWERTY%"}).
		Where(sq.Eq{"username": []string{"moe", "larry", "curly", "shemp"}}).
		Where(sq.Eq{"age": 45}).
		Where(sq.Or{
			sq.Eq{"col1": 1, "col2": 2},
			sq.Eq{"col1": 3, "col2": 4},
		}).
		Where(sq.Eq{"col1": 1, "col2": 2, "col3": 3, "col4": 4}).
		GroupBy("col1", "col2").
		Having("count(*) > 3").
		OrderBy("col1 DESC").
		Limit(30).
		Offset(100).
		ToSql()
	assert.NoError(t, err)
	expectedSql := "SELECT col1, col2, count(*) FROM users JOIN emails USING (email_id) INNER JOIN emails ON emails.id = users.email_id LEFT JOIN emails ON users.id = emails.user_id AND users.id = $1 RIGHT JOIN emails ON users.email_id = emails.id AND emails.id = $2 WHERE deleted_at IS NULL AND name LIKE $3 AND username IN ($4,$5,$6,$7) AND age = $8 AND (col1 = $9 AND col2 = $10 OR col1 = $11 AND col2 = $12) AND col1 = $13 AND col2 = $14 AND col3 = $15 AND col4 = $16 GROUP BY col1, col2 HAVING count(*) > 3 ORDER BY col1 DESC LIMIT 30 OFFSET 100"
	assert.Equal(t, expectedSql, sql)
	expectedArgs := []interface{}{34, 57, "%QWERTY%", "moe", "larry", "curly", "shemp", 45, 1, 2, 3, 4, 1, 2, 3, 4}
	assert.Equal(t, expectedArgs, args)
}
