CREATE OR REPLACE FUNCTION role.identifier_find(
    _id uuid DEFAULT NULL,
	_role_name VarChar(256) DEFAULT NULL,
    _date_added TIMESTAMP DEFAULT NULL,
    _permissions JSONB DEFAULT NULL -- Must only contain role_ids not user_ids
)
RETURNS TABLE (
    id uuid,
	role_name VarChar(256),
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
    WITH permitted_roles AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'ROLE'::permission.category
        AND permission_tenant = 'IDENTIFIER'::permission.tenant
    )
    SELECT i.id, i.role_name, i.date_added
    FROM role.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_roles WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_roles))
    AND (_id IS NULL OR i.id = _id)
    AND (_role_name IS NULL OR i.role_name = _role_name)
    AND (_date_added IS NULL OR i.date_added >= _date_added);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
