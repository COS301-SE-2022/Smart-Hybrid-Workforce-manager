CREATE OR REPLACE FUNCTION "user".credential_remove(
    _id uuid
)
RETURNS TABLE (
    id VARCHAR(256),
	secret VARCHAR(256),
	identifier VARCHAR(256),
	"type"  "user".credential_type,
	active BOOLEAN,
	failed_attempts INT,
    last_accessed TIMESTAMP
) AS 
$$
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

    RETURN QUERY
    DELETE FROM "user".credential as a WHERE a.id = _id
    RETURNING *;

END
$$ LANGUAGE plpgsql;