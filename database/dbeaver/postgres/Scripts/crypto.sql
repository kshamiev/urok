SELECT * FROM tokens t WHERE t."token" = encode(sha256('tm2Bvk8rE6Bk'::bytea), 'hex')

SELECT * FROM tokens t WHERE t."token" = encode(sha256("token"::bytea), 'hex')