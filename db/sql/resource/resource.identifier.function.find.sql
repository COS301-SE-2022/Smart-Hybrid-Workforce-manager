CREATE OR REPLACE FUNCTION resource.identifier_find(
    _id uuid DEFAULT NULL,
	_room_id uuid DEFAULT NULL,
    _location VARCHAR(256) DEFAULT NULL,
    _role_id uuid DEFAULT NULL,
    _resource_type resource.type DEFAULT NULL,
    _date_created TIMESTAMP DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	room_id uuid,
	location VARCHAR(256),
	resource_preference_id uuid,
	resource_type TIMESTAMP,
    date_created TIMESTAMP
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.id, i.room_id, i.location, i.resource_preference_id, i.resource_type, i.date_created
    FROM resource.identifier as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_room_id IS NULL OR i.room_id = _room_id)
    AND (_location IS NULL OR i.location = _location)
    AND (_role_id IS NULL OR i.role_id = _role_id)
    AND (_resource_type IS NULL OR i.resource_type = _resource_type)
    AND (_date_created IS NULL OR i.date_created >= _date_created);
END
$$ LANGUAGE plpgsql;