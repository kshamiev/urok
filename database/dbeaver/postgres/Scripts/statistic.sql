show config_file;

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
select count(DISTINCT datid) from pg_stat_activity;
show max_connections;


SELECT t.oid,t.*,c.relkind,format_type(nullif(t.typbasetype, 0), t.typtypmod) as base_type_name, d.description
FROM pg_catalog.pg_type t
LEFT OUTER JOIN pg_catalog.pg_class c ON c.oid=t.typrelid
LEFT OUTER JOIN pg_catalog.pg_description d ON t.oid=d.objoid
WHERE t.typname IS NOT NULL;



SELECT max(length(value_new)) FROM data_history dh
;

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
ORDER BY reltuples DESC
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

 
