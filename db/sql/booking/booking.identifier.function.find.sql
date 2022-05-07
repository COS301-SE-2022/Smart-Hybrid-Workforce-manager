CREATE OR REPLACE FUNCTION booking.identifier_find(
	_user_id uuid
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
    WHERE i.user_id = _user_id;
END
$$ LANGUAGE plpgsql;
