-- https://zelark.github.io/exploring-query-locks-in-postgres/
-- https://www.postgresql.org/docs/9.1/explicit-locking.html

create table toys (
  id serial not null,
  name character varying(36),
  usage integer not null default 0,
  constraint toys_pkey primary key (id)
);

insert into toys(name) values('car'),('digger'),('shovel');


-- Показывает текущие транзакции и блокировки

select
  lock.locktype,
  lock.relation::regclass,
  lock.mode,
  lock.transactionid as tid,
  lock.virtualtransaction as vtid,
  lock.pid,
  lock.granted
from pg_catalog.pg_locks lock
  left join pg_catalog.pg_database db
    on db.oid = lock.database
where (db.datname = 'tasksync' or db.datname is null)
  and not lock.pid = pg_backend_pid()
order by lock.pid;


-- Показывает запросы выполняющиеся в данный момент (полезно для транзакций)

select
	now() - stat.query_start as waiting_duration,
	stat.query,
	stat.state,
	stat.wait_event,
	stat.wait_event_type,
	stat.pid
from
	pg_catalog.pg_stat_activity as stat
where
	stat.datname = 'tasksync'
	and not (stat.state = 'idle'
		or stat.pid = pg_backend_pid());

-- Показывает кто кого заблокировал

select
	coalesce(bgl.relation::regclass::text,
	bgl.locktype) as locked_item,
	now() - bda.query_start as waiting_duration,
	bda.pid as blocked_pid,
	bda.query as blocked_query,
	bdl.mode as blocked_mode,
	bga.pid as blocking_pid,
	bga.query as blocking_query,
	bgl.mode as blocking_mode
from
	pg_catalog.pg_locks bdl
join pg_stat_activity bda
    on
	bda.pid = bdl.pid
join pg_catalog.pg_locks bgl
    on
	bgl.pid != bdl.pid
	and (bgl.transactionid = bdl.transactionid
		or bgl.relation = bdl.relation
		and bgl.locktype = bdl.locktype)
join pg_stat_activity bga
    on
	bga.pid = bgl.pid
	and bga.datid = bda.datid
where
	not bdl.granted
	and bga.datname = current_database();


-- advisory lock (рекомендованные блокировки)


SELECT * FROM pg_locks WHERE locktype = 'advisory';


