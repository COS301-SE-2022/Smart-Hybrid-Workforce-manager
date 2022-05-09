CREATE OR REPLACE FUNCTION booking.identifier_remove(
    _id uuid
)
RETURNS BOOLEAN AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM booking.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_booking'
            USING HINT = 'Please check the provided booking id parameter';
    END IF;

    DELETE FROM booking.identifier WHERE id = _id;

    RETURN TRUE;

END
$$ LANGUAGE plpgsql;