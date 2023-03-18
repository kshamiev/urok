package stmc

import (
	"context"
	"strings"

	"gitlab.tn.ru/golang/app/logger"
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
	logger.Get(ctx).Debug(qu)
	logger.Get(ctx).Debug(args)
	_, err = inst.DB.ExecContext(ctx, qu, args...)
	if err != nil {
		return err
	}
	return nil
}
