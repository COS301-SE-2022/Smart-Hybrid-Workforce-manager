CREATE OR REPLACE FUNCTION permission.identifier_remove(
    _permission_id uuid,
    _permission_id_type permission.id_type,
    _permission_type permission.type,
    _permission_category permission.category,
    _permission_tenant permission.tenant,
    _permission_tenant_id uuid
)
RETURNS TABLE (
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

    IF NOT EXISTS (
        SELECT 1
        FROM permission.identifier as b
        WHERE b.permission_id = _permission_id
        AND b.permission_id_type = _permission_id_type
        AND b.permission_type = _permission_type
        AND b.permission_category = _permission_category
        AND b.permission_tenant = _permission_tenant
        AND b.permission_tenant_id = _permission_tenant_id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_identifier'
            USING HINT = 'Please check the provided identifier parameters';
    END IF;

    RETURN QUERY
    DELETE FROM permission.identifier as a
    WHERE a.permission_id = _permission_id
    AND a.permission_id_type = _permission_id_type
    AND a.permission_type = _permission_type
    AND a.permission_category = _permission_category
    AND a.permission_tenant = _permission_tenant
    AND a.permission_tenant_id = _permission_tenant_id
    RETURNING a.permission_id, a.permission_id_type, a.permission_type, a.permission_category, a.permission_tenant, a.permission_tenant_id, a.date_added;

END
$$ LANGUAGE plpgsql;