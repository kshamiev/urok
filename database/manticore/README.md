### Manticore

```dockerfile
docker run --rm --name manticore \
    -e MCL=1 \
	-e QUERY_LOG_TO_STDOUT=true \
	-v $(pwd)/etc/manticore.conf:/etc/manticoresearch/manticore.conf \
	-v $(pwd)/data:/var/lib/manticore \
	-p 9308:9308 \
	-p 9312:9312 \
	-d manticoresearch/manticore

	-p 5432:5432 \

docker run --name manticore \
	-e QUERY_LOG_TO_STDOUT=true \
	-v $(pwd)/etc/manticore.conf:/etc/manticoresearch/manticore.conf \
	-v $(pwd)/data:/var/lib/manticore \
	-p 9306:9306 \
	-p 9308:9308 \
	-p 9312:9312 \
	-d manticoresearch/manticore

docker run --rm --name manticore \
	-v $(pwd)/etc/manticore.conf:/etc/manticoresearch/manticore.conf \
	-v $(pwd)/data:/var/lib/manticore \
	-p 192.168.0.101:9306:9306 \
	-p 192.168.0.101:9308:9308 \
	-p 192.168.0.101:9312:9312 \
	-d manticoresearch/manticore

gosu manticore indexer --all --rotate

```
