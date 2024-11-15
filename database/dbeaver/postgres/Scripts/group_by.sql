SELECT CURRENT_DATE - INTERVAL '40 year';

-- Пример группировки
-- Статистика продаж по товарам за прошедший месяц
SELECT
	p.id,
	p.name,
	(sum(s.units) * (p.price - p.cost)) AS profit
FROM
	goods p
LEFT JOIN sales s ON
	p.id = s.good_id
WHERE
	1=1
	AND s.created_at > CURRENT_DATE - INTERVAL '4 weeks'
GROUP BY
	p.id,
	p.name
HAVING
	sum(p.price * s.units) > 5000;
	
-- Находим товар который не был ни разу продан за последние 40 лет
SELECT
	p.id,
	p.name
FROM
	goods p
LEFT JOIN sales s ON 
	p.id = s.good_id AND s.created_at > CURRENT_DATE - INTERVAL '30 year'   
WHERE
   s.good_id IS NULL
;


