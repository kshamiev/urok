DO $$
DECLARE
	idstart int DEFAULT 1;
	ids RECORD;
BEGIN
	FOR ids IN SELECT d.id FROM documents d WHERE d.id >= idstart LIMIT 100 LOOP
    RAISE NOTICE 'Done refreshing materialized views.';
   
   
	END LOOP;	
	PERFORM (SELECT now());
	RAISE NOTICE '!!!!!';
END $$;

	

SELECT now();
