-- LATERAL

SELECT
	*
FROM
	Products
WHERE
	Price = (SELECT MIN(Price) FROM Products)
;
		
SELECT
	*
FROM
	Products
WHERE
	Price > (SELECT AVG(Price) FROM Products)
;
	
SELECT
	CreatedAt,
	Price,
	(SELECT ProductName FROM Products WHERE Products.Id = Orders.ProductId) AS Product
FROM
	Orders
;

SELECT
	ProductName,
	Company,
	Price,
	(SELECT	AVG(Price) FROM Products AS SubProds WHERE SubProds.Company = Prods.Company) AS AvgPrice
FROM
	Products AS Prods
WHERE
	Price > (SELECT AVG(Price) FROM Products AS SubProds WHERE SubProds.Company = Prods.Company)
;

SELECT 
*
FROM products p 
INNER JOIN LATERAL (SELECT * FROM orders WHERE p.id = orders.productid ) ord ON true 
;

SELECT 
*
FROM products p 
CROSS JOIN LATERAL (SELECT * FROM orders WHERE p.id = orders.productid ) ord 
;


SELECT 
*
FROM products p 
LEFT JOIN LATERAL (SELECT * FROM orders WHERE p.id = orders.productid ) ord  ON true 
;

SELECT 
*
FROM products p 
LEFT JOIN LATERAL (SELECT * FROM orders WHERE p.id = orders.productid ) ord  ON true 
WHERE 
	ord.id IS NULL
;

-- просто примеры
SELECT
	products.product_name,
	subquery1.category_name
FROM
	products,
	(
	SELECT
		categories.category_id,
		categories.category_name,
		COUNT(category_id) AS total
	FROM
		categories
	GROUP BY
		categories.category_id,
		categories.category_name
	) subquery1
WHERE
	subquery1.category_id = products.category_id
;


SELECT
	*
FROM
	laptop L1
CROSS JOIN LATERAL
         (
	SELECT
		MAX(price) max_price,
		MIN(price) min_price
	FROM
		Laptop L2
	JOIN Product P1 ON
		L2.model = P1.model
	WHERE
		maker = (
		SELECT
			maker
		FROM
			Product P2
		WHERE
			P2.model = L1.model
			)
		) X
;

SELECT
	*
FROM
	laptop L1
CROSS JOIN LATERAL
        (
	SELECT
		*
	FROM
		Laptop L2
	WHERE
		L1.model < L2.model
		OR (L1.model = L2.model
			AND L1.code < L2.code)
	ORDER BY
		model,
		code
	LIMIT 1
	) X
ORDER BY
	L1.model
;





