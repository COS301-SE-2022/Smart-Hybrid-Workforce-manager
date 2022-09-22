CREATE OR REPLACE FUNCTION "user".identifier_remove(
    _id uuid
)
RETURNS TABLE (
    id uuid,
	identifier VARCHAR(256),
	first_name VARCHAR(256),
	last_name VARCHAR(256),
	email VARCHAR(256),
	picture VARCHAR(256),
    date_created TIMESTAMP,
    work_from_home BOOLEAN,
    parking parking.type,
    office_days INTEGER,
    preferred_start_time TIME WITHOUT TIME ZONE,
    preferred_end_time TIME WITHOUT TIME ZONE,
    building_id uuid
) AS 
$$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM "user".identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_user'
            USING HINT = 'Please check the provided user id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM "user".identifier as a WHERE a.id = _id
    RETURNING *;

END
$$ LANGUAGE plpgsql;