CREATE OR REPLACE FUNCTION booking.meeting_room_find(
	_id uuid DEFAULT NULL, -- NULLABLE, If supplied try update else insert
	_booking_id uuid DEFAULT NULL,
	_team_id uuid DEFAULT NULL,
	_role_id uuid DEFAULT NULL,
	_additional_attendees INT DEFAULT NULL,
	_desks_attendees BOOLEAN DEFAULT NULL,
	_desks_additional_attendees BOOLEAN DEFAULT NULL,
    _permissions JSONB DEFAULT NULL -- Must only contain user_ids not role_ids
)
RETURNS TABLE (
	id uuid,
	booking_id uuid,
	team_id uuid,
	role_id uuid,
	additional_attendees INT,
	desks_attendees BOOLEAN,
	desks_additional_attendees BOOLEAN
) AS 
$$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_meeting_rooms AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'BOOKING'::permission.category
        AND (permission_tenant = 'USER'::permission.tenant OR permission_tenant = 'ROLE'::permission.tenant OR permission_tenant = 'TEAM'::permission.tenant)
    )
    SELECT i.id, i.booking_id, i.team_id, i.role_id, i.additional_attendees, i.desks_attendees, i.desks_additional_attendees
    FROM booking.meeting_room as i
    WHERE (EXISTS(SELECT 1 FROM permitted_meeting_rooms WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_meeting_rooms))
    AND (_id IS NULL OR i.id = _id)
    AND (_booking_id IS NULL OR i.booking_id = _booking_id)
    AND (_team_id IS NULL OR i.team_id = _team_id)
    AND (_role_id IS NULL OR i.role_id = _role_id)
    AND (_additional_attendees IS NULL OR i.additional_attendees = _additional_attendees)
    AND (_desks_attendees IS NULL OR i.desks_attendees = _desks_attendees)
    AND (_desks_additional_attendees IS NULL OR i.desks_additional_attendees = _desks_additional_attendees);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
