## Встраивание внешней статики в приложение.
## Программой находящейся внутри пакета генерируются исходные коды в которые
## встраиваются статические файлы с данными.
## EMBED_STATIC_RESOURCES - Группы ресурсов и папками самих ресурсов.
##     Формат: group_name1:path/to/folder1,group_name2:/absolut/path/to/folder2
## Относительные пути берутся от текущей рабочей директории
export EMBED_STATIC_RESOURCES=www:static_sample
embed:
	@echo "Встроенные ресурсы"
	@go run embedder/main.go
.PHONY: embed
