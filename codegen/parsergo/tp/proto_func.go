// nolint
package tp

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/strmangle"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// type conversion support functions
// вспомогательные функции реализующие уникальную обработку свойств для определяемых рабочих типов

// PbFromNullBytes перевод из примитива grpc в рабочий тип
func PbFromNullBytes(b []byte) null.Bytes {
	return null.BytesFrom(b)
}

// PbFromNullJSON перевод из примитива grpc в рабочий тип
func PbFromNullJSON(b []byte) null.JSON {
	return null.JSONFrom(b)
}

// PbFromNullString перевод из примитива grpc в рабочий
func PbFromNullString(s string) null.String {
	if s == "" {
		return null.String{}
	}
	return null.StringFrom(s)
}

// PbFromNullInt64 перевод из примитива grpc в рабочий
func PbFromNullInt64(n int64) null.Int64 {
	if n == 0 {
		return null.Int64{}
	}
	return null.NewInt64(n, true)
}

// PbFromNullFloat64 перевод из примитива grpc в рабочий
func PbFromNullFloat64(n float64) null.Float64 {
	if n == 0 {
		return null.Float64{}
	}
	return null.NewFloat64(n, true)
}

// PbFromNullInt32 перевод из примитива grpc в рабочий
func PbFromNullInt32(n int32) null.Int32 {
	if n == 0 {
		return null.Int32{}
	}
	return null.NewInt32(n, true)
}

// PbFromNullInt16 перевод из примитива grpc в рабочий
func PbFromNullInt16(n int32) null.Int16 {
	if n == 0 {
		return null.Int16{}
	}
	return null.NewInt16(int16(n), true)
}

// PbFromNullInt перевод из примитива grpc в рабочий
func PbFromNullInt(n int64) null.Int {
	if n == 0 {
		return null.Int{}
	}
	return null.NewInt(int(n), true)
}

// PbFromNullBool перевод из примитива grpc в рабочий
func PbFromNullBool(b bool) null.Bool {
	if !b {
		return null.Bool{}
	}
	return null.BoolFrom(b)
}

// PbFromNullTime перевод из примитива grpc в рабочий тип
func PbFromNullTime(d *timestamppb.Timestamp) null.Time {
	return null.TimeFrom(d.AsTime())
}

// PbToTime перевод в примитив grpc из рабочего типа
func PbToTime(d time.Time) *timestamppb.Timestamp {
	if d.IsZero() {
		return &timestamppb.Timestamp{}
	}
	return timestamppb.New(d)
}

// PbFromTime перевод из примитива grpc в рабочий тип
func PbFromTime(d *timestamppb.Timestamp) time.Time {
	if d == nil || (d.Nanos == 0 && d.Seconds == 0) {
		return time.Time{}
	}
	return d.AsTime()
}

// PbFromDecimal перевод из примитива grpc в рабочий тип
func PbFromDecimal(v float64) decimal.Decimal {
	d := decimal.NewFromFloat(v)
	return d
}

// PbFromDecimalArray перевод из примитива grpc в рабочий тип
func PbFromDecimalArray(list []float64) []decimal.Decimal {
	out := make([]decimal.Decimal, len(list))
	for i, item := range list {
		out[i] = decimal.NewFromFloat(item)
	}
	return out
}

// PbToDecimalArray перевод из рабочего типа в примитив grpc
func PbToDecimalArray(list []decimal.Decimal) []float64 {
	out := make([]float64, len(list))
	for i, item := range list {
		out[i], _ = item.Float64()
	}
	return out
}

// PbToNullUUID перевод из примитива grpc в рабочий тип
func PbToNullUUID(value string) uuid.NullUUID {
	if value == "" {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{Valid: true, UUID: uuid.MustParse(value)}
}

// PbFromNullUUID перевод из рабочего типа в примитив grpc
func PbFromNullUUID(value uuid.NullUUID) string {
	return value.UUID.String()
}

func PbToUUIDS(list []uuid.UUID) []string {
	uu := make([]string, len(list))
	for i := range list {
		uu[i] = list[i].String()
	}
	return uu
}

func PbFromUUIDS(list []string) []uuid.UUID {
	uu := make([]uuid.UUID, len(list))
	for i := range list {
		uu[i] = uuid.MustParse(list[i])
	}
	return uu
}

// UPSERT ALL

func upsertAll(tableName string, updateOnConflict bool, cntRow int, insert, conflict, update []string) string {
	buf := strmangle.GetBuffer()
	defer strmangle.PutBuffer(buf)

	columns := "DEFAULT VALUES"
	if len(insert) != 0 {
		columns = fmt.Sprintf("(%s) VALUES %s",
			strings.Join(insert, ", "),
			upsertAllPlaceholders(cntRow, len(insert)),
		)
	}

	_, _ = fmt.Fprintf(
		buf,
		"INSERT INTO %s %s ON CONFLICT ",
		tableName,
		columns,
	)

	if !updateOnConflict || len(update) == 0 {
		buf.WriteString("DO NOTHING")
	} else {
		buf.WriteByte('(')
		buf.WriteString(strings.Join(conflict, ", "))
		buf.WriteString(") DO UPDATE SET ")
		var i int
		var v string
		for i, v = range update {
			if i != 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(v)
			buf.WriteString(" = EXCLUDED.")
			buf.WriteString(v)
		}
	}

	if len(conflict) != 0 {
		buf.WriteString(" RETURNING ")
		buf.WriteString(strings.Join(conflict, ", "))
	}

	return buf.String()
}

func upsertAllPlaceholders(cntRow, cntCol int) string {
	buf := strmangle.GetBuffer()
	defer strmangle.PutBuffer(buf)

	if cntCol == 0 || cntRow == 0 {
		panic("cntCol либо cntRow равны 0")
	}
	buf.WriteByte('(')
	sum := cntCol * cntRow
	for i := 0; i < sum; i++ {
		if i != 0 {
			if cntRow > 1 && i%cntCol == 0 {
				buf.WriteString("),(")
			} else {
				buf.WriteByte(',')
			}
		}
		buf.WriteString("$" + strconv.Itoa(1+i))
	}
	buf.WriteByte(')')
	return buf.String()
}
