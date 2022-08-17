package main

import (
	"fmt"

	"github.com/manticoresoftware/go-sdk/manticore"
)

func main() {
	cl := manticore.NewClient()
	// cl.SetServer("192.168.0.103", 9312)
	cl.SetServer("127.0.0.1", 9312)
	// cl.SetServer("localhost", 9312)
	// cl.SetServer("localhost", 9308)
	if _, err := cl.Open(); err != nil {
		fmt.Println(err)
		return
	}

	res, err := cl.Sphinxql("RELOAD INDEXES")
	fmt.Println(res, err)

	// data, err := cl.Json("search", req2)
	// fmt.Println(data, err)

	q := manticore.NewSearch("Дом", "documents", "")
	// q := manticore.NewSearch("мухомор", "users", "")
	q.Offset = 0
	q.Limit = 3
	q.SetSortMode(manticore.SortExtended, "updated_at desc")
	res2, err2 := cl.RunQuery(q)
	fmt.Println(res2, err2)
}

var req1 = `
{
  "index": "users",
  "query": {
    "match": {
      "title, description": "дом"
    }
  },
  "limit": 10
}
`

var req2 = `
{
  "index": "users",
  "query": {
    "match": {
      "title,description": "Дом"
    }
  },
  "sort": [ { "updated_at":"desc" } ],
  "limit": 3
}
`

/*
10 000 000
main

2022-08-16 10:47:50.444
2022-08-16 11:06:50.444
186 000


range 1 000 000
2022-08-15 14:24:06.320
2022-08-15 14:34:59.649
10:53

range 100 000
2022-08-15 13:16:22.528
2022-08-15 13:26:50.872
10:28

range 10 000
2022-08-15 13:55:09.463
2022-08-15 14:06:27.283
11:18

delta

2022-08-15 13:53:14.961
2022-08-15 13:53:33.508

killlist = SELECT id FROM documents WHERE updated_at >=  (SELECT created_at FROM deltabreaker WHERE index_name='delta')

sql_query_killlist = \
        SELECT id FROM documents WHERE updated_ts>=@last_reindex UNION \
        SELECT id FROM documents_deleted WHERE deleted_ts>=@last_reindex
}


Total: 3
Total found: 21614


'дом' (Docs:21614, Hits:38903)



Прохор Пастушенко, [05.08.2022 09:48]
@s_a_nikolaev
Добрый день!
Подскажите пожалуйста, по "тюнингу" движка.
У нас такой кейс: 2 таблицы по 203 колонки (id bigint, group string, doc_id string и 200 text indexed) + engine='columnar' + min_infix_len=2.
Питоном (для тестов) мы генерируем .sql файлики с INSERT`ами. Одна инструкция INSERT представляет из себя вставку пачкой в 400 VALUES (в каждое text indexed поле мы пишем рандомное предложение в 100 слов). "Вес" одной инструкции - 120МБ.
Целевой объём в каждой таблице - 150_000_000 записей.
В случае, если увеличить пачку до 450-500, при выполнении команды
mysql -h IP -P9306 < /path/instructions_1.sql
Получаем ошибку:
ERROR 2013 (HY000) at line 1: Lost connection to MySQL server during query
И в логах:
WARNING: failed to receive MySQL request body (client=IP:59836(2), exp=16777215, error='Success')

Путём опытов и изучения документации поняли, что упираемся в параметр "max_packet_size" конфига. Выставили его в максимум - 128МБ.
Можно ли как-нибудь значительно (напр. до 1ГБ) увеличить данный параметр?
Подскажите, на какие параметры ещё стоит обратить внимание?
Вот наш конфиг:
searchd {
    ...
    read_buffer_docs = 512m
    read_buffer_hits = 512m
    read_unhinted = 512m
    max_packet_size = 128m
    max_open_files = 100000
    max_batch_queries = 102400
    docstore_cache_size = 1024m
    threads = 50
}
Спасибо!


*/
