CREATE OR REPLACE FUNCTION booking.identifier_store(
	_id uuid, -- NULLABLE, If supplied try update else insert
	_booking_id uuid,
	_team_id uuid,
	_role_id uuid,
	_additional_attendees INT,
	_desks_attendees BOOLEAN,
	_desks_additional_attendees uuid,
)
RETURNS uuid AS
$$
DECLARE
	__id uuid;
BEGIN
	IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM booking.meeting_room WHERE id = _id)) THEN
        UPDATE booking.meeting_room
        SET id = _id,
            booking_id = _booking_id,
            team_id = _team_id,
            role_id = _role_id,
            additional_attendees = _additional_attendees,
            desks_attendees = _desks_attendees,
            desks_additional_attendees = _desks_additional_attendees
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO booking.meeting_room(id, booking_id, team_id, role_id, additional_attendees, desks_attendees, desks_additional_attendees)
    	VALUES (COALESCE(_id, uuid_generate_v4()), _booking_id, _team_id, _role_id, _additional_attendees, _desks_attendees, _desks_additional_attendees)
		RETURNING meeting_room.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
