CREATE OR REPLACE FUNCTION resource.building_find(
    _id uuid DEFAULT NULL,
	_name VARCHAR(256) DEFAULT NULL,
    _location VARCHAR(256) DEFAULT NULL,
    _dimension VARCHAR(256) DEFAULT NULL,
    _permissions JSONB DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
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
    WITH permitted_buildings AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'RESOURCE'::permission.category
        AND permission_tenant = 'BUILDING'::permission.tenant
    )
    SELECT i.id, i.name, i.location, i.dimension
    FROM resource.building as i
    WHERE (EXISTS(SELECT 1 FROM permitted_buildings WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_buildings))
    AND (_id IS NULL OR i.id = _id)
    AND (_name IS NULL OR i.name = _name)
    AND (_location IS NULL OR i.location = _location)
    AND (_dimension IS NULL OR i.dimension = _dimension);
    
    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
