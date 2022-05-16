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
    IF (_user_id IS NOT NULL AND EXISTS(SELECT 1 FROM permission.user WHERE user_id = _user_id)) THEN
        UPDATE permission.user
        SET permission_type = _permission_type,
            permission_category = _permission_category,
            permission_tenant = _permission_tenant,
            permission_tenant_id = _permission_tenant_id
        WHERE user_id = _user_id;
    ELSE
    	INSERT INTO permission.user(user_id, permission_type, permission_category, permission_tenant, permission_tenant_id)
        VALUES (_user_id, _permission_type, _permission_category, _permission_tenant, _permission_tenant_id);
    END IF;
	RETURN true;
END
$$ LANGUAGE plpgsql;
