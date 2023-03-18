package stmc

import (
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

func ConvertString(v interface{}) string {
	if v == nil {
		return ""
	}
	return string(v.([]byte))
}

func ConvertBigint(v interface{}) int64 {
	if v == nil {
		return 0
	}
	i, _ := strconv.ParseInt(string(v.([]byte)), 10, 64)
	return i
}

func ConvertUint(v interface{}) uint32 {
	if v == nil {
		return 0
	}
	i, _ := strconv.ParseUint(string(v.([]byte)), 10, 32)
	return uint32(i)
}

func ConvertDecimal(v interface{}) decimal.Decimal {
	if v == nil {
		return decimal.Decimal{}
	}
	d, _ := decimal.NewFromString(string(v.([]byte)))
	return d
}

func ConvertTime(v interface{}) time.Time {
	if v == nil {
		return time.Time{}
	}
	i, _ := strconv.ParseInt(string(v.([]byte)), 10, 64)
	return time.Unix(i, 0)
}

func ConvertBool(v interface{}) bool {
	if v == nil {
		return false
	}
	return string(v.([]byte)) == "1"
}
