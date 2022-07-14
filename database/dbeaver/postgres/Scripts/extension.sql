select * from pg_available_extensions;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "moddatetime";
CREATE EXTENSION IF NOT EXISTS "ltree";

-- расширения для изучения и работы с индексами типа btree
CREATE EXTENSION IF NOT EXISTS "pageinspect";
CREATE EXTENSION IF NOT EXISTS "amcheck";

DROP EXTENSION IF EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS "moddatetime";
DROP EXTENSION IF EXISTS "ltree";

