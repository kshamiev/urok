## Manticore

## install Docker

```dockerfile
docker build --no-cache -t kshamiev/manticore:v1 -f manticore.Dockerfile .
docker push kshamiev/assembly:v1

docker run --rm -it kshamiev/manticore:v1

docker run --rm --name manticore \
	-v $(pwd)/tmp:/var/lib/manticore \
	-v $(pwd)/manticore.conf:/etc/manticoresearch/manticore.conf \
	-p 9308:9308 \
	-p 9312:9312 \
	-p 9306:9306 \
	-d kshamiev/manticore:v1 searchd --nodetach

docker exec -it manticore bash
docker logs manticore

searchd --status
searchd --stop
sudo -u manticore indexer documents_main --rotate
sudo -u manticore indexer documents_delta --rotate
sudo -u manticore indexer documents_rt --rotate
gosu manticore indexer --all --rotate

mysql -P9306 -h0
SHOW TABLES;
DESCRIBE users_main_idx;
SHOW META;
RELOAD INDEXES;
```

 1 sql_attr_bigint = id
 2 sql_attr_string = title
 3 sql_attr_uint = release_year
 4 sql_attr_float = price
 5 sql_attr_timestamp = created_at
 6 sql_attr_bool = is_flag
 7 sql_attr_json = data

1 rt_attr_bigint = id
2 rt_attr_string = title
3 rt_attr_uint = release_year
4 rt_attr_float = price
5 rt_attr_timestamp = created_at
6 rt_attr_bool = is_flag
7 rt_attr_json = data
