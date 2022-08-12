INSERT INTO films
(title, description, category_id, release_year, price)
SELECT title, description, category_id, release_year, price
FROM films WHERE id IN (349, 696);

INSERT INTO films 
	(title, description, category_id, release_year, price, created_at, updated_at, deleted_at)
SELECT title, description, category_id, release_year, price, created_at, updated_at, deleted_at FROM films;

SELECT count(*) FROM films;

UPDATE films SET created_at = now() - INTERVAL '3 MONTH', updated_at = now() - INTERVAL '2 MONTH' WHERE id < 1001;
UPDATE films SET created_at = now() - INTERVAL '5 MONTH', updated_at = now() - INTERVAL '4 MONTH' WHERE id > 1000;

SELECT updated_at, to_timestamp(extract(epoch from updated_at::timestamptz)) FROM manticore_breaker;

UPDATE films SET created_at = , updated_at = WHERE id BETWEEN 1 AND 10;
SELECT count(*) FROM films f WHERE created_at BETWEEN '2019-05-10 12:49:34.779' AND '2021-01-10 12:49:34.779';  

SELECT * FROM films f WHERE id BETWEEN 9000000 AND 9000001;

SELECT * FROM films f WHERE id > 10000000;
