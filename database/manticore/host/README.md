## Manticore

## Install System & Start

```text
systemctl status manticore
systemctl restart manticore

searchd --status
searchd --stop
sudo -u manticore indexer documents_main --rotate
sudo -u manticore indexer documents_delta --rotate
sudo -u manticore indexer --all --rotate

mysql -P9306 -h0
SHOW TABLES;
DESCRIBE users_main_idx;
SHOW META;
RELOAD INDEXES;
```

Можно работать только в одном режиме из двух.
В режиме индексов реального времени.
Либо в режиме простого индекса
При работе в режиме простого индекса.

1) Только декларативный подход.
2) Репликации недоступны.

