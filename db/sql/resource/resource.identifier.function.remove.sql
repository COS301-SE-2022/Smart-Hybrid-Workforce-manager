CREATE OR REPLACE FUNCTION resource.identifier_remove(
    _id uuid
)
RETURNS TABLE (
    id uuid,
	room_id uuid,
    name VARCHAR(256),
	xcoord float,
    ycoord float,
    width float,
    height float,
    rotation float,
	resource_type resource.type,
    date_created TIMESTAMP,
    decorations JSON
) AS 
$$
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

    RETURN QUERY
    DELETE FROM resource.identifier as a WHERE a.id = _id
    RETURNING *;

END
$$ LANGUAGE plpgsql;