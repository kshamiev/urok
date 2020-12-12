# Инициализация
DBHOST := "localhost"
DBPORT := "5432"
DBUSER := "postgres"
DBPASS := "postgres"
DBNAME := "million"
DBMIG := "migrations"

DIR := $(realpath -m .)
DBDSN := "host=$(DBHOST) port=$(DBPORT) user=$(DBUSER) password=$(DBPASS) dbname=$(DBNAME) sslmode=disable"

default: help

# Инженеринг protobuf files
proto:
	@cd $(DIR) && protoc --go_out=plugins=grpc:. *.proto;
.PHONY: proto

# Инженеринг типов на основе БД
db:
	@cd $(DIR) && go run generate/database/main.go
	@cd $(DIR) && go fmt ./... && goimports -w .;
	@cd $(DIR)/typs && go build;
.PHONY: db

# Инженеринг типов proto
pb:
	@if [ -f typs/advanced.go ]; then \
		rm typs/advanced.go; \
	fi
	@cd $(DIR) && go run generate/proto/main.go;
	@cd $(DIR) && protoc --proto_path=pb -I=thirdparty --go_out=plugins=grpc:pb --grpc-gateway_out=logtostderr=true:pb --swagger_out=logtostderr=true:pb pb/*.proto
	@cd $(DIR) && go fmt ./... && goimports -w .;
	@cd $(DIR)/typs && go build;
.PHONY: pb

# database dump
dbdump:
	pg_dump -F p -h $(DBHOST) -p $(DBPORT) -U $(DBUSER) -w -f dump.sql -d $(DBNAME)
.PHONY: dbdump

# database restore
dbinit:
	psql -h "$(DBHOST)" -p "$(DBPORT)" -U $(DBUSER) -w -f dump.sql -d $(DBNAME)
.PHONY: dbinit

# Создание шаблона миграции
mig:
	@gsmigrate --dir=${DBMIG} --drv="postgres" --dsn=${DBDSN} create tpl;
.PHONY: mig

# Статус миграции
mig-st:
	gsmigrate --dir=${DBMIG} --drv="postgres" --dsn=${DBDSN} status;
.PHONY: mig-st

# Миграция на одну позицию вниз
mig-down:
	gsmigrate --dir=${DBMIG} --drv="postgres" --dsn=${DBDSN} down;
.PHONY: mig-down

# Миграция вверх до конца
mig-up:
	gsmigrate --dir=${DBMIG} --drv="postgres" --dsn=${DBDSN} up;
.PHONY: mig-up

# Инженеринг моделей по существующей структуре БД
models:
	sqlboiler -c sqlboiler.yaml -p models -o models --no-auto-timestamps --no-tests --wipe psql
.PHONY: models

# FMT & GOIMPORT
fmt:
	go fmt ./... && goimports -w .
.PHONY: fmt

# Linters
lint:
	golangci-lint run
.PHONY: lint

# Help
h:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "    h			- Вывод этой документации"
	@echo "    proto		- Инженеринг protobuf файлов"
	@echo "    pb			- Инженеринг сопоставления типов и protobuf файлов"
	@echo "    mig			- Создание шаблона миграции "
	@echo "    mig-st		- Статус миграции"
	@echo "    mig-down		- Миграция на одну позицию вниз"
	@echo "    mig-up		- Миграция вверх до конца"
	@echo "    dbinit		- Восстановление БД из дампа dump.sql (БД должна существовать)"
	@echo "    dbdump		- Создание дампа БД dump.sql"
	@echo "    models		- Инженеринг моделей по существующей структуре БД"
	@echo "    lint		- Линтеры"
	@echo "    fmt			- Форматирование кодовой базы"
.PHONY: h
help: h
.PHONY: help
