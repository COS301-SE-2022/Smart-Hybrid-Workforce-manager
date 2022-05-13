CREATE OR REPLACE FUNCTION role.identifier_remove(
    _id uuid
)
RETURNS BOOLEAN AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM role.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_role'
            USING HINT = 'Please check the provided role id parameter';
    END IF;

    DELETE FROM role.identifier WHERE id = _id;

    RETURN TRUE;

END
$$ LANGUAGE plpgsql;