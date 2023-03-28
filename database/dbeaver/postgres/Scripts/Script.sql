SELECT * FROM users u WHERE (login = 'xTrKFop17W' AND description = 'KYi54DGgMs' OR login = 'rFIB0MLzJv' AND description = '1HzE0H4NZQ');
SELECT * FROM users u WHERE (login = 'xTrKFop17W' AND description = 'KYi54DGgMs') OR ( login = 'rFIB0MLzJv' AND description = '1HzE0H4NZQ');

UPDATE orders
SET
	(productcount, price) = 
(
SELECT 50, 50.50 
FROM orders AS o
	INNER JOIN customers ON o.customerid = customers.id 
	INNER JOIN products ON o.productid = products.id 
WHERE
	customers.firstname = 'Sam'
	AND products.company = 'HTC'
)
;
		

UPDATE orders
SET
	productcount = tt.productcount,
	price = tt.price
FROM (
SELECT 70 productcount, 70.70 price, customers.firstname, products.company
FROM orders AS o
	INNER JOIN customers ON o.customerid = customers.id 
	INNER JOIN products ON o.productid = products.id 
) AS tt
WHERE
	tt.firstname = 'Sam'
	AND tt.company = 'HTC'
;

