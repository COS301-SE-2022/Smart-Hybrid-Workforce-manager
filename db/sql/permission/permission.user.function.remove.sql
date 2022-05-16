CREATE OR REPLACE FUNCTION permission.user_remove(
    _user_id uuid,
    _permission_type permission.type,
    _permission_category permission.category,
    _permission_tenant permission.tenant,
    _permission_tenant_id uuid
)
RETURNS TABLE (
    user_id uuid,
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
        FROM permission.user as b
        WHERE b.user_id = _user_id
        AND b.permission_type = _permission_type
        AND b.permission_category = _permission_category
        AND b.permission_tenant = _permission_tenant
        AND b.permission_tenant_id = _permission_tenant_id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_user'
            USING HINT = 'Please check the provided user parameters';
    END IF;

    RETURN QUERY
    DELETE FROM permission.user as a
    WHERE a.user_id = _user_id
    AND a.permission_type = _permission_type
    AND a.permission_category = _permission_category
    AND a.permission_tenant = _permission_tenant
    AND a.permission_tenant_id = _permission_tenant_id
    RETURNING *;

END
$$ LANGUAGE plpgsql;