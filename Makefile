# Testing

default: help

test-tag:
	go test ./... -tags redis
.PHONY: test-tag

test-all:
	go test ./... -cover -v
.PHONY: test-all

test-pkg:
	go test ./test/testt/. -cover -v
.PHONY: test-pkg

test-func1:
	go test ./test/testt/. -run=Face -cover -v
.PHONY: test-func1

test-func2:
	go test ./... -run=Face -cover -v
.PHONY: test-func2

# Benchmark

bench-all:
	go test ./... -run=^# -bench=. -cover -v
.PHONY: bench-all

bench-pkg:
	go test ./sample/sort/. -run=^# -bench=. -cover -v
.PHONY: bench-pkg

bench-func1:
	go test ./sample/sort/. -run=^# -bench=SortBitwise -cover -v
.PHONY: bench-func1

bench-func2:
	go test ./... -run=^# -bench=SortBitwise -cover -v
.PHONY: bench-func2

# FMT & IMPORT
fmt:
	go fmt ./... && goimports -w .
.PHONY: fmt

# Linters
lint:
	golangci-lint run -c .golangci.yml
.PHONY: lint

h:
	@echo "Usage: make [target]"
	@echo "  target is:"
	@echo "		h			- Вывод этой документации"
	@echo "		test-tag		- Запуск тестов по тегу"
	@echo "		test-all		- Запуск всех тестов"
	@echo "		test-pkg		- Запуск тестов пакета"
	@echo "		test-func1		- Запуск теста функции вариант 1"
	@echo "		test-func2		- Запуск теста функции вариант 2"
	@echo "		bench-all		- Запуск всех benchmark тестов"
	@echo "		bench-pkg		- Запуск benchmark тестов пакета"
	@echo "		bench-func1		- Запуск теста benchmark функции вариант 1"
	@echo "		bench-func2		- Запуск теста benchmark функции вариант 2"
	@echo "		lint			- Запуск линтеров"
	@echo "		fmt			- форматирование и пути иппорта"
.PHONY: h
help: h
.PHONY: help

# https://www.zabbix.com/documentation/current/en/manual/installation/containers
zabbix-start:
	# 1. Create network dedicated for Zabbix component containers:
	docker network create --subnet 172.20.0.0/16 --ip-range 172.20.240.0/20 zabbix-net || true

	# 2. Start empty PostgreSQL server instance
	docker run --name postgres-server -t \
	-e POSTGRES_USER="zabbix" \
	-e POSTGRES_PASSWORD="zabbix_pwd" \
	-e POSTGRES_DB="zabbix" \
	--network=zabbix-net \
	--restart unless-stopped \
	-d postgres:latest

	# 3. Start Zabbix snmptraps instance
	docker run --name zabbix-snmptraps -t \
	-v /zbx_instance/snmptraps:/var/lib/zabbix/snmptraps:rw \
	-v /var/lib/zabbix/mibs:/usr/share/snmp/mibs:ro \
	--network=zabbix-net \
	-p 162:1162/udp \
	--restart unless-stopped \
	-d zabbix/zabbix-snmptraps:alpine-7.0-latest

	# 4. Start Zabbix server instance and link the instance with created PostgreSQL server instance
	docker run --name zabbix-server-pgsql -t \
	-e DB_SERVER_HOST="postgres-server" \
	-e POSTGRES_USER="zabbix" \
	-e POSTGRES_PASSWORD="zabbix_pwd" \
	-e POSTGRES_DB="zabbix" \
	-e ZBX_ENABLE_SNMP_TRAPS="true" \
	--network=zabbix-net \
	-p 10051:10051 \
	--volumes-from zabbix-snmptraps \
	--restart unless-stopped \
	-d zabbix/zabbix-server-pgsql:alpine-7.0-latest

	# 5. Start Zabbix web interface and link the instance with created PostgreSQL server and Zabbix server instances
	docker run --name zabbix-web-nginx-pgsql -t \
	-e ZBX_SERVER_HOST="zabbix-server-pgsql" \
	-e DB_SERVER_HOST="postgres-server" \
	-e POSTGRES_USER="zabbix" \
	-e POSTGRES_PASSWORD="zabbix_pwd" \
	-e POSTGRES_DB="zabbix" \
	--network=zabbix-net \
	-p 443:8443 \
	-p 80:8080 \
	-v /etc/ssl/nginx:/etc/ssl/nginx:ro \
	--restart unless-stopped \
	-d zabbix/zabbix-web-nginx-pgsql:alpine-7.0-latest
.PHONY:zabbix-start

zabbix-stop:
	docker stop postgres-server || true
	docker rm postgres-server || true
	docker stop zabbix-snmptraps || true
	docker rm zabbix-snmptraps || true
	docker stop zabbix-server-pgsql || true
	docker rm zabbix-server-pgsql || true
	docker stop zabbix-web-nginx-pgsql || true
	docker rm zabbix-web-nginx-pgsql || true
.PHONY:zabbix-stop