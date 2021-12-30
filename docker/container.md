#### ЗАПУСК КОНТЕЙНЕРА

	docker run -d --rm -p 8080:8080 --name test kshamiev/sungora:v1.10.106
	docker run -it --rm --net host --name test kshamiev/sungora:v1.10.108
	docker run -d --rm --net host --name test kshamiev/sungora:v1.10.110

	-d 			запуск в режиме демона (отпустить консоль)
	-P 			все открытые порты публичными и случайными
	--name		имя, которое мы хотим дать контейнеру
	--net		сеть в котрой будет запущен контейнер (сеть должна существовать)
	-p 8080:80	проброс порта 80 (порт контейнера) на порт 8080 (порт родительской ОС)
    -it         запуск с терминалом управления
    --rm        Запуск с последующим удалением после завершения работы
    --volume /tmp/data:/home/app/data монтирование папки хоста на папку контейнера

	docker run --name test1 --net qwerty -dp 8080:8080 kshamiev/test1
	docker run --name psql2 --net qwerty -dp 5433:5432 -e POSTGRES_PASSWORD=postgres postgres:10

#### РАБОТА С РАБОТАЮЩИМИ КОНТЕЙНЕРАМИ

Список всех запущенных контейнеров (-a завершенные контейнеры)

	docker ps -a

Вход в работающий в фоне контейнер

	docker exec -it containerNameOrID sh
	docker exec -it containerNameOrID bash

	netstat -lp

Просмотр логов работы контейнера

	docker logs containerNameOrID

Остановить работающий в фоне (в режиме демона) контейнер

	docker stop containerNameOrID

Удаление контейнеров

	docker rm containerNameOrID containerNameOrID
	docker rm $(docker ps -a -q -f status=exited)

Просмотр используемых портов работающего контейнера

	docker port containerNameOrID

#### СЕТИ ДОКЕРА

Список всех сетей

	docker network ls

Подробная информация по указанной сети

	docker network inspect bridge

Создание новой сети

	docker network create NameNewNet

