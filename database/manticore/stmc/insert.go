package stmc

import (
	"context"
	"fmt"
	"strings"
)

func (inst *Instance) Insert(ctx context.Context, object Index) error {
	objValue, colsTrans, err := getReflectObject(object)
	if err != nil {
		return err
	}

	// запрос
	delete(colsTrans, "highlight")
	properties, placeholders, args := insertObject(colsTrans, objValue, inst.Conf.DotDecimal)
	qu := `insert into ` + object.GetIndexName() + ` (` + strings.Join(properties, ", ")
	qu += `) values (` + strings.Join(placeholders, ", ") + `)`
	fmt.Println(qu)
	fmt.Println(args)
	res, err := inst.DB.ExecContext(ctx, qu, args...)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	objValue.FieldByName("ID").SetInt(id)
	return nil
}
