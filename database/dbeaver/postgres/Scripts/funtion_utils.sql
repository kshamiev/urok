CREATE OR REPLACE FUNCTION public.random_string(num integer, chars text DEFAULT '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'::text)
 RETURNS text
 LANGUAGE plpgsql
AS $function$
DECLARE
  res_str TEXT := '';
BEGIN
  IF num < 1 THEN
      RAISE EXCEPTION 'Invalid length';
  END IF;
  FOR __ IN 1..num LOOP
    res_str := res_str || substr(chars, floor(random() * length(chars))::int + 1, 1);
  END LOOP;
  RETURN res_str;
END $function$
;
