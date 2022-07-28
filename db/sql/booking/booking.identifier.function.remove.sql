CREATE OR REPLACE FUNCTION booking.identifier_remove(
    _id uuid
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
    automated BOOLEAN,
    date_created TIMESTAMP,
    dependent uuid
) AS 
$$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM booking.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_booking'
            USING HINT = 'Please check the provided booking id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM booking.identifier as a WHERE a.id = _id 
    RETURNING *;

END
$$ LANGUAGE plpgsql;