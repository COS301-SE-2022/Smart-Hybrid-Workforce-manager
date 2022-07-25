CREATE OR REPLACE FUNCTION resource.identifier_store(
    _id uuid, -- NULLABLE, If supplied try update else insert
    _room_id uuid,
    _name VARCHAR(256),
    _xcoord float,
    _ycoord float,
    _width float,
    _height float,
    _rotation float,
	_role_id uuid,
	_resource_type resource.type,
    _decorations JSON
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
            xcoord = _xcoord,
            ycoord = _ycoord,
            width = _width,
            height = _height,
            rotation = _rotation,        
            role_id = _role_id,
            resource_type = _resource_type
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
        INSERT INTO resource.identifier(id, room_id, name, xcoord, ycoord, width, height, rotation, role_id, resource_type, decorations)
        VALUES (COALESCE(_id, uuid_generate_v4()), _room_id, _name, _xcoord, _ycoord, _width, _height, _rotation, _role_id, _resource_type, _decorations)
        RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
