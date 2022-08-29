package main

import (
	"context"
	"log"

	"github.com/kshamiev/urok/database/manticore/manti"
	"github.com/kshamiev/urok/database/manticore/typs"
)

func main() {
	if err := manti.Init(); err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()

	qu := `
	SELECT * 
	FROM documents 
	WHERE MATCH('Дом')
	ORDER BY updated_at desc 
	LIMIT 0,3 
	OPTION ranker=proximity, cutoff=0, retry_count=0, retry_delay=0;
	`
	data := typs.NewDocuments(100)
	if err := manti.SearchData(ctx, data, qu); err != nil {
		log.Fatalln(err)
	}

	qu = `
	SELECT count(*) 
	FROM documents 
	WHERE MATCH('Дом')
	OPTION ranker=proximity, cutoff=0, retry_count=0, retry_delay=0;
	`
	cnt, err := manti.SearchCount(ctx, qu)
	if err != nil {
		log.Fatalln(err)
	}

	typs.Dumper(data, cnt)
}
