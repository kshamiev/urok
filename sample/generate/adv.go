// extended generation functions matching proto type and real type
// функции реализующие обработку пользовательских типов свойств для определяемых рабочих типов
// описание в протофайлах, приведение значений в обе стороны, вызов вспомогательных функций для обработки
//
// type conversion support functions
// вспомогательные функции реализующие уникальную обработку свойств для определяемых рабочих типов
// nolint
package typs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v8"

	"gitlab.services.mts.ru/3click/million/typ"
)

// Срез типов для которых нужно реализовать работу по GRPC
var Config = []interface{}{
	&ReqSample{},
	&ResSample{},
}

// extended generation functions matching proto type and real type
// функции реализующие обработку пользовательских типов свойств для определяемых рабочих типов
// описание в протофайлах, приведение значений в обе стороны, вызов вспомогательных функций для обработки
var ConfigHandlerFunc = map[string]func(int, string, string, string) (string, string, string){
	"decimal.Decimal":   GenerateFieldDecimal,
	"[]typ.UUID":        GenerateFieldUUIDSlice,
	"types.StringArray": GenerateFieldStringArray,
	"typ.UUID":          GenerateFieldUUID,
	"time.Time":         GenerateFieldTime,
	"null.Time":         GenerateFieldNullTime,
	"null.String":       GenerateFieldNullString,
	"null.Int":          GenerateFieldNullInt,
	"null.Bytes":        GenerateFieldNullBytes,
	"null.JSON":         GenerateFieldNullJSON,
}

// GenerateFieldNullJSON конвертация - сопоставление туда и обратно
func GenerateFieldNullJSON(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: tip.%s.JSON,\n", pProto, pType)
	tplMFrom = fmt.Sprintf("\t\t%s: pbFromNullJSON(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullBytes конвертация - сопоставление туда и обратно
func GenerateFieldNullBytes(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tip.%s.Bytes,\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: pbFromNullBytes(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullString конвертация - сопоставление туда и обратно
func GenerateFieldNullString(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tip.%s.String,\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: pbFromNullString(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullInt конвертация - сопоставление туда и обратно
func GenerateFieldNullInt(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint64 " + pMessage + " = " + strconv.Itoa(i+1) + " [(grpc.gateway.protoc_gen_swagger.options.openapiv2_field) = {type: INTEGER}];\n"
	tplMTo = fmt.Sprintf("%s: int64(tip.%s.Int),\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: pbFromNullInt(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullTime конвертация - сопоставление туда и обратно
func GenerateFieldNullTime(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tgoogle.protobuf.Timestamp " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: pbToTime(tip.%s.Time),\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: pbFromNullTime(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldTime конвертация - сопоставление туда и обратно
func GenerateFieldTime(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tgoogle.protobuf.Timestamp " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: pbToTime(tip.%s),\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: pbFromTime(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUUID конвертация - сопоставление туда и обратно
func GenerateFieldUUID(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tip.%s.String(),\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: typ.UUIDMustParse(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldDecimal конвертация - сопоставление туда и обратно
func GenerateFieldDecimal(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tip.%s.String(),\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: pbFromDecimal(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUUIDSlice конвертация - сопоставление туда и обратно
func GenerateFieldUUIDSlice(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\trepeated string " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: pbToUUIDS(tip.%s),\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: pbFromUUIDS(pit.%s),\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldStringArray конвертация - сопоставление туда и обратно
func GenerateFieldStringArray(i int, pType, pMessage, pProto string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\trepeated string " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tip.%s,\n", pProto, pType)
	tplMFrom = fmt.Sprintf("%s: pit.%s,\n", pType, pProto)
	return tplP, tplMFrom, tplMTo
}

// type conversion support functions
// вспомогательные функции реализующие уникальную обработку свойств для определяемых рабочих типов

// pbFromNullBytes перевод из примитива grpc в рабочий тип
func pbFromNullBytes(b []byte) null.Bytes {
	return null.BytesFrom(b)
}

// pbFromNullJSON перевод из примитива grpc в рабочий тип
func pbFromNullJSON(b []byte) null.JSON {
	return null.JSONFrom(b)
}

// pbFromNullString перевод из примитива grpc в рабочий
func pbFromNullString(s string) null.String {
	if s == "" {
		return null.String{}
	}
	return null.StringFrom(s)
}

// pbFromNullInt перевод из примитива grpc в рабочий
func pbFromNullInt(i int64) null.Int {
	if i == 0 {
		return null.Int{}
	}
	return null.IntFrom(int(i))
}

// pbFromNullTime перевод из примитива grpc в рабочий тип
func pbFromNullTime(d *timestamp.Timestamp) null.Time {
	dp, err := ptypes.Timestamp(d)
	if err != nil {
		return null.Time{}
	}
	return null.TimeFrom(dp)
}

// pbToTime перевод в примитив grpc из рабочего типа
func pbToTime(d time.Time) *timestamp.Timestamp {
	dp, err := ptypes.TimestampProto(d)
	if err != nil {
		dp, _ = ptypes.TimestampProto(time.Time{})
	}
	return dp
}

// pbFromTime перевод из примитива grpc в рабочий тип
func pbFromTime(d *timestamp.Timestamp) time.Time {
	dp, err := ptypes.Timestamp(d)
	if err != nil {
		dp = time.Time{}
	}
	return dp
}

// pbFromDecimal перевод из примитива grpc в рабочий тип
func pbFromDecimal(v string) decimal.Decimal {
	d, _ := decimal.NewFromString(v)
	return d
}

func pbToUUIDS(list []typ.UUID) []string {
	uu := make([]string, len(list))
	for i := range list {
		uu[i] = list[i].String()
	}
	return uu
}

func pbFromUUIDS(list []string) []typ.UUID {
	uu := make([]typ.UUID, len(list))
	for i := range list {
		uu[i] = typ.UUIDMustParse(list[i])
	}
	return uu
}
