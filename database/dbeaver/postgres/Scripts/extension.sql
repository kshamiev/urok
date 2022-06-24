select * from pg_available_extensions;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "moddatetime";

DROP EXTENSION IF EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS "moddatetime";

