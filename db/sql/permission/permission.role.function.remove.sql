CREATE OR REPLACE FUNCTION permission.role_remove(
    _role_id uuid,
    _permission_type permission.type,
    _permission_category permission.category,
    _permission_tenant permission.tenant,
    _permission_tenant_id uuid
)
RETURNS BOOLEAN AS $$
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

    DELETE FROM permission.role 
    WHERE role_id = _role_id
    AND permission_type = _permission_type
    AND permission_category = _permission_category
    AND permission_tenant = _permission_tenant
    AND permission_tenant_id = _permission_tenant_id;

    RETURN TRUE;

END
$$ LANGUAGE plpgsql;