CREATE OR REPLACE FUNCTION resource.identifier_store(
    _id uuid, -- NULLABLE, If supplied try update else insert
    _room_id uuid,
    _name VARCHAR(256),
	_location VARCHAR(256),
	_role_id uuid,
	_resource_type resource.type,
    _computer BOOLEAN,
    _capacity INTEGER,
    _disabled BOOLEAN
)
RETURNS uuid AS 
$$
DECLARE
	__id uuid;
BEGIN
    IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.identifier WHERE id = _id)) THEN
        UPDATE resource.identifier
        SET room_id = _room_id,
            name = _name,
            location = _location,            
            role_id = _role_id,
            resource_type = _resource_type
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
        IF(_resource_type = 'DESK') THEN
        BEGIN
            INSERT INTO resource.identifier(id, room_id, name, location, role_id, resource_type, decorations)
            VALUES (COALESCE(_id, uuid_generate_v4()), _room_id, _name, _location, _role_id, _resource_type, '{"computer": ' + _computer + '}')
            RETURNING identifier.id INTO __id;
        END

        IF(_resource_type = 'MEETINGROOM') THEN
        BEGIN
            INSERT INTO resource.identifier(id, room_id, name, location, role_id, resource_type, decorations)
            VALUES (COALESCE(_id, uuid_generate_v4()), _room_id, _name, _location, _role_id, _resource_type, '{"capacity": ' + _capacity + '}')
            RETURNING identifier.id INTO __id;
        END

        IF(_resource_type = 'PARKING') THEN
        BEGIN
            INSERT INTO resource.identifier(id, room_id, name, location, role_id, resource_type, decorations)
            VALUES (COALESCE(_id, uuid_generate_v4()), _room_id, _name, _location, _role_id, _resource_type, '{"disabled": ' + _disabled + '}')
            RETURNING identifier.id INTO __id;
        END
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
