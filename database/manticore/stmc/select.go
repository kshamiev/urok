package stmc

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"strconv"

	"gitlab.tn.ru/golang/app/logger"
)

func (inst *Instance) Select(ctx context.Context, object Index) error {
	objValue, colsTrans, err := getReflectObject(object)
	if err != nil {
		return err
	}

	// запрос
	qu := "select * from " + object.GetIndexName() + " where id = "
	qu += strconv.FormatInt(objValue.FieldByName("ID").Int(), 10)
	logger.Get(ctx).Debug(qu)
	rows, err := inst.DB.QueryContext(ctx, qu)
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

	var ok bool
	colsValues := map[string]string{}
	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			return err
		}
		for i, column := range cols {
			if _, ok = colsTrans[column]; !ok {
				continue
			}
			colsValues[colsTrans[column]] = string((*(values[i].(*interface{}))).([]byte))
		}
		selectObject(colsValues, objValue, inst.Conf.DotDecimal)
		return nil
	}
	return sql.ErrNoRows
}

func (inst *Instance) SelectCount(ctx context.Context, qu string, args ...interface{}) (int, error) {
	logger.Get(ctx).Debug(qu)
	logger.Get(ctx).Debug(args)
	var cnt int
	row := inst.DB.QueryRowContext(ctx, qu, args...)
	if row.Err() != nil {
		return cnt, row.Err()
	}
	if err := row.Scan(&cnt); err != nil {
		return cnt, err
	}
	return cnt, nil
}

func (inst *Instance) SelectSlice(ctx context.Context, result interface{}, qu string, args ...interface{}) error {
	// рефлексия объекта - сопоставление полей
	objType := reflect.TypeOf(result)
	if objType.Kind() != reflect.Ptr {
		return errors.New("result is not ptr: " + objType.String())
	}
	if objType.Elem().Kind() != reflect.Slice {
		return errors.New("result is not Slice: " + objType.String())
	}
	var isPtr bool
	obj := objType.Elem().Elem()
	if obj.Kind() == reflect.Ptr {
		isPtr = true
		obj = obj.Elem()
	}
	var fieldTag string
	colsTrans := map[string]string{}
	for i := 0; i < obj.NumField(); i++ {
		fieldTag = obj.Field(i).Tag.Get(propertyTag)
		if fieldTag == `` || fieldTag == `-` || !obj.Field(i).IsExported() {
			continue
		}
		colsTrans[fieldTag] = obj.Field(i).Name
	}

	// запрос
	logger.Get(ctx).Debug(qu)
	logger.Get(ctx).Debug(args)
	rows, err := inst.DB.QueryContext(ctx, qu, args...)
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

	var objValueSlice = reflect.MakeSlice(objType.Elem(), 0, 0)
	var objValue reflect.Value
	var ok bool
	colsValues := map[string]string{}
	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			return err
		}
		for i, column := range cols {
			if _, ok = colsTrans[column]; !ok {
				continue
			}
			colsValues[colsTrans[column]] = string((*(values[i].(*interface{}))).([]byte))
		}
		objValue = selectObject(colsValues, reflect.New(obj).Elem(), inst.Conf.DotDecimal)
		if isPtr {
			objValueSlice = reflect.Append(objValueSlice, objValue.Addr())
		} else {
			objValueSlice = reflect.Append(objValueSlice, objValue)
		}
	}
	reflect.ValueOf(result).Elem().Set(objValueSlice)
	return nil
}

func (inst *Instance) SelectParser(ctx context.Context, result Parser, qu string, args ...interface{}) error {
	// запрос
	logger.Get(ctx).Debug(qu)
	logger.Get(ctx).Debug(args)
	rows, err := inst.DB.QueryContext(ctx, qu, args...)
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
		result.Parse(dest)
	}
	return nil
}
