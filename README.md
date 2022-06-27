# urok

urok

### go.mod

добавление replace в go.mod

    go mod edit -replace github.com/user_name/repo_name=local_path

удаление replace из go.mod

    go mod edit -dropreplace github.com/user_name/repo_name

### Резервная копия и восстановление БД

```text
database restore    
psql -h "$(PG_HOST)" -p "$(PG_PORT)" -U $(PG_USER) -w -d $(PG_NAME) -f bin/dump.sql
psql -h localhost -p 5432 -U postgres -w -d urok -f database/dbeaver/postgres/Scripts/dump.sql

database full dump    
pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -w -d $(PG_NAME) -f bin/dump.sql
pg_dump -F p -h localhost -p 5432 -U postgres -w -d urok -f database/dbeaver/postgres/Scripts/dump.sql

database schema dump
pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -s -w -d $(PG_NAME) -f bin/dump.sql
pg_dump -F p -h localhost -p 5432 -U postgres -s -w -d test -f database/dbeaver/postgres/Scripts/dump_s.sql

database data dump
pg_dump -F p -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -a -w -d $(PG_NAME) -f bin/dump.sql
pg_dump -F p -h localhost -p 5432 -U postgres -a -w -d test -f database/dbeaver/postgres/Scripts/dump_d.sql
```