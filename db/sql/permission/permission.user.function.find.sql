CREATE OR REPLACE FUNCTION permission.user_find(
    _user_id uuid DEFAULT NULL,
	_permission_type permission.type DEFAULT NULL,
    _permission_category permission.category DEFAULT NULL,
    _permission_tenant permission.tenant DEFAULT NULL,
    _permission_tenant_id uuid DEFAULT NULL,
    _date_added TIMESTAMP DEFAULT NULL
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
    RETURN QUERY
    SELECT i.user_id, i.permission_type, i.permission_category, i.permission_tenant, i.permission_tenant_id, i.date_added
    FROM permission.user as i
    WHERE (_user_id IS NULL OR i.user_id = _user_id)
    AND (_permission_type IS NULL OR i.permission_type = _permission_type)
    AND (_permission_category IS NULL OR i.permission_category = _permission_category)
    AND (_permission_tenant IS NULL OR i.permission_tenant = _permission_tenant)
    AND (_permission_tenant_id IS NULL OR i.permission_tenant_id = _permission_tenant_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);
END
$$ LANGUAGE plpgsql;
