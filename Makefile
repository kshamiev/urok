# Testing

default: help

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
