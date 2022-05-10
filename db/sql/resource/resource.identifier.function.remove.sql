CREATE OR REPLACE FUNCTION resource.identifier_remove(
    _id uuid
)
RETURNS BOOLEAN AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM resource.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_resource'
            USING HINT = 'Please check the provided resource id parameter';
    END IF;

    DELETE FROM resource.identifier WHERE id = _id;

    RETURN TRUE;

END
$$ LANGUAGE plpgsql;