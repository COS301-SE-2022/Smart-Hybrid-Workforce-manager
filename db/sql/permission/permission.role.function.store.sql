CREATE OR REPLACE FUNCTION permission.role_store(
    _role_id uuid, -- NULLABLE, If supplied try update else insert
    _permission_type permission.type,
	_permission_category permission.category,
	_permission_tenant permission.tenant,
	_permission_tenant_id uuid
)
RETURNS BOOLEAN AS 
$$
BEGIN
    IF (_role_id IS NOT NULL AND EXISTS(SELECT 1 FROM permission.role WHERE role_id = _role_id)) THEN
        UPDATE permission.role
        SET permission_type = _permission_type,
            permission_category = _permission_category,
            permission_tenant = _permission_tenant,
            permission_tenant_id = _permission_tenant_id
        WHERE role_id = _role_id;
    ELSE
    	INSERT INTO permission.role(role_id, permission_type, permission_category, permission_tenant, permission_tenant_id)
        VALUES (_role_id, _permission_type, _permission_category, _permission_tenant, _permission_tenant_id);
    END IF;
	RETURN true;
END
$$ LANGUAGE plpgsql;
