analyze t;

explain (costs off) select * from t where a = 1;

explain (costs off) select * from t where a <= 100;

explain (costs off) select * from t where a <= 40000;

vacuum t;
explain (costs off) select a from t where a < 100;

explain (analyze, costs off) select a from t where a < 100;
explain (costs off) select a from t where a < 100;

explain (costs off) select * from t where a <= 100 and b = 'a';

explain (costs off) select * from t where b = 'a';

explain (costs off) select * from t where lower(b) = 'a';
explain (costs off) select * from t where b = 'a';

 create index on t(lower(b));

SELECT * FROM pg_stats AS tt WHERE tt.tablename = 't_lower_idx';
SELECT * FROM pg_catalog.pg_stat_all_indexes AS tt WHERE tt.relname = 't';
SELECT * FROM pg_catalog.pg_stat_all_tables AS dd WHERE dd.relname = 't';
SELECT * FROM pg_catalog.pg_stat_activity psa ;


SELECT version();



SELECT '12.34'::float8::numeric::money;


SELECT 'as   '::character(20) = 'as            '::CHARACTER(20); 


SELECT '\xDE AD BE EF';


SELECT 'yes'::boolean;

SELECT NULL::boolean;


SELECT ROW(1,2.5,'this is a test')::TEXT ;
SELECT ROW(1,2.5,'this is a test');

SELECT;
SELECT version();


