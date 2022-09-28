CREATE OR REPLACE FUNCTION team.association_find(
    _team_id uuid DEFAULT NULL,
	_team_id_association uuid DEFAULT NULL,
    _permissions JSONB DEFAULT NULL -- Must only contain team_id's
)
RETURNS TABLE (
    team_id uuid,
	team_id_association uuid
) AS 
$$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_teams AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'TEAM'::permission.category
        AND permission_tenant = 'ASSOCIATION'::permission.tenant
    )
    SELECT i.team_id, i.team_id_association
    FROM team.association as i
    WHERE (EXISTS(SELECT 1 FROM permitted_teams WHERE permission_tenant_id is null) OR i.team_id = ANY(SELECT * FROM permitted_teams))
    AND (EXISTS(SELECT 1 FROM permitted_teams WHERE permission_tenant_id is null) OR i.team_id_association = ANY(SELECT * FROM permitted_teams))
    AND (_team_id IS NULL OR i.team_id = _team_id)
    AND (_team_id_association IS NULL OR i.team_id_association = _team_id_association);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
