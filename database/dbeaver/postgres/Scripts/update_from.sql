UPDATE Ware
SET
	Ware.Company_ID = Company.ID
FROM Ware
  INNER JOIN Company ON Ware.Brend = Company.Name;

UPDATE Goods
	INNER JOIN Goods5 ON Goods.ID = Goods5.ID
SET
	Goods.Supplier_ID = Goods5.Supplier_ID;

UPDATE Goods
	INNER JOIN Goods5 ON Goods.ID = Goods5.ID
SET
	Goods.Valuta_ID = Goods5.Valuta_ID
WHERE
  [Conditon];

 -- psql
 
UPDATE orders
SET
	productcount = 100,
	price = 100.100
FROM customers, products
WHERE
	orders.customerid = customers.id AND customers.firstname = 'Sam'  
	AND orders.productid = products.id AND products.company = 'HTC'
;

UPDATE orders
SET
	productcount = products.productcount,
	price = products.price
FROM customers, products
WHERE
	orders.customerid = customers.id AND customers.firstname = 'Sam'  
	AND orders.productid = products.id AND products.company = 'HTC'
;

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
