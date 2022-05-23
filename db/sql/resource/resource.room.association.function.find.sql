CREATE OR REPLACE FUNCTION resource.room_association_find(
    _room_id uuid DEFAULT NULL,
	_room_id_association uuid DEFAULT NULL,
    _permissions JSONB DEFAULT NULL
)
RETURNS TABLE (
    room_id uuid,
	room_id_association uuid
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
    WITH permitted_associations AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'RESOURCE'::permission.category
        AND permission_tenant = 'ROOMASSOCIATION'::permission.tenant
    )
    SELECT i.room_id, i.room_id_association
    FROM resource.room_association as i
    WHERE (EXISTS(SELECT 1 FROM permitted_associations WHERE permission_tenant_id is null) OR i.room_id = ANY(SELECT * FROM permitted_associations))
    AND (_room_id IS NULL OR i.room_id = _room_id)
    AND (_room_id_association IS NULL OR i.room_id_association = _room_id_association);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
