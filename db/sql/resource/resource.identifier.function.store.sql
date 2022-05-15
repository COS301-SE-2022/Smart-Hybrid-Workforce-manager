CREATE OR REPLACE FUNCTION resource.identifier_store(
    _id uuid, -- NULLABLE, If supplied try update else insert
    _room_id uuid,
	_location VARCHAR(256),
	_role_id uuid,
	_resource_type resource.type
)
RETURNS uuid AS 
$$
DECLARE
	__id uuid;
BEGIN
    IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.identifier WHERE id = _id)) THEN
        UPDATE resource.identifier
        SET room_id = _room_id,
            location = _location,
            role_id = _role_id,
            resource_type = _resource_type
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO resource.identifier(id, room_id, location, role_id, resource_type)
        VALUES (COALESCE(_id, uuid_generate_v4()), _room_id, _location, _role_id, _resource_type)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
