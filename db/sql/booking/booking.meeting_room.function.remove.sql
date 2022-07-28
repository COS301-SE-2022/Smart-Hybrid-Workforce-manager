CREATE OR REPLACE FUNCTION booking.meeting_room_remove(
    _id uuid
)
RETURNS TABLE (
    id uuid,
    booking_id uuid,
    team_id uuid,
    role_id uuid,
    additional_attendess INT,
    desks_attendees BOOLEAN,
    desks_additional_attendees BOOLEAN
) AS 
$$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM booking.meeting_room as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_meeting_room_boooking'
            USING HINT = 'Please check the provided meeting_room_booking id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM booking.meeting_room as a WHERE a.id = _id 
    RETURNING *;

END
$$ LANGUAGE plpgsql;