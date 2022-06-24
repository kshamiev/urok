CREATE TRIGGER carriers_delivery_statuses_updated_at
    BEFORE
        UPDATE
    ON
        public.carriers_delivery_statuses
    FOR EACH ROW
EXECUTE PROCEDURE moddatetime('updated_at');
