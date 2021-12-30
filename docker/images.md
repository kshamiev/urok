#### Получение нужного образа контейнера из репозитория

	docker pull busybox
	docker pull alpine
	docker pull golang:1.17
	docker pull ubuntu:20.04

#### Просмотр списка полученных и доступных образов в системе

	docker images -a

#### Удаление образа

	docker rmi ImageID1 ImageID2
	docker rmi $(docker images -f dangling=true -q)

#### Создание нового образа

(в корне проекта с Dockerfile)

    docker build --rm -t kshamiev/sungora:v1.10.110 .

#### Поиск образа на хабе

	docker search


