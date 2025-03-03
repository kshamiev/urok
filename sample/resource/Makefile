# Инициализация
export DIR   := $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

## Встраивание внешней статики в приложение.
## Программой находящейся внутри пакета генерируются исходные коды в которые
## встраиваются статические файлы с данными.
embed:
	@echo "Встроенный ресурс www"
	@if [ -d "www" ]; then \
	export EMBEDDER_STATIC_BASE_DIR="${DIR}"; \
	export EMBEDDER_STATIC_RESOURCES="www:www"; \
	go run embedder/main.go; \
	fi
.PHONY: embed
