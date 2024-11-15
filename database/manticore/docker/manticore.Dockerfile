FROM ubuntu:22.04

LABEL maintainer="Konstantin Shamiev aka ilosa <konstantin@shamiev.ru>"

RUN apt-get update \
    && apt-get -y install mysql-client \
    && apt-get -y install postgresql-client-14 \
    && apt-get -y install wget

RUN mkdir -p /var/run/manticore \
    && mkdir -p /usr/local/lib/manticore \
    && mkdir -p /var/lib/manticore/replication \
    && cd tmp

RUN wget https://repo.manticoresearch.com/repository/manticoresearch_jammy/dists/manticore_5.0.2-220530-348514c86_amd64.tgz https://repo.manticoresearch.com/repository/manticoresearch_jammy/dists/jammy/main/binary-amd64/manticore-columnar-lib_1.15.4-220522-2fef34e_amd64.deb \
    && tar -xvf manticore_5.0.2-220530-348514c86_amd64.tgz \
    && dpkg -i manticore-*5.0.2-220530-348514c86_amd64.deb manticore-columnar-lib_1.15.4-220522-2fef34e_amd64.deb

RUN wget https://repo.manticoresearch.com/repository/morphology/en.pak.tgz \
    && tar -xvf en.pak.tgz \
    && mv en.pak /usr/share/manticore

RUN wget https://repo.manticoresearch.com/repository/morphology/ru.pak.tgz \
    && tar -xvf ru.pak.tgz \
    && mv ru.pak /usr/share/manticore

RUN rm ./*.deb && rm ./*.tgz

EXPOSE 9306
EXPOSE 9308
EXPOSE 9312
