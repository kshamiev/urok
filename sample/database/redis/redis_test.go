package redis

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Object struct {
	Str string
	Num int
}

func TestClient(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.TODO()
	key := "mykey"

	for i := 1; i < 21; i++ {
		obj := &Object{
			Str: "mystring" + strconv.Itoa(i),
			Num: 40 + i,
		}
		cacheObj, err := json.Marshal(obj)
		if err != nil {
			t.Fatal(err)
		}
		if err := redisClient.Set(ctx, key+strconv.Itoa(i), cacheObj, time.Hour).Err(); err != nil {
			t.Fatal(err)
		}
	}

	var wanted Object
	for i := 1; i < 21; i++ {
		if data, err := redisClient.Get(ctx, key+strconv.Itoa(i)).Bytes(); err == nil {
			err = json.Unmarshal(data, &wanted)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(wanted)
		} else {
			t.Fatal(err)
		}
	}

}

func TestRing(t *testing.T) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	ctx := context.TODO()
	key := "mykey"

	for i := 1; i < 21; i++ {
		obj := &Object{
			Str: "mystring" + strconv.Itoa(i),
			Num: 40 + i,
		}
		cacheObj, err := json.Marshal(obj)
		if err != nil {
			t.Fatal(err)
		}
		if err := ring.Set(ctx, key+strconv.Itoa(i), cacheObj, time.Hour).Err(); err != nil {
			t.Fatal(err)
		}
	}

	var wanted Object
	for i := 1; i < 21; i++ {
		if data, err := ring.Get(ctx, key+strconv.Itoa(i)).Bytes(); err == nil {
			err = json.Unmarshal(data, &wanted)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(wanted)
		} else {
			t.Fatal(err)
		}
	}

}

func TestCache(t *testing.T) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	ctx := context.TODO()
	key := "mykey"

	for i := 1; i < 21; i++ {
		obj := &Object{
			Str: "mystring" + strconv.Itoa(i),
			Num: 40 + i,
		}
		if err := mycache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   key + strconv.Itoa(i),
			Value: obj,
			TTL:   time.Hour,
		}); err != nil {
			t.Fatal(err)
		}
	}

	var wanted Object
	for i := 1; i < 21; i++ {
		if mycache.Exists(ctx, key+strconv.Itoa(i)) {
			if err := mycache.Get(ctx, key+strconv.Itoa(i), &wanted); err == nil {
				t.Log(wanted)
			} else {
				t.Fatal(err)
			}
		} else {
			t.Log("not exists " + strconv.Itoa(i))
		}
	}
}
