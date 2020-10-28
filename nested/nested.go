package nested

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Set struct {
	ID       int    `boil:"id" json:"id"`
	ParentID int    `boil:"parent_id" json:"parent_id"`
	Keyl     int    `boil:"keyl" json:"keyl"`
	Keyr     int    `boil:"keyr" json:"keyr"`
	Level    int    `boil:"level" json:"level"`
	Name     string `boil:"name" json:"name"`
	db       *sql.DB
}

func NewSet(db *sql.DB, id int) *Set {
	return &Set{
		ID: id,
		db: db,
	}
}

// Загрузка узла
func (set *Set) Load(ctx context.Context) error {
	if set.ID == 0 {
		return nil
	}
	return queries.Raw(sqlLoadFromID, set.ID).Bind(ctx, set.db, set)
}

// Создание узла
func (set *Set) Create(ctx context.Context, name string) (*Set, error) {
	if err := set.Load(ctx); err != nil {
		return nil, err
	}

	tx, err := set.db.Begin()
	if err != nil {
		return nil, err
	}
	mu.Lock()
	defer func() {
		_ = tx.Rollback()
		mu.Unlock()
	}()

	req := &Set{}
	if set.ID == 0 {
		// node up level
		if errN := queries.Raw(sqlMaxKeyr).Bind(ctx, tx, req); errN != nil {
			return nil, errN
		}
		if req.Keyr > 0 {
			err = queries.Raw(sqlAddParent, req.Keyr+1, req.Keyr+2, name).Bind(ctx, tx, req)
			if err != nil {
				return nil, err
			}
		} else {
			err = queries.Raw(sqlAddFirst, name).Bind(ctx, tx, req)
			if err != nil {
				return nil, err
			}
		}
	} else {
		// child node
		if _, err = queries.Raw(sqlShift, set.Keyr, set.Keyr).ExecContext(ctx, tx); err != nil {
			return nil, err
		}
		err = queries.Raw(
			sqlAddChild, set.ID, set.Keyr, set.Keyr+1, set.Level+1, name,
		).Bind(ctx, tx, req)
		if err != nil {
			return nil, err
		}
	}

	_ = tx.Commit()

	req = NewSet(set.db, req.ID)
	err = req.Load(ctx)
	return req, err
}

// Перемещение узла
func (set *Set) Move(ctx context.Context, child *Set) error {
	if err := set.checkMove(ctx, child); err != nil {
		return err
	}

	tx, err := set.db.Begin()
	if err != nil {
		return err
	}
	mu.Lock()
	defer func() {
		_ = tx.Rollback()
		mu.Unlock()
	}()

	// прячем в минус
	if _, errr := queries.Raw(sqlMinus, child.Keyl, child.Keyr).ExecContext(ctx, tx); errr != nil {
		return errr
	}

	stepDel := child.Keyr - child.Keyl + 1

	// вырезаем - сдвигаем дерево
	if _, errr := queries.Raw(
		sqlCut, stepDel, child.Keyl, stepDel, child.Keyr,
	).ExecContext(ctx, tx); errr != nil {
		return errr
	}

	if set.ID > 0 {
		// инициализация родителя
		if child.Keyr < set.Keyr {
			set.Keyr -= stepDel
			if child.Keyl < set.Keyl {
				set.Keyl -= stepDel
			}
		}
		// вставка - раздвигаем дерево
		if _, errr := queries.Raw(
			sqlPaste, stepDel, set.Keyr, stepDel, set.Keyr,
		).ExecContext(ctx, tx); errr != nil {
			return errr
		}

		//	вычисление смещения ключей для перемещаемой ветки
		stepIns := set.Keyr - child.Keyl
		stepLevel := set.Level - child.Level + 1
		set.Keyr += stepDel

		// выводим в плюс спрятанный узел
		if _, errr := queries.Raw(
			sqlPlus, stepIns, stepIns, stepLevel,
		).ExecContext(ctx, tx); errr != nil {
			return errr
		}
	} else {
		req := &Set{}
		if errr := queries.Raw(sqlMaxKeyr).Bind(ctx, tx, req); errr != nil {
			return errr
		}

		//	вычисление смещения ключей для перемещаемой ветки
		stepIns := req.Keyr - child.Keyl + 1
		stepLevel := child.Level*-1 + 1

		// выводим в плюс спрятанный узел
		if _, errr := queries.Raw(
			sqlPlus, stepIns, stepIns, stepLevel,
		).ExecContext(ctx, tx); errr != nil {
			return errr
		}
	}

	if _, errr := queries.Raw(
		sqlParentSetID, set.ID, child.ID,
	).ExecContext(ctx, tx); errr != nil {
		return errr
	}

	_ = tx.Commit()
	err = child.Load(ctx)

	return err
}

