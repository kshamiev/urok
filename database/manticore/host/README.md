## Manticore

## Install System & Start

```text
systemctl status manticore
systemctl restart manticore

searchd --status
sudo -u manticore indexer --all --rotate
systemctl restart manticore
sudo -u manticore indexer documents_main --rotate
sudo -u manticore indexer documents_delta --rotate

mysql -P9306 -h0
SHOW TABLES;
DESCRIBE users_main_idx;
SHOW META;
RELOAD INDEXES;

searchd --config /etc/manticoresearch/manticore.conf --stop
searchd --config dev.conf --status
sudo -u manticore searchd --config dev.conf
sudo -u manticore indexer --config dev.conf --all --rotate
```

Можно работать только в одном режиме из двух.
В режиме индексов реального времени.
Либо в режиме простого индекса
При работе в режиме простого индекса.

1) Только декларативный подход.
2) Репликации недоступны.

Возможен ли поиск во время перестройки простого индекса.