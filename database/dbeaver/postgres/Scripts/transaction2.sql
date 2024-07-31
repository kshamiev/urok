begin;
select * from toys for update;
update toys set usage = usage + 1 where id = 1;
commit;

-- advisory lock (рекомендованные блокировки)

SELECT pg_try_advisory_lock(3034);

SELECT * FROM pg_locks WHERE pid = pg_backend_pid() AND locktype = 'advisory';

SELECT  pg_terminate_backend(3178);
