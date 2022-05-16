CREATE OR REPLACE FUNCTION booking.identifier_find(
    _id uuid DEFAULT NULL,
	_user_id uuid DEFAULT NULL,
    _resource_type resource.type DEFAULT NULL,
    _resource_preference_id uuid DEFAULT NULL,
    _resource_id uuid DEFAULT NULL,
    _start TIMESTAMP DEFAULT NULL,
    _end TIMESTAMP DEFAULT NULL,
    _booked BOOLEAN DEFAULT NULL,
    _date_created TIMESTAMP DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	user_id uuid,
	resource_type resource.type,
	resource_preference_id uuid,
	resource_id uuid,
	start TIMESTAMP,
	"end" TIMESTAMP,
    booked BOOLEAN,
    date_created TIMESTAMP
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.id, i.user_id, i.resource_type, i.resource_preference_id, i.resource_id, i.start, i."end", i.booked, i.date_created
    FROM booking.identifier as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_resource_type IS NULL OR i.resource_type = _resource_type)
    AND (_resource_preference_id IS NULL OR i.resource_preference_id = _resource_preference_id)
    AND (_resource_id IS NULL OR i.resource_id = _resource_id)
    AND (_start IS NULL OR i.start >= _start)
    AND (_end IS NULL OR i."end" >= _end)
    AND (_booked IS NULL OR i.booked = _booked)
    AND (_date_created IS NULL OR i.date_created >= _date_created);
END
$$ LANGUAGE plpgsql;
