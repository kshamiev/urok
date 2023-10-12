DO $$
    DECLARE
        userid record;
        roleid bigint;
    BEGIN
        SELECT id INTO roleid From roles where name = 'Guest' ;
        RAISE NOTICE 'Guest role_id %', roleid;
        FOR userid IN SELECT id From users
            LOOP
                IF  (SELECT count(*) FROM users_roles as ur INNER JOIN roles as r on ur.role_id = r.id WHERE r.name = 'Guest' AND ur.user_id = userid.id) = 0 THEN
                    INSERT INTO users_roles (user_id, role_id)
                    VALUES				   (userid.id, roleid);
                    RAISE NOTICE 'Add GuestRole to UserId %', userid.id;
                END IF;
            END LOOP;
        RAISE NOTICE 'Done';
    END;
$$ language plpgsql;
