CREATE OR REPLACE FUNCTION permission.identifier_find(
    _id uuid DEFAULT NULL,
    _permission_id uuid DEFAULT NULL,
    _permission_id_type permission.id_type DEFAULT NULL,
    _permission_type permission.type DEFAULT NULL,
    _permission_category permission.category DEFAULT NULL,
    _permission_tenant permission.tenant DEFAULT NULL,
    _permission_tenant_id uuid DEFAULT NULL,
    _date_added TIMESTAMP DEFAULT NULL,
    _permissions JSONB DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
    permission_id uuid,
    permission_id_type permission.id_type,
	permission_type permission.type,
	permission_category permission.category,
	permission_tenant permission.tenant,
	permission_tenant_id uuid,
	date_added TIMESTAMP
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
        SELECT a.permission_tenant_id FROM _permissions_table as a
        WHERE a.permission_type = 'VIEW'::permission.type
        AND a.permission_category = 'PERMISSION'::permission.category
    )
    SELECT i.id, i.permission_id, i.permission_id_type, i.permission_type, i.permission_category, i.permission_tenant, i.permission_tenant_id, i.date_added
    FROM permission.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_identifiers as b WHERE b.permission_tenant_id is null) OR i.permission_id = ANY(SELECT * FROM permitted_identifiers))
    AND (_id IS NULL OR i.id = _id)
    AND (_permission_id IS NULL OR i.permission_id = _permission_id)
    AND (_permission_id_type IS NULL OR i.permission_id_type = _permission_id_type)
    AND (_permission_type IS NULL OR i.permission_type = _permission_type)
    AND (_permission_category IS NULL OR i.permission_category = _permission_category)
    AND (_permission_tenant IS NULL OR i.permission_tenant = _permission_tenant)
    AND (_permission_tenant_id IS NULL OR i.permission_tenant_id = _permission_tenant_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
