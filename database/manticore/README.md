## Manticore

## install Docker

```dockerfile
docker run --name manticore \
    -e MCL=1 \
    -e QUERY_LOG_TO_STDOUT=true \
    -v $(pwd)/etc/manticore.conf:/etc/manticoresearch/manticore.conf \
	-v $(pwd)/tmp:/var/lib/manticore \
	-p 192.168.0.101:9308:9308 \
	-p 192.168.0.101:9312:9312 \
	-d manticoresearch/manticore

docker run --name manticore \
    -v $(pwd)/etc/manticore.conf:/etc/manticoresearch/manticore.conf \
	-v $(pwd)/tmp:/var/lib/manticore \
	-p 9308:9308 \
	-p 9312:9312 \
	-d manticoresearch/manticore

docker exec -it manticore bash

gosu manticore indexer --all --rotate
searchd --status
```

## install System