INSERT INTO films
(title, description, category_id, release_year, price)
SELECT title, description, category_id, release_year, price
FROM films WHERE id IN (349, 696);

INSERT INTO documents 
	(title, description, category_id, release_year, price, created_at, updated_at, deleted_at)
SELECT title, description, category_id, release_year, price, created_at, updated_at, deleted_at FROM films;

INSERT INTO documents  
	(title, description, category_id, release_year, price, created_at, updated_at, deleted_at)
SELECT title, description, category_id, release_year, price, created_at, updated_at, deleted_at FROM documents1 ORDER BY updated_at DESC LIMIT 10000000;

SELECT count(*) FROM documents;
SELECT count(*) FROM documents1;

UPDATE documents SET created_at = now() - INTERVAL '3 MONTH', updated_at = now() - INTERVAL '2 MONTH' WHERE id < 1001;
UPDATE documents SET created_at = now() - INTERVAL '5 MONTH', updated_at = now() - INTERVAL '4 MONTH' WHERE id > 1000;

SELECT start_at, to_timestamp(extract(epoch from start_at::timestamptz)) FROM manticore_breaker;

DELETE FROM documents WHERE id > 4000000;

UPDATE documents SET id = id - 1;
UPDATE documents SET updated_at = now(), deleted_at = NULL WHERE id > 4;
UPDATE documents SET DATA = '{"id":12, "name":"funtik", "price":34.67}';
UPDATE documents SET is_flag = true;


SELECT * FROM documents1 ORDER BY updated_at DESC LIMIT 10;
