CREATE OR REPLACE FUNCTION permission.identifier_remove(
    _id uuid
)
RETURNS TABLE (
    id uuid,
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
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_permission'
            USING HINT = 'Please check the provided permission id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM permission.identifier as a WHERE a.id = _id
    RETURNING *;

END
$$ LANGUAGE plpgsql;