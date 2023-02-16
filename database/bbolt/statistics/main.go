package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("database/bbolt/data.db", 0666, &bolt.Options{
		Timeout:  time.Second * 10,
		ReadOnly: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	go func() {
		// Grab the initial stats.
		prev := db.Stats()
		for {
			// Wait for 10s.
			time.Sleep(10 * time.Second)

			// Grab the current stats and diff them.
			stats := db.Stats()
			diff := stats.Sub(&prev)

			// Encode stats to JSON and print to STDERR.
			json.NewEncoder(os.Stderr).Encode(diff)

			// Save stats for the next loop.
			prev = stats
		}
	}()
	time.Sleep(time.Hour)
}
