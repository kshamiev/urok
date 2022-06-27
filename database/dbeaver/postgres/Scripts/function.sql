--обращение к аргументам по номерам
CREATE OR REPLACE FUNCTION concat_lower_or_upper(a TEXT, b TEXT, uppercase boolean DEFAULT FALSE)
RETURNS TEXT AS
$$
SELECT
	CASE
		WHEN $3 THEN UPPER($1 || ' ' || $2)
		ELSE LOWER($1 || ' ' || $2)
	END;

$$
LANGUAGE SQL IMMUTABLE STRICT
;

SELECT concat_lower_or_upper('Hello', 'World', true);
SELECT concat_lower_or_upper('Hello', 'World');

--обращение к аргументам по именам
CREATE OR REPLACE FUNCTION concat_lower_or_upper(a TEXT, b TEXT, uppercase boolean DEFAULT FALSE)
RETURNS TEXT AS
$$
SELECT
	CASE
		WHEN uppercase THEN UPPER(a || ' ' || b)
		ELSE LOWER(a || ' ' || b)
	END;

$$
LANGUAGE SQL IMMUTABLE STRICT
;

-- Именная передача
SELECT concat_lower_or_upper(a => 'Hello', b => 'World', uppercase=> true);
SELECT concat_lower_or_upper(a => 'Hello', b => 'World');
SELECT concat_lower_or_upper(a := 'Hello', uppercase := true, b := 'World');

-- ПРИМЕР С СОСТАВНЫМ ТИПОМ И ВОЗВРАЩЕНИЕМ МНОЖЕСТВА (SETOF)

CREATE OR REPLACE FUNCTION get_product(ord orders)
RETURNS SETOF products AS $$
SELECT
	p.*
FROM
	products AS p
WHERE
	p.productcount = ord.productcount;

$$ LANGUAGE SQL;

SELECT get_product(o.*) FROM orders AS o WHERE o.id = 1;


-- ПРИМЕР С СОСТАВНЫМ ТИПОМ И ВОЗВРАЩЕНИЕМ МНОЖЕСТВА (SETOF) + ВАРИАНТ С ВЫХОДНЫМИ ПАРАМЕТРАМИ

-- TARGET SELECT, WHERE, HAVING
CREATE OR REPLACE FUNCTION get_products_out(ord orders, res OUT products)
RETURNS SETOF products -- (если возращается множество)
AS $$ 
SELECT
	p.*
FROM
	products AS p
WHERE
	p.productcount = ord.productcount;

$$ LANGUAGE SQL;

SELECT get_products_out(o.*) FROM orders AS o WHERE o.id = 1;

-- TARGET FROM
CREATE OR REPLACE FUNCTION get_products_out_from(productcount int, res OUT products)
RETURNS SETOF products -- (если возращается множество)
AS $$ 
SELECT
	p.*
FROM
	products AS p
WHERE
	p.productcount = get_products_out_from.productcount;

$$ LANGUAGE SQL;

SELECT * FROM get_products_out_from(2);

-- 
CREATE OR REPLACE FUNCTION sum_n_product (x int, y int, OUT sum int, OUT product int)
RETURNS SETOF record -- (если возращается множество)
AS $$
SELECT
	x + y, x * y
UNION
SELECT
	x * y, x + y
;

$$ LANGUAGE SQL;

SELECT * FROM sum_n_product(11,42);
SELECT (sum_n_product(11,42)).product;
SELECT product(sum_n_product(11,42));


-- 
CREATE OR REPLACE FUNCTION sum_n_product_t (x int, y int)
RETURNS TABLE(sum int, product int) -- (если возращается множество)
AS $$
SELECT
	x + y, x * y
UNION
SELECT
	x * y, x + y
;

$$ LANGUAGE SQL;

SELECT * FROM sum_n_product_t(11,42);
SELECT (sum_n_product_t(11,42)).product;
SELECT product(sum_n_product_t(11,42));

-- TODO
-- SELECT name, child FROM nodes, LATERAL listchildren(name) AS child;