func (set *Set) checkMove(ctx context.Context, child *Set) error {
	if err := set.Load(ctx); err != nil {
		return err
	}
	if err := child.Load(ctx); err != nil {
		return err
	}

	req := &Set{}
	if err := queries.Raw(
		sqlUnlinkCheck, set.Keyl, set.Keyr, set.Keyl, set.Keyr, set.Keyl, set.Keyr, set.ID, child.ID,
	).Bind(ctx, set.db, req); err != nil {
		return err
	}
	if req.ID != 1 {
		return errors.New("impossible linked, these objects are already linked")
	}

	return nil
}

// Копирование узла
func (set *Set) Copy(ctx context.Context, child *Set) error {
	if err := set.Load(ctx); err != nil {
		return err
	}
	if err := child.Load(ctx); err != nil {
		return err
	}

	tx, err := set.db.Begin()
	if err != nil {
		return err
	}
	mu.Lock()
	defer func() {
		_ = tx.Rollback()
		mu.Unlock()
	}()

	if _, errr := queries.Raw(sqlLoadCopy, child.Keyl, child.Keyr).ExecContext(ctx, tx); errr != nil {
		return errr
	}

	var stepIns int
	if set.ID > 0 {
		stepCopy := child.Keyr - child.Keyl + 1
		// вставка - раздвигаем дерево
		if _, errr := queries.Raw(
			sqlPaste, stepCopy, set.Keyr, stepCopy, set.Keyr,
		).ExecContext(ctx, tx); errr != nil {
			return errr
		}

		//	вычисление смещения ключей для перемещаемой ветки
		stepIns = set.Keyr - child.Keyl
		stepLevel := set.Level - child.Level + 1
		set.Keyr += stepCopy

		// выводим в плюс спрятанный узел
		if _, errr := queries.Raw(
			sqlPlus, stepIns, stepIns, stepLevel,
		).ExecContext(ctx, tx); errr != nil {
			return errr
		}
	} else {
		req := &Set{}
		if errr := queries.Raw(sqlMaxKeyr).Bind(ctx, tx, req); errr != nil {
			return errr
		}

		//	вычисление смещения ключей для перемещаемой ветки
		stepIns = req.Keyr - child.Keyl + 1
		stepLevel := child.Level*-1 + 1

		// выводим в плюс спрятанный узел
		if _, errr := queries.Raw(
			sqlPlus, stepIns, stepIns, stepLevel,
		).ExecContext(ctx, tx); errr != nil {
			return errr
		}
	}

	if _, errr := queries.Raw(
		sqlParentSetKeyl, set.ID, stepIns+child.Keyl,
	).ExecContext(ctx, tx); errr != nil {
		return errr
	}

	obj := &Set{}
	if errr := queries.Raw(sqlLoadFromKeyl, stepIns+child.Keyl).Bind(ctx, tx, obj); errr != nil {
		return errr
	}

	if errr := set.parentUpdate(ctx, tx, obj); errr != nil {
		return errr
	}

	_ = tx.Commit()
	err = child.Load(ctx)

	return err
}

