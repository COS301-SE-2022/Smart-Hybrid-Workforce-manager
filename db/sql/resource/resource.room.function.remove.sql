CREATE OR REPLACE FUNCTION resource.room_remove(
    _id uuid
)
RETURNS TABLE (
    id uuid,
	building_id uuid,
    name VARCHAR(256),
    xcoord float,
    ycoord float,
    zcoord float,
	dimension VARCHAR(256)
) AS 
$$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM resource.room as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_resource'
            USING HINT = 'Please check the provided resource id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM resource.room as a WHERE a.id = _id
    RETURNING *;

END
$$ LANGUAGE plpgsql;