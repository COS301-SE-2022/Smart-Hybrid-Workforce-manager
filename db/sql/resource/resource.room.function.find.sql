CREATE OR REPLACE FUNCTION resource.room_find(
    _id uuid DEFAULT NULL,
	_building_id uuid DEFAULT NULL,
    _name VARCHAR(256) DEFAULT NULL,
    _location VARCHAR(256) DEFAULT NULL,
    _dimension VARCHAR(256) DEFAULT NULL,
    _permissions JSONB DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	building_id uuid,
    name VARCHAR(256),
	location VARCHAR(256),
	dimension VARCHAR(256)
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
    WITH permitted_rooms AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'RESOURCE'::permission.category
        AND permission_tenant = 'ROOM'::permission.tenant
    )
    SELECT i.id, i.building_id, i.name, i.location, i.dimension
    FROM resource.room as i
    WHERE (EXISTS(SELECT 1 FROM permitted_rooms WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_rooms))
    AND (_id IS NULL OR i.id = _id)
    AND (_building_id IS NULL OR i.building_id = _building_id)
    AND (_name IS NULL OR i.name = _name)
    AND (_location IS NULL OR i.location = _location)
    AND (_dimension IS NULL OR i.dimension = _dimension);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
