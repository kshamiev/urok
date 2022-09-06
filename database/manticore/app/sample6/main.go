package main

import (
	"context"
	"log"

	"github.com/kshamiev/urok/database/manticore/manti"
	"github.com/kshamiev/urok/database/manticore/typs"
	"github.com/kshamiev/urok/debug"
)

func main() {
	if err := manti.Init(); err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()

	data := typs.NewDocuments(100)

	s := manti.NewSearch("дом", "documents")
	s.Select("*")
	s.Order("updated_at desc")
	s.Limit(0, 3)
	if err := s.Fetch(ctx, data); err != nil {
		log.Fatalln(err)
	}
	debug.Dumper(data)
}