func (set *Set) parentUpdate(ctx context.Context, tx *sql.Tx, obj *Set) error {
	if _, err := queries.Raw(
		sqlParentSet, obj.ID, obj.Keyl, obj.Keyr, obj.Level+1,
	).ExecContext(ctx, tx); err != nil {
		return err
	}
	var data []*Set
	if err := queries.Raw(sqlLoadFromParentID, obj.ID).Bind(ctx, tx, &data); err != nil {
		return err
	}
	for i := range data {
		if err := set.parentUpdate(ctx, tx, data[i]); err != nil {
			return err
		}
	}
	return nil
}

// Удаление узла
func (set *Set) Delete(ctx context.Context) error {
	if set.ID == 0 {
		return nil
	}
	if err := set.Load(ctx); err != nil {
		return err
	}

	tx, err := set.db.Begin()
	if err != nil {
		return err
	}
	mu.Lock()
	defer func() {
		_ = tx.Rollback()
		mu.Unlock()
	}()

	if _, err := queries.Raw(sqlDelete, set.Keyl, set.Keyr).ExecContext(ctx, tx); err != nil {
		return err
	}

	step := set.Keyr - set.Keyl + 1

	if _, err := queries.Raw(
		sqlCut, step, set.Keyl, step, set.Keyr,
	).ExecContext(ctx, tx); err != nil {
		return err
	}

	_ = tx.Commit()

	return nil
}

// Получение узлов которые можно привязать к текущему
func (set *Set) GetUnlink(ctx context.Context) (data []*Set, err error) {
	if set.ID > 0 {
		if err := set.Load(ctx); err != nil {
			return nil, err
		}
		if err := queries.Raw(
			sqlGetUnlink1, set.Keyl, set.Keyr, set.Keyl, set.Keyr, set.Keyl, set.Keyr, set.ParentID,
		).Bind(ctx, set.db, &data); err != nil {
			return nil, err
		}
	} else if err := queries.Raw(sqlGetUnlink2).Bind(ctx, set.db, &data); err != nil {
		return nil, err
	}

	return
}

// Песочный путь
func (set *Set) GetSandPath(ctx context.Context) (data []*Set, err error) {
	if err := set.Load(ctx); err != nil {
		return nil, err
	}
	if err := queries.Raw(sqlGetSandPath, set.Keyl, set.Keyr).Bind(ctx, set.db, &data); err != nil {
		return nil, err
	}
	return
}

// Выборка дочерней ветки
func (set *Set) GetBranch(ctx context.Context) (data []*Set, err error) {
	if err := set.Load(ctx); err != nil {
		return nil, err
	}
	if err := queries.Raw(sqlGetBranch, set.Keyl, set.Keyr).Bind(ctx, set.db, &data); err != nil {
		return nil, err
	}
	return
}

// Выборка количества узлов ветки
func (set *Set) GetBranchCountNode(ctx context.Context) (int, error) {
	if err := set.Load(ctx); err != nil {
		return 0, err
	}
	obj := &Set{}
	if err := queries.Raw(sqlGetBranchCountNode, set.Keyl, set.Keyr).Bind(ctx, set.db, obj); err != nil {
		return 0, err
	}
	return obj.ID, nil
}

// Выборка дочерних узлов
func (set *Set) GetChildNode(ctx context.Context) (data []*Set, err error) {
	if err := set.Load(ctx); err != nil {
		return nil, err
	}
	if err := queries.Raw(sqlGetChildNode, set.ID).Bind(ctx, set.db, &data); err != nil {
		return nil, err
	}
	return
}

// Выборка количества дочерних узлов
func (set *Set) GetChildNodeCount(ctx context.Context) (int, error) {
	if err := set.Load(ctx); err != nil {
		return 0, err
	}
	obj := &Set{}
	if err := queries.Raw(sqlGetChildNodeCount, set.ID).Bind(ctx, set.db, obj); err != nil {
		return 0, err
	}
	return obj.ID, nil
}
