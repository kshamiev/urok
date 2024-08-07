package stmc

import (
	"context"
	"fmt"
	"strings"
)

func (inst *Instance) Replace(ctx context.Context, object Index) error {
	objValue, colsTrans, err := getReflectObject(object)
	if err != nil {
		return err
	}

	// запрос
	delete(colsTrans, "highlight")
	properties, placeholders, args := insertObject(colsTrans, objValue, inst.Conf.DotDecimal)
	qu := `replace into ` + object.GetIndexName() + ` (` + strings.Join(properties, ", ")
	qu += `) values (` + strings.Join(placeholders, ", ") + `)`
	fmt.Println(qu)
	fmt.Println(args)
	_, err = inst.DB.ExecContext(ctx, qu, args...)
	if err != nil {
		return err
	}
	return nil
}
