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
gosu manticore indexer --all --rotate

mysql -P9306 -h0
SHOW TABLES;
DESCRIBE users_main_idx;
SHOW META;
RELOAD INDEXES;
```
