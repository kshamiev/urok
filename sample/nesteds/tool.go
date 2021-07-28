package nesteds

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Tool struct {
	Tree int `boil:"tree" json:"tree"`
	db   *sql.DB
}

func NewTool(db *sql.DB, tree int) *Tool {
	return &Tool{
		Tree: tree,
		db:   db,
	}
}

// Проверка целостности дерева
func (t *Tool) Check(ctx context.Context) error {
	//	Левый ключ ВСЕГДА меньше правого
	reqList := []Set{}
	if err := queries.Raw(sqlCheck1, t.Tree).Bind(ctx, t.db, &reqList); err != nil {
		return err
	}
	if len(reqList) > 0 {
		return errors.New("левый ключ не меньше правого")
	}
	//	Наименьший левый ключ ВСЕГДА равен 1
	//	Наибольший правый ключ ВСЕГДА равен двойному числу узлов
	req := &Set{}
	if err := queries.Raw(sqlCheck2, t.Tree).Bind(ctx, t.db, req); err != nil {
		return err
	}
	if req.Keyl != 1 {
		return errors.New("наименьший левый ключ не равен 1")
	}
	if req.Keyr != req.ID*2 {
		return errors.New("наибольший правый ключ не равен двойному числу узлов")
	}
	//	Разница между правым и левым ключом ВСЕГДА нечетное число
	reqList = []Set{}
	if err := queries.Raw(sqlCheck3, t.Tree).Bind(ctx, t.db, &reqList); err != nil {
		return err
	}
	if len(reqList) > 0 {
		return errors.New("разница между правым и левым ключом четное число")
	}
	//	Если уровень узла нечетное число то тогда левый ключ ВСЕГДА нечетное число, то же самое и для четных чисел
	reqList = []Set{}
	if err := queries.Raw(sqlCheck4, t.Tree).Bind(ctx, t.db, &reqList); err != nil {
		return err
	}
	if len(reqList) > 0 {
		return errors.New("четность левого ключа и его уровня не совпадают")
	}
	//	Ключи ВСЕГДА уникальны, вне зависимости от того правый он или левый
	reqList = []Set{}
	if err := queries.Raw(sqlCheck5, t.Tree).Bind(ctx, t.db, &reqList); err != nil {
		return err
	}
	if len(reqList) > 0 {
		return errors.New("ключи не уникальны")
	}
	return nil
}

// Восстановление дерева по рефлексивным связям
func (t *Tool) Repair(ctx context.Context) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	mu.Lock()
	defer func() {
		_ = tx.Rollback()
		mu.Unlock()
	}()

	req := &Set{}
	err = queries.Raw(sqlLoadFromTree, t.Tree).Bind(ctx, tx, req)
	if err != nil {
		return err
	}

	if _, err = queries.Raw(sqlRepair1, 1, 1, req.ID).ExecContext(ctx, tx); err != nil {
		return err
	}

	var keyl int
	if keyl, err = t.repair(ctx, tx, req.ID, 2, 2); err != nil {
		return err
	}

	if _, err = queries.Raw(sqlRepair2, keyl, req.ID).ExecContext(ctx, tx); err != nil {
		return err
	}

	_ = tx.Commit()

	return nil
}

func (t *Tool) repair(ctx context.Context, tx *sql.Tx, parentID, keyl, level int) (int, error) {
	req := []Set{}
	if err := queries.Raw(sqlNodes, parentID).Bind(ctx, tx, &req); err != nil {
		return 0, err
	}
	for i := range req {
		if _, err := queries.Raw(sqlRepair1, keyl, level, req[i].ID).ExecContext(ctx, tx); err != nil {
			return 0, err
		}
		var err error
		keyl, err = t.repair(ctx, tx, req[i].ID, keyl+1, level+1)
		if err != nil {
			return 0, err
		}
		if _, err := queries.Raw(sqlRepair2, keyl, req[i].ID).ExecContext(ctx, tx); err != nil {
			return 0, err
		}
		keyl++
	}
	return keyl, nil
}
