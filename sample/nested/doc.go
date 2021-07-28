// Nested Set
// Реализация хранения и работы с деревьями в БД
// Пакет реализует работу с одним деревом занимающим всю таблицу
//
// Пример описания Nested Set
// https://open2web.com.ua/blog/derevo-katalogov-nested-sets-vlozhennye-mnozhestva-i-upravlenie-im.html
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
