CREATE OR REPLACE FUNCTION "user".credential_remove(
    _id uuid
)
RETURNS BOOLEAN AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM "user".credential as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_user'
            USING HINT = 'Please check the provided user id parameter';
    END IF;

    DELETE FROM "user".credential WHERE id = _id;

    RETURN TRUE;

END
$$ LANGUAGE plpgsql;