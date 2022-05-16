CREATE OR REPLACE FUNCTION permission.role_remove(
    _role_id uuid,
    _permission_type permission.type,
    _permission_category permission.category,
    _permission_tenant permission.tenant,
    _permission_tenant_id uuid
)
RETURNS TABLE (
    role_id uuid,
	permission_type permission.type,
	permission_category permission.category,
	permission_tenant permission.tenant,
	permission_tenant_id uuid,
	date_added TIMESTAMP
) AS 
$$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM permission.role as b
        WHERE b.role_id = _role_id
        AND b.permission_type = _permission_type
        AND b.permission_category = _permission_category
        AND b.permission_tenant = _permission_tenant
        AND b.permission_tenant_id = _permission_tenant_id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_role'
            USING HINT = 'Please check the provided role parameters';
    END IF;

    RETURN QUERY
    DELETE FROM permission.role as a
    WHERE a.role_id = _role_id
    AND a.permission_type = _permission_type
    AND a.permission_category = _permission_category
    AND a.permission_tenant = _permission_tenant
    AND a.permission_tenant_id = _permission_tenant_id
    RETURNING *;

END
$$ LANGUAGE plpgsql;