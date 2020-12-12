# Million

base and share code from use project service

---
### install tools

    go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
    go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

    go get github.com/volatiletech/sqlboiler
    go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

    go get github.com/webnice/migrate/gsmigrate

    go get github.com/swaggo/swag/cmd/swag
    
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.30.0

Описание пакетов

    /models
        модели boiler для примитвной работы с БД в парадигме ОРМ
    /pb
        прототипы для общения между приложениями по GRPC
    /pkg
        инструментарий для реализации работы приложений 
    /typ
        простые типы общего служебного и специфичного назначения.
        (уникальные и значимые типы)
    /typs
        основные рабочие типы приложения
        - могут встраивать в себя модели boiler
        - имеют методы для передачи себя по GRPC
        - обеспечивает дополнитеьный уровень расширения общего функционала (валидация свойств...) 


---
### Test

    go run cmd/grpc/report.go

---
### GRPC

#### typs/generate
- Удаляем старую сборку (`typs/advanced.go`).
- Обьявляем или изменяем и регистрируем свои типы (`typs/generate/main.go`)
- По необходимости реализуем кастомные обработчики (`typs/generate/main.go`, `typs/adv.go`).
- Запускаем команду:

    make pb

Далее описываете и реализуете свой сервис

[описание инструмента](generate/proto/README.md "описание инструмента")

#### dop info

***protobuf files***

Компиляция находясь в целевом каталоге:

    protoc --go_out=plugins=grpc:. *.proto;
    protoc --go_out=plugins=grpc:. --grpc-gateway_out=:. --swagger_out=:. *.proto
    protoc --proto_path=pb -I=thirdparty --go_out=plugins=grpc:pb --grpc-gateway_out=logtostderr=true:pb --swagger_out=logtostderr=true:pb pb/*.proto
    
    protoc-gen-grpc-gateway --help
    protoc-gen-swagger --help 

***examples***

https://github.com/grpc/grpc-go/tree/master/examples

Полезная ссылка: https://github.com/grpc-ecosystem/grpc-gateway

---
### Kubernetes

    data/config-kubectl
    kubectl get pods
    kubectl logs -f [NAME|ID пода] --tail 100
    kubectl port-forward postgresql-0 5433:5432
    kubectl port-forward [NAME|ID пода] Plocal:Premote

---
### Swagger Документирование api

Для работы со свагером мы используем библиотеку: [swaggo](https://github.com/swaggo/swag#api-operation)

Описание документирования api:
<pre>
//+funcName godoc
//+@Summary Авторизация пользователя по логину и паролю (ldap).     пишем кратко о чем речь и что принимает на входе
// @Description Возвращается токен авторизации и пользователья      пишем что возвращает и возможно подробности
//+@Tags tagName                                                    группировка api запросов
//+@Router /page/page [post]                                        относительный роутинг от базового и метод
//+@Param name TARGET TYPE_REQUEST true "com"                       входящие параметры
//+@Success 200 {TYPE_RESPONSE} string "com"                        положительный ответ
//+@Failure 400 {TYPE_RESPONSE} request.Error "com"                 отрицательный ответ
//+@Failure 401 {TYPE_RESPONSE} request.Error "user unauthorized"   пользователь не авторизован
// @Accept json                                                     тип принимаемых данных
// @Produce json                                                    тип возвращаемых данных
// @Security ApiKeyAuth                                             запрос авторизованный по ключу или токену
</pre>

<pre>
+ Обязательные теги и теги по контексту (параметров может и не быть...)
TARGET          = header | path | query  | body | formData
TYPE_REQUEST    = string | int  | number | bool | file | userGolangStruct
TYPE_RESPONSE   = string | int  | number | bool | file | object | array
</pre>

Пример:
<pre>
// Login авторизация пользователя по логину и паролю ldap
// @Summary авторизация пользователя по логину и паролю (ldap).
// @Description возвращается токен авторизации
// @Tags Auth
// @Router /auth/login [post]
// @Param credentials body models.Credentials true "реквизиты доступа"
// @Success 200 {string} string "успешная авторизация"
// @Accept json                                                    
// @Produce json                                                   
// @Security ApiKeyAuth
</pre>

Проблемы:

* Не умеет работать с алиасами в импортах.
* Типы slice, map не поддерживаются для входных параметров (нужно оборачивать в отдельные типы) 

Принятые коды ответов:

- 200 Любой положительный ответ
- 301 Редирект (перманентный). Переход на другой запрос
- 302 Редирект (от логики). Переход на другой запрос
- 400 Ошибка работы с данными приложения
- 401 Пользователь не авторизован
- 403 Отказано в операции за отсутствием прав 
- 404 Данные по запросу не найдены
- 500 Ошибка работы сервера

Формат принимаемых и отдаваемых данных для API:

- Данные передаются в формате JSON

---
### Jaeger
https://www.jaegertracing.io/docs/1.20/getting-started/

Запуск в докере:

docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.20

http://localhost:16686

---
### Oracle
В терминале:

    sudo apt install libaio-dev
    
    mkdir -p /opt/oracle
    cd /opt/oracle
    wget https://download.oracle.com/otn_software/linux/instantclient/193000/instantclient-basic-linux.x64-19.3.0.0.0dbru.zip
    wget https://download.oracle.com/otn_software/linux/instantclient/193000/instantclient-sdk-linux.x64-19.3.0.0.0dbru.zip
    wget https://download.oracle.com/otn_software/linux/instantclient/193000/instantclient-sqlplus-linux.x64-19.3.0.0.0dbru.zip
    unzip instantclient-basic-linux.x64-19.3.0.0.0dbru.zip 
    unzip instantclient-sdk-linux.x64-19.3.0.0.0dbru.zip 
    unzip instantclient-sqlplus-linux.x64-19.3.0.0.0dbru.zip 
    
    mkdir -p /usr/lib/pkgconfig
    cd /usr/lib/pkgconfig
    создать файл /usr/lib/pkgconfig/oci8.pc со следующим содержимым:
    
    //==========================================BEG
    prefixdir=/opt/oracle/instantclient_19_3
    libdir=${prefixdir}
    includedir=${prefixdir}/sdk/include

    Name: OCI
    Description: Oracle database driver
    Version: 19.3
    Libs: -L${libdir} -lclntsh
    Cflags: -I${includedir}
    //==========================================END
    
    далее выполнить команду в этом же каталоге:
    pkg-config oci8.pc
    
    Прописать переменные окружения в /etc/environment
    
    PKG_CONFIG_PATH=/usr/lib/pkgconfig
    LD_LIBRARY_PATH=/opt/oracle/instantclient_19_3
    TNS_ADMIN=/opt/oracle/instantclient_19_3/network/admin
    
    sh -c "echo /opt/oracle/instantclient_19_3 >/etc/ld.so.conf.d/oracle-instantclient.conf"
    sudo ldconfig
    
    go get github.com/mattn/go-oci8
