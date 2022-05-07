CREATE OR REPLACE FUNCTION resource.identifier_create(
	_location VARCHAR(256),
	_role_id uuid,
	_resource_type resource.type
)
RETURNS BOOLEAN AS 
$$
BEGIN
    INSERT INTO resource.identifier(location, role_id, resource_type)
    VALUES (_location, _role_id, _resource_type);
    RETURN TRUE;
END
$$ LANGUAGE plpgsql;
