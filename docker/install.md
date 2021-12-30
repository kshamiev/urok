### Понятия и термины

#### Images (образы)

Схемы нашего приложения, которые являются основой контейнеров. В примере выше мы использовали
команду ```docker pull busybox``` чтобы скачать образ busybox.

#### Containers (контейнеры)

Создаются на основе образа и запускают само приложение. Мы создали контейнер командой ```docker run```, и использовали
образ busybox, скачанный ранее. Список запущенных контейнеров можно увидеть с помощью команды ```docker ps```.

#### Docker Daemon (демон Докера)

Фоновый сервис, запущенный на хост-машине, который отвечает за создание, запуск и уничтожение Докер-контейнеров. Демон
это процесс запущенный в ОС, с которой взаимодействует клиент.

#### Docker Client (клиент Докера)

Утилита командной строки, которая позволяет пользователю взаимодействовать с демоном. Существуют другие формы клиента,
например, Kitematic, с графическим интерфейсом.

#### Docker Hub - Регистр Докер-образов. Грубо говоря, архив всех доступных образов.

https://hub.docker.com/

    Если нужно, то можно содержать собственный регистр и использовать его для получения образов.

Важно понимать разницу между базовыми и дочерними образами:

    Base images (базовые образы) — это образы, которые не имеют родительского образа. 
		Обычно это образы с операционной системой, такие как ubuntu, busybox или debian.
    Child images (дочерние образы) — это образы, построенные на базовых образах и обладающие дополнительной функциональностью.

Существуют официальные и пользовательские образы, и любые из них могут быть базовыми и дочерними.

    Официальные образы — это образы, которые официально поддерживаются командой Docker. 
		Обычно в их названии одно слово. В списке выше python, ubuntu, busybox и hello-world — базовые образы.
    Пользовательские образы — образы, созданные простыми пользователями вроде меня и вас.
		Они построены на базовых образах. Обычно, они называются по формату user/image-name.

### INSTALL

#### Удаление утсановленного ранее докера

    sudo apt-get remove docker docker-engine docker.io containerd runc

#### Установка докера

(https://docs.docker.com/install/linux/docker-ce/ubuntu/)

Set up the repository

    sudo apt-get update

    sudo apt-get install \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg-agent \
        software-properties-common

    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

    sudo apt-key fingerprint 0EBFCD88

    sudo add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
       $(lsb_release -cs) \
       stable"

Install Docker Engine - Community

    sudo apt-get update

    sudo apt-get install docker-ce docker-ce-cli containerd.io

    sudo docker run hello-world

Еще нужно установленные

- python --version
- pip --version

Решение проблемы с правами вызова команд докера

    sudo groupadd docker
    
    sudo usermod -aG docker $USER

#### ПОДКЛЮЧЕНИЕ КОНТЕЙНЕРА К БД НА ХОСТЕ

	Нвстройка БД

	узнаем ip адресс хоста
	$ hostname -I | cut -d ' ' -f1

	прописываем ip в конфигурацию для прослушивания
	/etc/postgresql/10/main/postgresql.conf
	listen_addresses = 'localhost,192.168.0.82'  # можно просто звездочку (*) напечатать
	
	в конец конфигурации добавить (для решения проблемы прав доступа)
	/etc/postgresql/10/main/pg_hba.conf
	host     all             all             172.17.0.0/8            md5
	
	перезапускаем Postgres
	$ service postgresql restart

	Ну и далее в скриптах приложения вместо localhost указываем локальный айпишник.
	При запуске контейнера для подключения к БД ничего указывать не надо.

#### ДОСТУП В КОНТЕЙНЕР ПО http (8080,80,443,...)

	Нужно при запуске контенйера пробросить нужные порты (-p 8080:8080)
	
	В самом приложении в хосте к вебсерверу нужно прописывать 0.0.0.0 вместо localhost
