CREATE OR REPLACE FUNCTION resource.building_find(
    _id uuid DEFAULT NULL,
	_name VARCHAR(256) DEFAULT NULL,
    _location VARCHAR(256) DEFAULT NULL,
    _dimension VARCHAR(256) DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	name VARCHAR(256),
	location VARCHAR(256),
	dimension VARCHAR(256)
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.id, i.name, i.location, i.dimension
    FROM resource.building as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_name IS NULL OR i.name = _name)
    AND (_location IS NULL OR i.location = _location)
    AND (_dimension IS NULL OR i.dimension = _dimension);
END
$$ LANGUAGE plpgsql;
