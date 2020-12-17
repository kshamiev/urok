// Nested Set
// Реализация хранения и работы с деревьями в БД
// Пакет реализует работу с неограниченным количеством независимых деревьев в одной таблице
// Особенности реализации:
// Перемещать и копирвоать ветки и узлы в корень (верхний уровень) нельзя (в корне только корневые узлы отдельных деревьев)
// При создании корневого узла возвращается уже существующий, если он уже был ранее создан
//
// Пример описания Nested Set
// https://open2web.com.ua/blog/derevo-katalogov-nested-sets-vlozhennye-mnozhestva-i-upravlenie-im.html
/*
CREATE TABLE public.trees (
	id serial NOT NULL,
	parent_id int4 NOT NULL DEFAULT 0,
	keyl int4 NOT NULL DEFAULT 1,
	keyr int4 NOT NULL DEFAULT 2,
	"level" int4 NOT NULL DEFAULT 1,
	"tree" int4 NOT NULL DEFAULT 1,
	"name" text NOT NULL,
	CONSTRAINT tree_pk PRIMARY KEY (id)
);
*/
package nesteds

import "sync"

var mu sync.Mutex
