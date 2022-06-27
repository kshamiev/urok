-- LEFT JOIN 
SELECT
	p.id,
	p.name,
	s.units 
FROM
	goods p
LEFT JOIN sales s ON
	p.id = s.good_id
;

-- INNER JOIN
SELECT
	p.id,
	p.name,
	s.units 
FROM
	goods p
INNER JOIN sales s ON
	p.id = s.good_id
;

-- CROSS JOIN
SELECT
	p.id,
	p.name,
	s.units 
FROM
	goods p
CROSS JOIN sales s
;
SELECT
	p.id,
	p.name,
	s.units 
FROM
	goods p, sales s
;
SELECT
	p.id,
	p.name,
	s.units 
FROM
	goods p
INNER JOIN sales s ON true
;



