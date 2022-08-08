## Manticore

## install Docker

```dockerfile
docker run --rm --name manticore \
	-v $(pwd)/tmp:/var/lib/manticore \
	-v $(pwd)/manticore.conf:/etc/manticoresearch/manticore.conf \
	-p 9308:9308 \
	-p 9312:9312 \
	-d manticoresearch/manticore

docker exec -it manticore bash

gosu manticore indexer --all --rotate
```
