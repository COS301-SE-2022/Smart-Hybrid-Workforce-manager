CREATE OR REPLACE FUNCTION resource.building_remove(
    _id uuid
)
RETURNS TABLE (
    id uuid,
	name VARCHAR(256),
	location VARCHAR(256),
	dimension VARCHAR(256)
) AS 
$$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM resource.building as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_building'
            USING HINT = 'Please check the provided building id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM resource.building as a WHERE a.id = _id
    RETURNING *;

END
$$ LANGUAGE plpgsql;