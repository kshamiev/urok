INSERT INTO users (login, description, json_struct, json_string) SELECT random_string(10), random_string(10), '{}', '{}' FROM users;
SELECT count(*) FROM users u ;
-- 1441883

-- GO-PG
INSERT INTO users (id, login, description, price, json_struct, json_string) VALUES 
(1, 'testLogin', 'number g: yWl8brH7up', 45.67, '{}', '{}'),
(2,'gpscdIEk','number g: 7DVWGSmfRy', 73.81, '{}', '{}'),
(3,'v3iwypkK','number g: A52oiAje8S', 29.51, '{}', '{}')
ON CONFLICT (id) DO UPDATE SET description = EXCLUDED.description, price = EXCLUDED.price RETURNING *;

-- BOIL

-- updateOnConflict true

INSERT INTO "users" ("id", "login", "description", "price") VALUES 
($1,$2,$3,$4) 
ON CONFLICT ("id") DO UPDATE SET "description" = EXCLUDED."description", "price" = EXCLUDED."price" RETURNING *;

-- updateOnConflict false

INSERT INTO "users" ("id", "login", "description", "price") VALUES 
($1,$2,$3,$4)
ON CONFLICT DO NOTHING RETURNING *;

-- DEV

INSERT INTO users (login, description, price, json_struct, json_string) VALUES 
('testLogin', 'number g: yWl8brH7up', 45.67, '{}', '{}'),
('gpscdIEk','number g: 7DVWGSmfRy', 73.81, '{}', '{}'),
('v3iwypkK','number g: A52oiAje8S', 29.51, '{}', '{}')
ON CONFLICT (id) DO UPDATE SET description = EXCLUDED.description, price = EXCLUDED.price RETURNING *;
