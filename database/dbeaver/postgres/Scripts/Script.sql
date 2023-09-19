-- Количество

SELECT
	count(*)
FROM
	goods g
WHERE 
	g."cost" < 100
;

SELECT
	count(*)
FROM
	goods g
WHERE 
	g.cost_null < 100
;

-- Среднее

SELECT
	avg(g."cost")
FROM
	goods g
WHERE
	1 = 1
	AND g.id < 100
;

SELECT
	avg(g.cost_null)
FROM
	goods g
WHERE
	1 = 1
	AND g.id < 100
;



