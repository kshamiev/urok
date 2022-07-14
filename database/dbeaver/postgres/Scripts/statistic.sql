-- типы индексов
SELECT amname FROM pg_am;
SELECT * FROM pg_am;

-- Планировщик запросов учитывает информацию о порядке следования данных в запрашиваемой таблице чтобы определить какой индекс использовать.
-- индексное сканирование или сканирование по битовой карте
select attname, correlation from pg_stats where tablename = 't';

-- показывает сколько страниц данных занимает индекс
select relpages from pg_class where relname='t_c_idx';
select relpages from pg_class where relname='t_c_idx1';
SELECT * FROM pg_stats AS tt WHERE tt.tablename = 't_lower_idx';
select * from pg_stats WHERE tablename = 't';

-- получение свойства индекса (тип индекса)
select a.amname, p.name, pg_indexam_has_property(a.oid,p.name)
from pg_am a,
unnest(array['can_order','can_unique','can_multi_col','can_exclude']) p(name)
where a.amname = 'btree' order by a.amname;

/*
can_order
Метод доступа позволяет указать порядок сортировки значений при создании индекса (в настоящее время применимо только для btree);
can_unique
Поддержка ограничения уникальности и первичного ключа (применимо только для btree);
can_multi_col
Индекс может быть построен по нескольким столбцам;
can_exclude
Поддержка ограничения исключения EXCLUDE
*/

-- получение свойств существующего прикладного индекса (какие действия можно с ним делать)
select p.name, pg_index_has_property('t_c_idx'::regclass,p.name)
from unnest(array['clusterable','index_scan','bitmap_scan','backward_scan']) p(name);

-- получение свойств существующего прикладного индекса (какие ресурсы использованы и сколько)
select * from bt_metap('users_pk');

select type, live_items, dead_items, avg_item_size, page_size, free_size
from bt_page_stats('users_pk',1);

/*
clusterable
Возможность переупорядочивания строк таблицы в соответствии с данным индексом (кластеризация одноименной командой CLUSTER);
index_scan
Поддержка индексного сканирования. Это свойство может показаться странным, однако не все индексы могут выдавать TID по одному — некоторые выдают все результаты сразу и поддерживают только сканирование битовой карты;
bitmap_scan
Поддержка сканирования битовой карты;
backward_scan
Выдача результата в порядке, обратном указанному при создании индекса.
*/

-- свойство индексированного столбца
select p.name, pg_index_column_has_property('t_c_idx'::regclass,1,p.name)
from unnest(array['asc','desc','nulls_first','nulls_last','orderable','distance_orderable','returnable','search_array','search_nulls']) p(name);

SELECT * FROM pg_catalog.pg_stat_all_indexes AS tt WHERE tt.relname = 't';
SELECT * FROM pg_catalog.pg_stat_all_tables AS dd WHERE dd.relname = 't';
SELECT * FROM pg_catalog.pg_stat_activity psa ;

-- получение не работающих индексов (битых, сломанных индексов)
select indexrelid::regclass index_name, indrelid::regclass table_name from pg_index where not indisvalid;

show config_file;
SELECT version();
SELECT datname FROM pg_database;

-- схема таблица row_count
select table_schema, 
       table_name, 
       (xpath('/row/cnt/text()', xml_count))[1]::text::int as row_count
from (
  select table_name, table_schema, 
         query_to_xml(format('select count(*) as cnt from %I.%I', table_schema, table_name), false, true, '') as xml_count
  from information_schema.tables
  where table_schema = 'public' --<< change here for the schema you want
) t
;

-- Текущие соединения и запросы к БД
select * from pg_stat_activity;
select count(datid) from pg_stat_activity;
show max_connections;

-- возможные типы данных 
SELECT t.*, d.description
FROM pg_catalog.pg_type t
LEFT OUTER JOIN pg_catalog.pg_class c ON c.oid=t.typrelid
LEFT OUTER JOIN pg_catalog.pg_description d ON t.oid=d.objoid
WHERE t.typname IS NOT NULL;

-- максимальная длина строки в таблице с данными
SELECT max(length(value_new)) FROM data_history dh;

SELECT 
  nspname AS schema_name,
  relname AS table_name,
  reltuples AS table_count_rows,
  pg_size_pretty(pg_relation_size(relname::text)) AS table_size
FROM pg_class C
LEFT JOIN pg_namespace N ON (N.oid = C.relnamespace)
WHERE 
  nspname NOT IN ('pg_catalog', 'information_schema') AND
  relkind='r' 
ORDER BY reltuples DESC;
;

select pg_size_pretty(pg_relation_size('dedicated_distance'));
select pg_size_pretty(pg_relation_size('tariff_price'));

SELECT nspname || '.' || relname AS "relation",
    pg_size_pretty(pg_relation_size(C.oid)) AS "size"
  FROM pg_class C
  LEFT JOIN pg_namespace N ON (N.oid = C.relnamespace)
  WHERE nspname NOT IN ('pg_catalog', 'information_schema')
  ORDER BY pg_relation_size(C.oid) DESC
 ;

SELECT nspname || '.' || relname AS "relation",
    pg_size_pretty(pg_total_relation_size(C.oid)) AS "total_size"
  FROM pg_class C
  LEFT JOIN pg_namespace N ON (N.oid = C.relnamespace)
  WHERE nspname NOT IN ('pg_catalog', 'information_schema')
    AND C.relkind <> 'i'
    AND nspname !~ '^pg_toast'
  ORDER BY pg_total_relation_size(C.oid) DESC
  LIMIT 20
 ;

 
