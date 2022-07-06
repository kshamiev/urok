--обращение к аргументам по номерам
CREATE OR REPLACE FUNCTION concat_lower_or_upper(a TEXT, b TEXT, uppercase boolean DEFAULT FALSE)
RETURNS TEXT AS
$$
DECLARE
	res TEXT; 
BEGIN
SELECT
	CASE
		WHEN uppercase THEN UPPER($1 || ' ' || $2)
		ELSE LOWER($1 || ' ' || $2)
	END INTO res;
RETURN res;
END;
$$ LANGUAGE plpgsql IMMUTABLE STRICT;

SELECT concat_lower_or_upper('Hello', 'World', true);
SELECT concat_lower_or_upper('Hello', 'World');

--обращение к аргументам по именам
CREATE OR REPLACE FUNCTION concat_lower_or_upper(a TEXT, b TEXT, uppercase boolean DEFAULT FALSE)
RETURNS TEXT AS
$$
DECLARE
	res TEXT; 
BEGIN
SELECT
	CASE
		WHEN uppercase THEN UPPER(a || ' ' || b)
		ELSE LOWER(a || ' ' || b)
	END INTO res;
RETURN res;
END;
$$ LANGUAGE plpgsql IMMUTABLE STRICT;

-- Именная передача
SELECT concat_lower_or_upper(a => 'Hello', b => 'World', uppercase=> true);
SELECT concat_lower_or_upper(a => 'Hello', b => 'World');
SELECT concat_lower_or_upper(a := 'Hello', uppercase := true, b := 'World');

-- ПРИМЕР С СОСТАВНЫМ ТИПОМ И ВОЗВРАЩЕНИЕМ МНОЖЕСТВА (SETOF)
-- если есть таблица подходящая под возвращаемые данные

CREATE OR REPLACE FUNCTION get_product_SETOF(ord orders)
RETURNS SETOF products AS $$ 
BEGIN
RETURN QUERY SELECT p.id, p.productname, p.company, p.productcount, p.price
FROM
	products AS p
WHERE
	p.productcount = ord.productcount;
END;
$$ LANGUAGE plpgsql VOLATILE;

SELECT get_product_SETOF(o.*) FROM orders AS o WHERE o.id = 1;

CREATE OR REPLACE FUNCTION get_products_SETOF_from(productcount int)
RETURNS SETOF products
AS $$ 
BEGIN
RETURN QUERY SELECT
	p.*
FROM
	products AS p
WHERE
	p.productcount = get_products_out_from.productcount;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM get_products_SETOF_from(2);

-- ПРИМЕР С СОСТАВНЫМ ТИПОМ И ВОЗВРАЩЕНИЕМ МНОЖЕСТВА (TABLE)
-- анонимный составной тип если нет подходящей таблицы

CREATE OR REPLACE FUNCTION get_product_TABLE(ord orders)
RETURNS TABLE (id int, productname TEXT, company TEXT, productcount int, price numeric) AS $$
BEGIN
RETURN QUERY SELECT p.id, p.productname, p.company, p.productcount, p.price
FROM
	products AS p
WHERE
	p.productcount = ord.productcount;
END;
$$ LANGUAGE plpgsql VOLATILE;

SELECT get_product2(o.*) FROM orders AS o WHERE o.id = 1;

-- ПРИМЕР С СОСТАВНЫМ ТИПОМ И ВОЗВРАЩЕНИЕМ МНОЖЕСТВА (RECORD) + ВАРИАНТ С ВЫХОДНЫМИ ПАРАМЕТРАМИ

CREATE OR REPLACE FUNCTION sum_n_product (x int, y int, OUT sum int, OUT product int)
RETURNS SETOF record
AS $$
BEGIN
RETURN QUERY SELECT
	x + y, x * y
UNION
SELECT
	x * y, x + y
;
END;
$$ LANGUAGE plpgsql;

SELECT * FROM sum_n_product(11,42);
SELECT (sum_n_product(11,42)).product;
SELECT product(sum_n_product(11,42));

-- TODO
-- SELECT name, child FROM nodes, LATERAL listchildren(name) AS child;

