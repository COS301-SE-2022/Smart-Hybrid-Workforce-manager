CREATE OR REPLACE FUNCTION booking.identifier_create(
	_user_id uuid,
	_resource_type resource.type,
	_resource_preference_id uuid,
	_start TIMESTAMP WITHOUT TIME ZONE,
	_end TIMESTAMP WITHOUT TIME ZONE
)
RETURNS BOOLEAN AS 
$$
BEGIN
    INSERT INTO booking.identifier(user_id, resource_type, resource_preference_id, start, "end")
    VALUES (_user_id, _resource_type, _resource_preference_id, _start, _end);
    RETURN TRUE;
END
$$ LANGUAGE plpgsql;
