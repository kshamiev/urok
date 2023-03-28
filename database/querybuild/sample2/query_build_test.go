package main

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/assert"
)

var qu = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func TestUpdate(t *testing.T) {
	sql, args, err := qu.
		Update("orders").
		Set("productcount", sq.Expr("products.productcount")).
		Set("price", sq.Expr("products.price")).
		From("customers, products").
		Where(sq.Expr("orders.customerid = customers.id")).
		Where(sq.Eq{"customers.firstname": "Sam"}).
		Where(sq.Expr("orders.productid = products.id")).
		Where(sq.Eq{"products.company": "HTC"}).
		ToSql()
	assert.NoError(t, err)
	expectedSql := "UPDATE orders " +
		"SET productcount = products.productcount, price = products.price " +
		"FROM customers, products " +
		"WHERE " +
		"orders.customerid = customers.id AND customers.firstname = $1 " +
		"AND orders.productid = products.id AND products.company = $2"
	assert.Equal(t, expectedSql, sql)
	expectedArgs := []interface{}{"Sam", "HTC"}
	assert.Equal(t, expectedArgs, args)

	sql, args, err = qu.
		Update("orders").
		Set("productcount", 100).
		Set("price", 100.100).
		From("customers, products").
		Where(sq.Expr("orders.customerid = customers.id")).
		Where(sq.Eq{"customers.firstname": "Sam"}).
		Where(sq.Expr("orders.productid = products.id")).
		Where(sq.Eq{"products.company": "HTC"}).
		ToSql()
	assert.NoError(t, err)
	expectedSql = "UPDATE orders " +
		"SET productcount = $1, price = $2 " +
		"FROM customers, products " +
		"WHERE " +
		"orders.customerid = customers.id AND customers.firstname = $3 " +
		"AND orders.productid = products.id AND products.company = $4"
	assert.Equal(t, expectedSql, sql)
	expectedArgs = []interface{}{100, 100.1, "Sam", "HTC"}
	assert.Equal(t, expectedArgs, args)

	sql, _, err = qu.
		Update("accounts").
		Set(
			"(contact_first_name, contact_last_name)",
			sq.Select("first_name", "last_name").
				From("salesmen").
				Where("salesmen.id = accounts.sales_id"),
		).
		ToSql()
	assert.NoError(t, err)
	expectedSql = "UPDATE accounts SET (contact_first_name, contact_last_name) = " +
		"(SELECT first_name, last_name FROM salesmen " +
		"WHERE salesmen.id = accounts.sales_id)"
	assert.Equal(t, expectedSql, sql)

	sql, _, err = qu.
		Update("summary s").
		Set(
			"(sum_x, sum_y, avg_x, avg_y)",
			sq.Select("sum(x)", "sum(y)", "avg(x)", "avg(y)").
				From("data d").
				Where("d.group_id = s.group_id"),
		).
		ToSql()
	assert.NoError(t, err)
	expectedSql = "UPDATE summary s SET (sum_x, sum_y, avg_x, avg_y) = " +
		"(SELECT sum(x), sum(y), avg(x), avg(y) FROM data d " +
		"WHERE d.group_id = s.group_id)"
	assert.Equal(t, expectedSql, sql)
}
