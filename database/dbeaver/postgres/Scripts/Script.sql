SELECT to_tsvector('english', 'a fat  cat sat on a mat - it ate age a fat fff fff fff fff fff fff fff fff fff fff fff fff fff fff fff rats');
SELECT to_tsquery('english', 'The & Fat & Rats');


SELECT get_current_ts_config ( ) ;


SELECT to_tsvector('english', u.login || u.description) FROM users u;
SELECT to_tsvector('russian', u.login || u.description) FROM users u;
SELECT to_tsvector('english', u.description) AS en, to_tsvector(u.description) AS ru FROM users u;


SELECT u.login || u.description AS doument FROM users u
;

SELECT 'a fat cat sat on a mat and ate a fat rat'::tsvector @@ 'cat & rat'::tsquery;
SELECT 'fat & cow'::tsquery @@ 'a fat cat sat on a mat and ate a fat rat'::tsvector;

SELECT to_tsvector('a fat cat sat on a mat and ate a fat rat') @@ to_tsquery('cats & rats');
SELECT to_tsquery('cats & rats');

SELECT u.id, u.description FROM users u WHERE to_tsvector(u.description) @@ to_tsquery('корабль');
SELECT u.id, u.description FROM users u WHERE u.description ILIKE '%корабль%';
SELECT u.id, u.description FROM users u WHERE u.description ILIKE '%леонид%';





