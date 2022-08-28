package main

import (
	"fmt"
	"log"

	"github.com/kshamiev/urok/database/manticore/manti"
	"github.com/kshamiev/urok/database/manticore/typs"
)

func main() {

	if err := manti.Init(); err != nil {
		log.Fatalln(err)
	}

	qu := `
	SELECT * 
	FROM documents 
	WHERE MATCH('Дом')
	ORDER BY updated_at desc 
	LIMIT 0,3 
	OPTION ranker=proximity, cutoff=0, retry_count=0, retry_delay=0;
	`

	data := typs.NewDocuments(100)

	if err := manti.SearchCustom(data, qu); err != nil {
		log.Fatalln(err)
	}

	typs.Dumper(data)
	fmt.Println(data.Data[0].Price.String())
	fmt.Println(data.Data[0].UpdatedAt.String())
}
