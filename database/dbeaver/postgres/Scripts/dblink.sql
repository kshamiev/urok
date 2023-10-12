CREATE EXTENSION IF NOT EXISTS dblink;
SELECT dblink_connect('migration', 'host=? port=? user=? password=? dbname=?');

-- сущности--

-- пользователи
INSERT INTO public.users SELECT * FROM dblink('migration', $$
SELECT
    id,
    coalesce(user_name, format('LOGIN-%s', id)),
    first_name,
    middle_name,
    last_name,
    coalesce(email, format('EMAIL-%s', id)),
    phone_number,
    avatar_url,
    coalesce(locale, 'ru_RU'),
    coalesce(invited, false)
FROM
    public.user
$$) AS origin (
               id           BIGINT,
               user_name    TEXT,
               first_name   TEXT,
               middle_name  TEXT,
               last_name    TEXT,
               email        TEXT,
               phone_number TEXT,
               avatar_url   TEXT,
               locale       TEXT,
               invited      BOOLEAN
    )
ON CONFLICT DO NOTHING;
SELECT setval('users_id_seq', MAX(id), true)
FROM public.users;
-- свойства пользовательского профиля
INSERT INTO public.user_profiles SELECT * FROM dblink('migration', $$
SELECT
    u.id,
    u.counterparty_id,
    c.uuid,
    u.cart::jsonb,
    CASE WHEN u.decimal_separator=1 THEN 'Comma'
         WHEN u.decimal_separator=2 THEN 'Dot'
         WHEN u.decimal_separator=3 THEN 'Space'
    END AS decimal_separator,
    CASE WHEN u.thousands_separator=1 THEN 'Comma'
         WHEN u.thousands_separator=2 THEN 'Dot'
         WHEN u.thousands_separator=3 THEN 'Space'
    END AS thousands_separator,
    CASE WHEN u.currency_sign_position=1 THEN 'Before'
         WHEN u.currency_sign_position=2 THEN 'After'
    END AS currency_sign,
    u.geo_map_type::text
FROM public.user u
INNER JOIN public.counterparty c
        ON c.id = u.counterparty_id
$$) AS origin (
               user_id             BIGINT,
               counterparty_id     BIGINT,
               counterparty_guid   UUID,
               cart                JSONB,
               decimal_separator   TEXT,
               thousands_separator TEXT,
               currency_sign       TEXT,
               geo_map_type        TEXT
    )
ON CONFLICT DO NOTHING;

-- параметры допуска
INSERT INTO public.permissions SELECT * FROM dblink('migration', $$
SELECT id, name, comment FROM public.permission
$$) AS origin (
               id      BIGINT,
               name    TEXT,
               comment TEXT
    )
ON CONFLICT DO NOTHING;
SELECT setval('permissions_id_seq', MAX(id), true)
FROM public.permissions;
-- роли
INSERT INTO public.roles SELECT * FROM dblink('migration', $$
SELECT id, name, comment FROM public.role
$$) AS origin (
               id      BIGINT,
               name    TEXT,
               comment TEXT
    )
ON CONFLICT DO NOTHING;
SELECT setval('roles_id_seq', MAX(id), true)
FROM public.roles;
-- связи --

-- между пользователями и ролями
INSERT INTO public.users_roles SELECT * FROM dblink('migration', $$
SELECT rua.user_id, rua.role_id, c.id, c.uuid FROM public.role_user_assignment AS rua
   INNER JOIN public.user_partner_counterparty upc ON upc.user_id = rua.user_id AND upc.counterparty_id != 0
   INNER JOIN public.counterparty c ON c.id = upc.counterparty_id
WHERE rua.counterparty_id IS NULL
$$) AS origin (
               user_id         BIGINT,
               role_id         BIGINT,
               counterparty_id BIGINT,
               counterparty_guid UUID
    )
WHERE
        user_id IN (SELECT id FROM public.users) AND
        role_id IN (SELECT id FROM public.roles)
ON CONFLICT DO NOTHING;

INSERT INTO public.users_roles SELECT * FROM dblink('migration', $$
SELECT rua.user_id, rua.role_id, c.id, c.uuid FROM public.role_user_assignment AS rua
   INNER JOIN public.counterparty c ON c.id = rua.counterparty_id
WHERE rua.counterparty_id IS NOT NULL
$$) AS origin (
               user_id         BIGINT,
               role_id         BIGINT,
               counterparty_id BIGINT,
               counterparty_guid UUID
    )
WHERE
        user_id IN (SELECT id FROM public.users) AND
        role_id IN (SELECT id FROM public.roles)
ON CONFLICT DO NOTHING;

-- между ролями и параметрами допуска
INSERT INTO public.roles_permissions SELECT * FROM dblink('migration', $$
SELECT role_id, permission_id FROM public.role_permission
$$) AS origin (
               role_id       BIGINT,
               permission_id BIGINT
    )
WHERE
        role_id       IN (SELECT id FROM public.roles) AND
        permission_id IN (SELECT id FROM public.permissions)
ON CONFLICT DO NOTHING;

-- перманентные токены
INSERT INTO public.tokens (token, user_id, title, expiration_date, allowed_ip, is_revoked, created_at, "view")
SELECT
    encode(sha256(token::bytea), 'hex'),
    user_id,
    title,
    expiration_date,
    allowed_ip,
    is_revoked,
    now(),
    left(token, 4) || repeat('*', length(token) - 4) as view
FROM dblink('migration', $$
SELECT token, user_id, title, expiration_date, allowed_ip, revoke FROM public.token
$$) AS origin (
               token text,
               user_id bigint,
               title text,
               expiration_date timestamp with time zone,
               allowed_ip inet[],
               is_revoked bool
    )
WHERE
        user_id IN (SELECT id FROM public.users)
ON CONFLICT DO NOTHING;

SELECT dblink_disconnect('migration');
