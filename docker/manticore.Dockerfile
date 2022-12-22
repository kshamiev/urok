FROM ubuntu:22.04 as build1

RUN apt-get update \
    && apt-get -y install wget \
    && cd /tmp \
    && wget https://go.dev/dl/go1.17.13.linux-amd64.tar.gz \
    && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.13.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

WORKDIR /home/app
COPY . .

RUN rm -rdf /home/app/bin \
    && go build -o bin/app .


FROM ubuntu:22.04

RUN apt-get update \
    && apt-get -y install mysql-client \
    && apt-get -y install postgresql-client-14 \
    && apt-get -y install wget

RUN mkdir -p /var/run/manticore \
    && mkdir -p /usr/local/lib/manticore \
    && mkdir -p /var/lib/manticore/replication \
    && mkdir /tmp/temp && cd /tmp/temp

RUN wget https://repo.manticoresearch.com/repository/manticoresearch_jammy/dists/manticore_5.0.2-220530-348514c86_amd64.tgz https://repo.manticoresearch.com/repository/manticoresearch_jammy/dists/jammy/main/binary-amd64/manticore-columnar-lib_1.15.4-220522-2fef34e_amd64.deb \
    && tar -xvf manticore_5.0.2-220530-348514c86_amd64.tgz \
    && dpkg -i manticore-*5.0.2-220530-348514c86_amd64.deb manticore-columnar-lib_1.15.4-220522-2fef34e_amd64.deb

RUN wget https://repo.manticoresearch.com/repository/morphology/en.pak.tgz \
    && tar -xvf en.pak.tgz \
    && mv en.pak /usr/share/manticore

RUN wget https://repo.manticoresearch.com/repository/morphology/ru.pak.tgz \
    && tar -xvf ru.pak.tgz \
    && mv ru.pak /usr/share/manticore

RUN rm -rdf /tmp/temp

WORKDIR /home/app
RUN mkdir etc
COPY --from=build1 /home/app/bin bin
EXPOSE 9306
EXPOSE 9308
EXPOSE 9312
EXPOSE 7070

CMD bin/app -c config.yml
