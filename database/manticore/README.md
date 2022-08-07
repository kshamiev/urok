## Manticore

## Install System & Start

```text
sudo apt remove manticore*

wget https://repo.manticoresearch.com/manticore-repo.noarch.deb
sudo dpkg -i manticore-repo.noarch.deb
sudo apt update
sudo apt install manticore manticore-columnar-lib

info from indexer
sudo -u manticore indexer
apt-file find libmysqlclient.so.20

sudo apt-get install libmysqlclient20 libodbc1 libpq5 libexpat1

systemctl is-enabled manticore
systemctl enable manticore
systemctl disable manticore

systemctl status manticore
systemctl stop manticore
systemctl restart manticore
systemctl start manticore
 
searchd --status
sudo indexer --all --rotate
gosu manticore indexer --all --rotate

debug
sudo journalctl --unit manticore
sudo journalctl -xe
```

Можно работать только в одном режиме из двух.
В режиме индексов реального времени.
Либо в режиме простого индекса
При работе в режиме простого индекса.

1) Только декларативный подход.
2) Репликации на уровне manticore недоступны.

Возможен ли поиск во время перестройки простого индекса.