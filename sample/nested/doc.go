// Nested Set
// Реализация хранения и работы с деревьями в БД
// Пакет реализует работу с одним деревом занимающим всю таблицу
/*
CREATE TABLE public.tree (
	id serial NOT NULL,
	parent_id int4 NOT NULL DEFAULT 0,
	keyl int4 NOT NULL DEFAULT 1,
	keyr int4 NOT NULL DEFAULT 2,
	"level" int4 NOT NULL DEFAULT 1,
	"name" text NOT NULL,
	CONSTRAINT tree_pk PRIMARY KEY (id)
);
*/
package nested

import "sync"

var mu sync.Mutex
