-- восстановление autoincrement в postgres
select setval('order_options_id_seq', 10, false); 

