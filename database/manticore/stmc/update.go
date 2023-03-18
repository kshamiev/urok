package stmc

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

func (inst *Instance) Update(ctx context.Context, object Index) error {
	objValue, colsTrans, err := getReflectObject(object)
	if err != nil {
		return err
	}

	// запрос
	delete(colsTrans, "id")
	delete(colsTrans, "highlight")
	properties, args := updateObject(colsTrans, objValue, inst.Conf.DotDecimal)
	qu := `update ` + object.GetIndexName() + ` set ` + strings.Join(properties, ", ")
	qu += ` where id = ` + strconv.FormatInt(objValue.FieldByName("ID").Int(), 10)
	fmt.Println(qu)
	fmt.Println(args)
	_, err = inst.DB.ExecContext(ctx, qu, args...)
	if err != nil {
		return err
	}
	return nil
}
