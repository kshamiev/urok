package stmc

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

// index
// m := objValue.MethodByName("GetIndexNane")
// if !m.IsValid() {
// return objValue, nil, "", errors.New("Invalid method GetIndexName: " + objValue.String())
// }
// indexMC = m.Call([]reflect.Value{})[0].String()
func getReflectObject(
	object interface{},
) (
	objValue reflect.Value, colsTrans map[string]string, err error,
) {
	// рефлексия объекта
	objValue = reflect.ValueOf(object)
	if objValue.Kind() != reflect.Ptr {
		return objValue, nil, errors.New("object is not ptr:" + objValue.String())
	}
	if objValue.IsNil() {
		return objValue, nil, errors.New("object is nil: " + objValue.String())
	}
	// сопоставление полей
	objValue = objValue.Elem()
	var fieldTag string
	colsTrans = map[string]string{}
	for i := 0; i < objValue.NumField(); i++ {
		fieldTag = objValue.Type().Field(i).Tag.Get(propertyTag)
		if fieldTag == `` || fieldTag == `-` || !objValue.Type().Field(i).IsExported() {
			continue
		}
		colsTrans[fieldTag] = objValue.Type().Field(i).Name
	}
	return objValue, colsTrans, nil
}

func insertObject(
	colsTrans map[string]string, objValue reflect.Value, dot decimal.Decimal,
) (
	properties []string, placeholders []string, args []interface{},
) {
	var (
		t         time.Time
		d         decimal.Decimal
		ok        bool
		tag, name string
		prop      reflect.Value
	)
	properties = make([]string, 0, len(colsTrans))
	placeholders = make([]string, 0, len(colsTrans))
	args = make([]interface{}, 0, len(colsTrans))
	for tag, name = range colsTrans {
		prop = objValue.FieldByName(name)
		switch prop.Type().String() {
		case propertyTypeString:
			properties = append(properties, tag)
			placeholders = append(placeholders, "?")
			args = append(args, prop.String())
		case propertyTypeInt64:
			properties = append(properties, tag)
			placeholders = append(placeholders, "?")
			args = append(args, prop.Int())
		case propertyTypeDecimal:
			d, ok = prop.Interface().(decimal.Decimal)
			if ok {
				properties = append(properties, tag)
				placeholders = append(placeholders, "?")
				args = append(args, d.Mul(dot).IntPart())
			}
		case propertyTypeTime:
			t, ok = prop.Interface().(time.Time)
			if ok {
				properties = append(properties, tag)
				placeholders = append(placeholders, "?")
				args = append(args, t.Unix())
			}
		case propertyTypeBool:
			properties = append(properties, tag)
			placeholders = append(placeholders, "?")
			args = append(args, prop.Bool())
		}
	}
	return properties, placeholders, args
}

func updateObject(
	colsTrans map[string]string, objValue reflect.Value, dot decimal.Decimal,
) (
	properties []string, args []interface{},
) {
	var (
		t         time.Time
		d         decimal.Decimal
		ok        bool
		tag, name string
		prop      reflect.Value
	)
	properties = make([]string, 0, len(colsTrans))
	args = make([]interface{}, 0, len(colsTrans))
	for tag, name = range colsTrans {
		prop = objValue.FieldByName(name)
		switch prop.Type().String() {
		case propertyTypeInt64:
			properties = append(properties, tag+"=?")
			args = append(args, prop.Int())
		case propertyTypeDecimal:
			d, ok = prop.Interface().(decimal.Decimal)
			if ok {
				properties = append(properties, tag+"=?")
				args = append(args, d.Mul(dot).IntPart())
			}
		case propertyTypeTime:
			t, ok = prop.Interface().(time.Time)
			if ok {
				properties = append(properties, tag+"=?")
				args = append(args, t.Unix())
			}
		case propertyTypeBool:
			properties = append(properties, tag+"=?")
			args = append(args, prop.Bool())
		}
	}
	return properties, args
}

func selectObject(colsValues map[string]string, objValue reflect.Value, dot decimal.Decimal) reflect.Value {
	var (
		name, value string
		prop        reflect.Value
	)
	for name, value = range colsValues {
		prop = objValue.FieldByName(name)
		switch prop.Type().String() {
		case propertyTypeString:
			prop.SetString(value)
		case propertyTypeInt64:
			n, _ := strconv.ParseInt(value, 10, 64)
			prop.SetInt(n)
		case propertyTypeDecimal:
			d, _ := decimal.NewFromString(value)
			d = d.Div(dot)
			prop.Set(reflect.ValueOf(d))
		case propertyTypeTime:
			n, _ := strconv.ParseInt(value, 10, 64)
			prop.Set(reflect.ValueOf(time.Unix(n, 0)))
		case propertyTypeBool:
			prop.SetBool(value == "1")
		}
	}
	return objValue
}
