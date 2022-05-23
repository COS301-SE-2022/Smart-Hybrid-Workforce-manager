CREATE OR REPLACE FUNCTION booking.identifier_find(
    _id uuid DEFAULT NULL,
	_user_id uuid DEFAULT NULL,
    _resource_type resource.type DEFAULT NULL,
    _resource_preference_id uuid DEFAULT NULL,
    _resource_id uuid DEFAULT NULL,
    _start TIMESTAMP DEFAULT NULL,
    _end TIMESTAMP DEFAULT NULL,
    _booked BOOLEAN DEFAULT NULL,
    _date_created TIMESTAMP DEFAULT NULL,
    _permissions JSONB DEFAULT NULL -- Must only contain user_ids not role_ids
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
    WITH permitted_users AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'BOOKING'::permission.category
        AND permission_tenant = 'USER'::permission.tenant
    )
    SELECT i.id, i.user_id, i.resource_type, i.resource_preference_id, i.resource_id, i.start, i."end", i.booked, i.date_created
    FROM booking.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_users WHERE permission_tenant_id is null) OR i.user_id = ANY(SELECT * FROM permitted_users))
    AND (_id IS NULL OR i.id = _id)
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_resource_type IS NULL OR i.resource_type = _resource_type)
    AND (_resource_preference_id IS NULL OR i.resource_preference_id = _resource_preference_id)
    AND (_resource_id IS NULL OR i.resource_id = _resource_id)
    AND (_start IS NULL OR i.start >= _start)
    AND (_end IS NULL OR i."end" <= _end)
    AND (_booked IS NULL OR i.booked = _booked)
    AND (_date_created IS NULL OR i.date_created >= _date_created);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
