-- блокировка на таблицу
begin; 
lock table toys in access exclusive mode;
rollback;

-- блокировка на строки

begin; 
select * from toys for update; -- исключительная на запись и на чтение
update toys set usage = usage + 1 where id = 1;
commit;

begin; 
select * from toys for share; -- на запись
update toys set usage = usage + 1 where id = 1;
commit;


-- advisory lock (рекомендованные блокировки)

SELECT pg_try_advisory_lock(3034);

SELECT * FROM pg_locks WHERE pid = pg_backend_pid() AND locktype = 'advisory';

SELECT  pg_terminate_backend(75);
