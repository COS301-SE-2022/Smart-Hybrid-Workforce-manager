CREATE OR REPLACE FUNCTION resource.room_find(
    _id uuid DEFAULT NULL,
	_building_id uuid DEFAULT NULL,
    _name VARCHAR(256) DEFAULT NULL,
    _location VARCHAR(256) DEFAULT NULL,
    _dimension VARCHAR(256) DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	building_id uuid,
    name VARCHAR(256),
	location VARCHAR(256),
	dimension VARCHAR(256)
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.id, i.building_id, i.name, i.location, i.dimension
    FROM resource.room as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_building_id IS NULL OR i.building_id = _building_id)
    AND (_name IS NULL OR i.name = _name)
    AND (_location IS NULL OR i.location = _location)
    AND (_dimension IS NULL OR i.dimension = _dimension);
END
$$ LANGUAGE plpgsql;
