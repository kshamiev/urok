-- LEFT JOIN 
SELECT
	p.id,
	p.name,
	s.units 
FROM
	products p
LEFT JOIN sales s ON
	p.id = s.product_id
;

-- INNER JOIN
SELECT
	p.id,
	p.name,
	s.units 
FROM
	products p
INNER JOIN sales s ON
	p.id = s.product_id
;

-- CROSS JOIN
SELECT
	p.id,
	p.name,
	s.units 
FROM
	products p
CROSS JOIN sales s
;
SELECT
	p.id,
	p.name,
	s.units 
FROM
	products p, sales s
;
SELECT
	p.id,
	p.name,
	s.units 
FROM
	products p
INNER JOIN sales s ON true
;



