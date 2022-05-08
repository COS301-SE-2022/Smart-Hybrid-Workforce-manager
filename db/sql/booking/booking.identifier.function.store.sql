CREATE OR REPLACE FUNCTION booking.identifier_create(
	_user_id uuid,
	_resource_type resource.type,
	_resource_preference_id uuid,
	_start TIMESTAMP WITHOUT TIME ZONE,
	_end TIMESTAMP WITHOUT TIME ZONE,
	_id uuid DEFAULT NULL
)
RETURNS uuid AS 
$$
BEGIN
    INSERT INTO booking.identifier(id, user_id, resource_type, resource_preference_id, start, "end")
    VALUES (COALESCE(_id, uuid_generate_v4()), _user_id, _resource_type, _resource_preference_id, _start, _end);
    RETURN identifier.id;
END
$$ LANGUAGE plpgsql;
