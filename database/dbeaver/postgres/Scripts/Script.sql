INSERT INTO films
(title, description, category_id, release_year, price)
SELECT title, description, category_id, release_year, price
FROM films WHERE id IN (349, 696);
