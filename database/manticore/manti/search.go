package manti

import (
	"context"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type Parser interface {
	Parse(map[string]interface{})
}

type Search struct {
	sel    []string
	limit  int
	offset int
}

func NewSearch(sel ...string) Search {
	return Search{
		sel: sel,
	}
}

func (s *Search) Limit(limit, offset int) {
	s.limit = limit
	s.offset = offset
}

func SearchCustom(ctx context.Context, result Parser, qu string, args ...interface{}) error {
	rows, err := instance.DB.QueryContext(ctx, qu, args...)
	if err != nil {
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	values := make([]interface{}, len(cols))
	for i := range values {
		values[i] = new(interface{})
	}

	dest := make(map[string]interface{})
	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			return err
		}
		for i, column := range cols {
			dest[column] = *(values[i].(*interface{}))
		}
		if err := rows.Err(); err != nil {
			return err
		}
		result.Parse(dest)
	}
	return nil
}

func ConvertString(v interface{}) string {
	return string(v.([]byte))
}
func ConvertBigint(v interface{}) int64 {
	i, _ := strconv.ParseInt(string(v.([]byte)), 10, 64)
	return i
}
func ConvertUint(v interface{}) uint {
	i, _ := strconv.ParseUint(string(v.([]byte)), 10, 32)
	return uint(i)
}
func ConvertFloat(v interface{}) decimal.Decimal {
	d, _ := decimal.NewFromString(string(v.([]byte)))
	return d
}
func ConvertTime(v interface{}) time.Time {
	i, _ := strconv.ParseInt(string(v.([]byte)), 10, 64)
	return time.Unix(i, 0)
}
func ConvertBool(v interface{}) bool {
	if "1" == string(v.([]byte)) {
		return true
	}
	return false
}
