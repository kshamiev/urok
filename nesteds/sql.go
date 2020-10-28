package nesteds

const table = `trees`
const sel = `id, keyl, keyr, level, tree, parent_id, name`

// Gets
const (
	sqlLoadFromTree     = `SELECT ` + sel + ` FROM ` + table + ` WHERE tree = $1 ORDER BY id ASC LIMIT 1`
	sqlLoadFromID       = `SELECT ` + sel + ` FROM ` + table + ` WHERE id = $1`
	sqlLoadFromKeyl     = `SELECT ` + sel + ` FROM ` + table + ` WHERE keyl = $1`
	sqlLoadFromParentID = `SELECT ` + sel + ` FROM ` + table + ` WHERE parent_id = $1`
	// unlink
	sqlGetUnlink1 = `SELECT ` + sel + ` FROM ` + table + ` WHERE
	(( keyl < $1 AND keyr < $2 ) OR ( keyl > $3 AND keyr > $4 ) OR ( keyl > $5 AND keyr < $6 AND parent_id != $7 ))
	AND tree = $8
	ORDER BY name ASC`
	sqlGetUnlink2 = `SELECT ` + sel + ` FROM ` + table + ` WHERE level > 1 AND tree = $1 ORDER BY mame ASC`
	// Песочный путь
	sqlGetSandPath = `SELECT ` + sel + ` FROM ` + table + `	WHERE keyl <= $1 AND keyr >= $2 AND tree = $3 ORDER BY keyl ASC`
	// Выборка дочерней ветки
	sqlGetBranch          = `SELECT ` + sel + ` FROM ` + table + ` WHERE keyl > $1 AND keyr < $2 AND tree = $3 ORDER  BY keyl ASC`
	sqlGetBranchCountNode = `SELECT COUNT(*) as id FROM ` + table + ` WHERE keyl > $1 AND keyr < $2 AND tree = $3`
	sqlGetChildNode       = `SELECT ` + sel + ` FROM ` + table + ` WHERE parent_id = $1 AND tree = $2 ORDER  BY keyl ASC`
	sqlGetChildNodeCount  = `SELECT COUNT(*) as id FROM ` + table + ` WHERE parent_id = $1 AND tree = $2`
)

// Add
const (
	//  первый узел
	sqlAddFirst = `INSERT INTO ` + table + ` (keyl, keyr, level, tree, name) VALUES(1, 2, 1, $1, $2) RETURNING "id"`
	//  сдвиг правой и родительской стороны
	sqlShift = `UPDATE ` + table + `
	SET
		keyr = keyr + 2,
		keyl = case when keyl > $1 then keyl + 2 ELSE keyl END
	WHERE
		keyr >= $2
		AND tree = $3`
	//  добавление дочернего узла
	sqlAddChild = `INSERT INTO ` + table + ` (parent_id, keyl, keyr, level, tree, name) VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id"`
)

// Move & Copy
const (
	// unlink check
	sqlUnlinkCheck = `SELECT COUNT(id) id FROM ` + table + ` WHERE
	(( keyl < $1 AND keyr < $2 ) OR ( keyl > $3 AND keyr > $4 ) OR ( keyl > $5 AND keyr < $6 AND parent_id != $7 ))
	AND id = $8 AND tree = $9`
	//	псевдо удаление узла из дерева ( прячем в минус )
	sqlMinus = `UPDATE ` + table + ` SET keyl = keyl * -1, keyr = keyr * -1 WHERE keyl >= $1 AND Keyr <= $2 AND tree = $3`
	//	обновление дерева вставка
	sqlPaste = `UPDATE ` + table + `
	SET
		keyr = keyr + $1,
		keyl = case when keyl > $2 then keyl + $3 ELSE keyl END
	WHERE
		keyr >= $4
		AND tree = $5`
	// перемещение узла (вывод из тени, из минуса  в плюс)
	sqlPlus = `UPDATE ` + table + `
	SET
		keyl = keyl * -1 + $1,
		keyr = keyr * -1 + $2,
		level = level + $3
	WHERE
		keyl < 0 AND keyr < 0 AND tree = $4`
	sqlParentSetID   = `UPDATE ` + table + ` SET parent_id = $1 WHERE id = $2`
	sqlParentSetKeyl = `UPDATE ` + table + ` SET parent_id = $1 WHERE keyl = $2`

	sqlLoadCopy = `INSERT INTO ` + table + ` ("name", keyl, keyr, "level", tree, parent_id) 
	SELECT "name", keyl*-1, keyr*-1, "level", tree, parent_id 
	FROM ` + table + ` WHERE keyl >= $1 AND keyr <= $2 AND tree = $3 ORDER BY keyl ASC`
	sqlParentSet = `UPDATE ` + table + ` SET parent_id = $1 WHERE keyl > $2 AND keyr < $3 AND level = $4 AND tree = $5`
)

// Delete
const (
	// удаление узла
	sqlDelete = `DELETE FROM ` + table + ` WHERE keyl >= $1 AND keyr <= $2 AND tree = $3`
	// обновление дерева после удаления узла
	sqlCut = `UPDATE ` + table + `
	SET
		keyr = keyr - $1,
		keyl = case when keyl > $2 then keyl - $3 ELSE keyl END
	WHERE
		keyr > $4
		AND tree = $5`
)

// Check & Repair
const (
	sqlNodes   = `SELECT id FROM ` + table + ` WHERE parent_id = $1 ORDER BY keyl`
	sqlRepair1 = `UPDATE ` + table + ` SET keyl = $1, level = $2 WHERE id = $3`
	sqlRepair2 = `UPDATE ` + table + ` SET keyr = $1 WHERE id = $2`

	//	Левый ключ ВСЕГДА меньше правого
	sqlCheck1 = `SELECT id FROM ` + table + ` WHERE keyl >= keyr AND tree = $1`
	//	Наименьший левый ключ ВСЕГДА равен 1
	//	Наибольший правый ключ ВСЕГДА равен двойному числу узлов
	sqlCheck2 = `SELECT COUNT(id) id, min(keyl) keyl, max(keyr) keyr FROM ` + table + ` WHERE tree = $1`
	//	Разница между правым и левым ключом ВСЕГДА нечетное число
	sqlCheck3 = `SELECT id FROM ` + table + ` WHERE mod((keyr-keyl),2) = 0 AND tree = $1`
	//	Если уровень узла нечетное число то тогда левый ключ ВСЕГДА нечетное число, то же самое и для четных чисел
	sqlCheck4 = `SELECT id FROM ` + table + ` WHERE mod((keyl-level+2),2) = 1 AND tree = $1`
	//	Ключи ВСЕГДА уникальны, вне зависимости от того правый он или левый
	sqlCheck5 = `
	SELECT t1.id FROM ` + table + ` as t1, ` + table + ` as t2
	WHERE (t1.keyl = t2.keyr or t1.keyr = t2.keyl) AND t1.tree = $1 AND t2.tree = $1
	UNION
	SELECT t1.id FROM ` + table + ` as t1, ` + table + ` as t2 
	WHERE (t1.keyl = t2.keyl or t1.keyr = t2.keyr) AND t1.tree = $1 AND t2.tree = $1
	GROUP BY t1.id
	HAVING COUNT(t1.id)>1`
)
