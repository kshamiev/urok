CREATE TABLE foo (c1 integer, c2 text);
INSERT INTO foo
  SELECT i, md5(random()::text)
  FROM generate_series(1, 1000000) AS i;
 
-- анализ таблицы сбрасывает кеш статистики
ANALYZE foo;
-- статистика запроса ожидаемая планировщиком 
EXPLAIN SELECT * FROM foo;

INSERT INTO foo
  SELECT i, md5(random()::text)
  FROM generate_series(1, 10) AS i;
EXPLAIN SELECT * FROM foo;
 
-- реальная статистика запроса
EXPLAIN (ANALYZE, costs off) SELECT * FROM foo;
EXPLAIN (ANALYZE) SELECT * FROM foo;
EXPLAIN (ANALYZE, BUFFERS) SELECT * FROM foo;
-- ANALYZE анализирем реальное выполнение запроса
-- BUFFERS 

-- ВЫВОД
-- Seq Scan читается вся таблица.
-- Index Scan используется индекс для условий WHERE, читает таблицу при отборе строк.
-- Bitmap Index Scan сначала Index Scan, затем контроль выборки по таблице. Эффективно для большого количества строк.
-- Index Only Scan самый быстрый. Читается только индекс. Покрывающий индекс.
 
-- cost "время" получения первой строки, время получения всех данных в милисекундах.
-- rows приблизительное количество возвращаемых строк при выполнении запроса 
-- width средний размер одной строки в байтах.

-- actual time реальное время в миллисекундах, затраченное для получения первой строки и всех строк соответственно.
-- rows реальное количество полученных строк
-- loops сколько раз пришлось выполнить внутренние операции по запросу.

-- Buffers: shared read - количество блоков, считанное с диска.
-- Buffers: shared hit — количество блоков, считанных из кэша PostgreSQL.

SELECT * FROM foo;
