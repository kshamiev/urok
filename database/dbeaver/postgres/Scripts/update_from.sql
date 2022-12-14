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
