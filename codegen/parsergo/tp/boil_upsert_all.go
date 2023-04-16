// nolint
package tp

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (o FileServiceSlice) UpsertAll(
	ctx context.Context,
	exec boil.ContextExecutor,
	updateOnConflict bool,
	conflictColumns []string,
	updateColumns,
	insertColumns boil.Columns,
) (int64, error) {
	var i int
	if o == nil {
		return 0, errors.New("срез пустой")
	}

	insert := insertColumns.UpdateColumnSet(
		fileServiceAllColumns,
		nil,
	)
	update := updateColumns.UpdateColumnSet(
		fileServiceAllColumns,
		fileServicePrimaryKeyColumns,
	)
	if len(conflictColumns) == 0 {
		conflictColumns = make([]string, len(fileServicePrimaryKeyColumns))
		copy(conflictColumns, fileServicePrimaryKeyColumns)
	}
	if updateOnConflict && len(update) == 0 {
		return 0, errors.New("поля для обновления не указаны")
	}

	qu := upsertAll("minio", updateOnConflict, len(o), insert, conflictColumns, update)

	values := make([]interface{}, 0, len(o)*len(insert))
	value := reflect.Indirect(reflect.ValueOf(o[0]))
	cnt := value.NumField()
	colMap := make(map[string]int, cnt)
	for i = 0; i < cnt; i++ {
		colMap[value.Type().Field(i).Tag.Get("boil")] = i
	}
	var ok bool
	for i = range insert {
		if _, ok = colMap[insert[i]]; !ok {
			panic("поля вставки не соответствуют структуре")
		}
	}
	for i = range o {
		value = reflect.Indirect(reflect.ValueOf(o[i]))
		for i = range insert {
			values = append(values, reflect.Indirect(value.Field(colMap[insert[i]])).Interface())
		}
	}
	res, err := exec.ExecContext(ctx, qu, values...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (o OrderSlice) UpsertAll(
	ctx context.Context,
	exec boil.ContextExecutor,
	updateOnConflict bool,
	conflictColumns []string,
	updateColumns,
	insertColumns boil.Columns,
) (int64, error) {
	var i int
	if o == nil {
		return 0, errors.New("срез пустой")
	}

	if !boil.TimestampsAreSkipped(ctx) {
		var i int
		currTime := time.Now().In(boil.GetLocation())
		for i = range o {
			o[i].UpdatedAt = currTime
		}
	}

	insert := insertColumns.UpdateColumnSet(
		orderAllColumns,
		nil,
	)
	update := updateColumns.UpdateColumnSet(
		orderAllColumns,
		orderPrimaryKeyColumns,
	)
	if len(conflictColumns) == 0 {
		conflictColumns = make([]string, len(orderPrimaryKeyColumns))
		copy(conflictColumns, orderPrimaryKeyColumns)
	}
	if updateOnConflict && len(update) == 0 {
		return 0, errors.New("поля для обновления не указаны")
	}

	qu := upsertAll("orders", updateOnConflict, len(o), insert, conflictColumns, update)

	values := make([]interface{}, 0, len(o)*len(insert))
	value := reflect.Indirect(reflect.ValueOf(o[0]))
	cnt := value.NumField()
	colMap := make(map[string]int, cnt)
	for i = 0; i < cnt; i++ {
		colMap[value.Type().Field(i).Tag.Get("boil")] = i
	}
	var ok bool
	for i = range insert {
		if _, ok = colMap[insert[i]]; !ok {
			panic("поля вставки не соответствуют структуре")
		}
	}
	for i = range o {
		value = reflect.Indirect(reflect.ValueOf(o[i]))
		for i = range insert {
			values = append(values, reflect.Indirect(value.Field(colMap[insert[i]])).Interface())
		}
	}
	res, err := exec.ExecContext(ctx, qu, values...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (o RoleSlice) UpsertAll(
	ctx context.Context,
	exec boil.ContextExecutor,
	updateOnConflict bool,
	conflictColumns []string,
	updateColumns,
	insertColumns boil.Columns,
) (int64, error) {
	var i int
	if o == nil {
		return 0, errors.New("срез пустой")
	}

	insert := insertColumns.UpdateColumnSet(
		roleAllColumns,
		nil,
	)
	update := updateColumns.UpdateColumnSet(
		roleAllColumns,
		rolePrimaryKeyColumns,
	)
	if len(conflictColumns) == 0 {
		conflictColumns = make([]string, len(rolePrimaryKeyColumns))
		copy(conflictColumns, rolePrimaryKeyColumns)
	}
	if updateOnConflict && len(update) == 0 {
		return 0, errors.New("поля для обновления не указаны")
	}

	qu := upsertAll("roles", updateOnConflict, len(o), insert, conflictColumns, update)

	values := make([]interface{}, 0, len(o)*len(insert))
	value := reflect.Indirect(reflect.ValueOf(o[0]))
	cnt := value.NumField()
	colMap := make(map[string]int, cnt)
	for i = 0; i < cnt; i++ {
		colMap[value.Type().Field(i).Tag.Get("boil")] = i
	}
	var ok bool
	for i = range insert {
		if _, ok = colMap[insert[i]]; !ok {
			panic("поля вставки не соответствуют структуре")
		}
	}
	for i = range o {
		value = reflect.Indirect(reflect.ValueOf(o[i]))
		for i = range insert {
			values = append(values, reflect.Indirect(value.Field(colMap[insert[i]])).Interface())
		}
	}
	res, err := exec.ExecContext(ctx, qu, values...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (o UserSlice) UpsertAll(
	ctx context.Context,
	exec boil.ContextExecutor,
	updateOnConflict bool,
	conflictColumns []string,
	updateColumns,
	insertColumns boil.Columns,
) (int64, error) {
	var i int
	if o == nil {
		return 0, errors.New("срез пустой")
	}

	if !boil.TimestampsAreSkipped(ctx) {
		var i int
		currTime := time.Now().In(boil.GetLocation())
		for i = range o {
			o[i].UpdatedAt = currTime
		}
	}

	insert := insertColumns.UpdateColumnSet(
		userAllColumns,
		nil,
	)
	update := updateColumns.UpdateColumnSet(
		userAllColumns,
		userPrimaryKeyColumns,
	)
	if len(conflictColumns) == 0 {
		conflictColumns = make([]string, len(userPrimaryKeyColumns))
		copy(conflictColumns, userPrimaryKeyColumns)
	}
	if updateOnConflict && len(update) == 0 {
		return 0, errors.New("поля для обновления не указаны")
	}

	qu := upsertAll("users", updateOnConflict, len(o), insert, conflictColumns, update)

	values := make([]interface{}, 0, len(o)*len(insert))
	value := reflect.Indirect(reflect.ValueOf(o[0]))
	cnt := value.NumField()
	colMap := make(map[string]int, cnt)
	for i = 0; i < cnt; i++ {
		colMap[value.Type().Field(i).Tag.Get("boil")] = i
	}
	var ok bool
	for i = range insert {
		if _, ok = colMap[insert[i]]; !ok {
			panic("поля вставки не соответствуют структуре")
		}
	}
	for i = range o {
		value = reflect.Indirect(reflect.ValueOf(o[i]))
		for i = range insert {
			values = append(values, reflect.Indirect(value.Field(colMap[insert[i]])).Interface())
		}
	}
	res, err := exec.ExecContext(ctx, qu, values...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
