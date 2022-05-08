CREATE OR REPLACE FUNCTION booking.identifier_store(
	_id uuid, -- NULLABLE, If supplied try update else insert
	_user_id uuid,
	_resource_type resource.type,
	_resource_preference_id uuid,
	_start TIMESTAMP WITHOUT TIME ZONE,
	_end TIMESTAMP WITHOUT TIME ZONE,
	_booked BOOLEAN DEFAULT NULL -- Defaults to false and is not considered for creation
)
RETURNS uuid AS 
$$
DECLARE
	__id uuid;
BEGIN
	IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM booking.identifier WHERE id = _id)) THEN
        UPDATE booking.identifier
        SET user_id = _user_id,
            resource_type = _resource_type,
            resource_preference_id = _resource_preference_id,
            start = _start,
            "end" = _end,
            booked = _booked
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO booking.identifier(id, user_id, resource_type, resource_preference_id, start, "end")
    	VALUES (COALESCE(_id, uuid_generate_v4()), _user_id, _resource_type, _resource_preference_id, _start, _end)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
