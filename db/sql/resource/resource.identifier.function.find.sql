CREATE OR REPLACE FUNCTION resource.identifier_find(
    _id uuid DEFAULT NULL,
	_room_id uuid DEFAULT NULL,
    _name VARCHAR(256) DEFAULT NULL,
    _xcoord float DEFAULT NULL,
    _ycoord float DEFAULT NULL,
    _width float DEFAULT NULL,
    _height float DEFAULT NULL,
    _rotation float DEFAULT NULL,
    _resource_type resource.type DEFAULT NULL,
    _date_created TIMESTAMP DEFAULT NULL,
    _permissions JSONB DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	room_id uuid,
    name VARCHAR(256),
	xcoord float,
    ycoord float,
    width float,
    height float,
    rotation float,
	resource_type resource.type,
    date_created TIMESTAMP,
    decorations JSON
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
    WITH permitted_identifiers AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'RESOURCE'::permission.category
        AND permission_tenant = 'IDENTIFIER'::permission.tenant
    )
    SELECT i.id, i.room_id, i.name, i.xcoord, i.ycoord, i.width, i.height, i.rotation, i.resource_type, i.date_created, i.decorations
    FROM resource.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_identifiers WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_identifiers))
    AND (_id IS NULL OR i.id = _id)
    AND (_room_id IS NULL OR i.room_id = _room_id)
    AND (_name IS NULL OR i.name = _name)
    AND (_xcoord IS NULL OR i.xcoord = _xcoord)
    AND (_ycoord IS NULL OR i.ycoord = _ycoord)
    AND (_width IS NULL OR i.width = _width)
    AND (_height IS NULL OR i.height = _height)
    AND (_rotation IS NULL OR i.rotation = _rotation)
    AND (_resource_type IS NULL OR i.resource_type = _resource_type)
    AND (_date_created IS NULL OR i.date_created >= _date_created);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;