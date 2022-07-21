CREATE OR REPLACE FUNCTION permission.identifier_store(
    _permission_id uuid, -- NULLABLE, If supplied try update else insert
    _permission_id_type permission.id_type,
    _permission_type permission.type,
	_permission_category permission.category,
	_permission_tenant permission.tenant,
	_permission_tenant_id uuid
)
RETURNS uuid AS 
$$
DECLARE
	__id uuid;
BEGIN
    INSERT INTO permission.identifier(permission_id, permission_id_type, permission_type, permission_category, permission_tenant, permission_tenant_id)
    VALUES (_permission_id, _permission_id_type, _permission_type, _permission_category, _permission_tenant, _permission_tenant_id)
    RETURNING identifier.id INTO __id;

    RETURN __id;
END
$$ LANGUAGE plpgsql;
