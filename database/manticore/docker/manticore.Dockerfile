FROM ubuntu:22.04

LABEL maintainer="Konstantin Shamiev aka ilosa <konstantin@shamiev.ru>"

RUN apt-get update \
    && apt-get install make \
    && apt-get -y install mc \
    && apt-get -y install mysql-client \
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

RUN wget https://go.dev/dl/go1.17.13.linux-amd64.tar.gz \
    && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.13.linux-amd64.tar.gz

RUN rm ./*.deb && rm ./*.tgz && rm ./*.gz

ENV PATH=$PATH:/usr/local/go/bin

EXPOSE 9306
EXPOSE 9308
EXPOSE 9312

# docker build --no-cache -t kshamiev/manticore:v1 -f manticore.Dockerfile .
# docker push kshamiev/assembly:v1
# docker run --rm -it kshamiev/assembly:v1