CREATE OR REPLACE FUNCTION resource.room_store(
    _id uuid, -- NULLABLE, If supplied try update else insert
    _building_id uuid,
    _name VARCHAR(256),
	_location VARCHAR(256),
	_dimension VARCHAR(256)
)
RETURNS uuid AS 
$$
DECLARE
	__id uuid;
BEGIN
    IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.room WHERE id = _id)) THEN
        UPDATE resource.room
        SET name = _name,
            building_id = _building_id,
            location = _location,
            dimension = _dimension
        WHERE id = _id
		RETURNING room.id INTO __id;
    ELSE
    	INSERT INTO resource.room(id, building_id, name, location, dimension)
        VALUES (COALESCE(_id, uuid_generate_v4()), _building_id, _name, _location, _dimension)
		RETURNING room.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
