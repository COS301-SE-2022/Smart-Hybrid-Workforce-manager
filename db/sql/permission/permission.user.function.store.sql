CREATE OR REPLACE FUNCTION permission.user_store(
    _user_id uuid, -- NULLABLE, If supplied try update else insert
    _permission_type permission.type,
	_permission_category permission.category,
	_permission_tenant permission.tenant,
	_permission_tenant_id uuid
)
RETURNS BOOLEAN AS 
$$
BEGIN

    INSERT INTO permission.user(user_id, permission_type, permission_category, permission_tenant, permission_tenant_id)
    VALUES (_user_id, _permission_type, _permission_category, _permission_tenant, _permission_tenant_id);

	RETURN true;
END
$$ LANGUAGE plpgsql;
