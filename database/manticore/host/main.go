package main

import (
	"fmt"

	"github.com/manticoresoftware/go-sdk/manticore"
)

func main() {
	cl := manticore.NewClient()
	cl.SetServer("localhost", 9312)
	if _, err := cl.Open(); err != nil {
		fmt.Println(err)
		return
	}

	// data, err := cl.Json("/search", req)
	// fmt.Println(data, err)

	// res, err := cl.Sphinxql("RELOAD INDEXES")
	// fmt.Println(res, err)

	// q := manticore.NewSearch("Дом на берегу озера", "users", "")
	// q := manticore.NewSearch("дом на холме", "users", "")
	q := manticore.NewSearch("Мухомор", "users", "")
	q.Offset = 0
	q.Limit = 5
	res2, err2 := cl.RunQuery(q)
	fmt.Println(res2, err2)
	// fmt.Println(len(res2.Matches), res2.Total)

	// Total: 3
	// Total found: 2162
	// 'дом' (Docs:2162, Hits:3891)

}

var req = `
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

/*
start
1000000 0
update
999998 2
restart


killlist = SELECT id FROM documents WHERE updated_at >=  (SELECT created_at FROM deltabreaker WHERE index_name='delta')


sql_query_killlist = \
        SELECT id FROM documents WHERE updated_ts>=@last_reindex UNION \
        SELECT id FROM documents_deleted WHERE deleted_ts>=@last_reindex
}

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
