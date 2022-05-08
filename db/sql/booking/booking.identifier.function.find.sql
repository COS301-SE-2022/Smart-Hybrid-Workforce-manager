CREATE OR REPLACE FUNCTION booking.identifier_find(
	_user_id uuid,
    _resource_type resource.type DEFAULT NULL,
    _start TIMESTAMP DEFAULT NULL,
    _end TIMESTAMP DEFAULT NULL,
    _booked BOOLEAN DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	user_id uuid,
	resource_type resource.type,
	resource_preference_id uuid,
	start TIMESTAMP,
	"end" TIMESTAMP,
    booked BOOLEAN,
    date_created TIMESTAMP
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.id, i.user_id, i.resource_type, i.resource_preference_id, i.start, i."end", i.booked, i.date_created
    FROM booking.identifier as i
    WHERE i.user_id = _user_id
    AND (_resource_type IS NULL OR i.resource_type = _resource_type)
    AND (_start IS NULL OR i.start >= _start)
    AND (_end IS NULL OR i."end" >= _end)
    AND (_booked IS NULL OR i.booked = _booked);
END
$$ LANGUAGE plpgsql;
