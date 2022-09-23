CREATE OR REPLACE FUNCTION role.identifier_remove(
    _id uuid
)
RETURNS TABLE (
    id uuid,
	name VARCHAR(256),
    color VARCHAR(256),
    lead_id uuid,
	date_added TIMESTAMP
) AS
$$
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

    RETURN QUERY
    DELETE FROM role.identifier as a WHERE a.id = _id
    RETURNING *;

END
$$ LANGUAGE plpgsql;