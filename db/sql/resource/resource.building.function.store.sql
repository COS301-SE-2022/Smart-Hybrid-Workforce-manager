CREATE OR REPLACE FUNCTION resource.building_store(
    _id uuid, -- NULLABLE, If supplied try update else insert
    _name VARCHAR(256),
	_location VARCHAR(256),
	_dimension VARCHAR(256)
)
RETURNS uuid AS 
$$
DECLARE
	__id uuid;
BEGIN
    IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.building WHERE id = _id)) THEN
        UPDATE resource.building
        SET name = _name,
            location = _location,
            dimension = _dimension
        WHERE id = _id
		RETURNING building.id INTO __id;
    ELSE
    	INSERT INTO resource.building(id, name, location, dimension)
        VALUES (COALESCE(_id, uuid_generate_v4()), _name, _location, _dimension)
		RETURNING building.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
