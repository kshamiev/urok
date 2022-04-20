SELECT
	t.table_name,
	d.description
FROM
	pg_catalog.pg_description AS d
INNER JOIN pg_catalog.pg_class AS cl ON
	d.objoid = cl.oid
INNER JOIN information_schema."tables" t ON
	t.table_name = cl.relname
INNER JOIN pg_namespace AS n ON
	n.oid = cl.relnamespace
	AND n.nspname = t.table_schema
ORDER BY
	t.table_name ASC
