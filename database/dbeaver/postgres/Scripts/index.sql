create table t(a integer, b text, c boolean);

insert into t(a,b,c)
  select s.id, chr((32+random()*94)::integer), random() < 0.01
  from generate_series(1,100000) as s(id)
  order by random();

create index on t using btree (a);
create index on t(b);
create index on t(a,b);
create index on t(lower(b));
create index on t(c);
create index on t(c) where c;
 
analyze t;
vacuum t;
SELECT count(*) FROM t;

-- Index Scan
explain (ANALYZE, costs off) select * from t where a = 1;

-- Bitmap Heap Scan
explain (ANALYZE, costs off) select * from t where a <= 100;

-- Bitmap Heap Scan
explain (ANALYZE, costs off) select * from t where a <= 100 and b = 'a';
explain (ANALYZE, costs off) select * from t where a <= 100;

-- Seq Scan
explain (ANALYZE, costs off) select * from t where a <= 40000;

-- Index Only Scan
explain (costs off) select a from t where a < 100;

-- Heap Fetches количество обращений к таблице данных
explain (analyze, costs off) select a from t where a < 100;

explain (costs off) select * from t where lower(b) = 'a';

explain (costs off) select * from t where c;
explain (costs off) select * from t where not c;

