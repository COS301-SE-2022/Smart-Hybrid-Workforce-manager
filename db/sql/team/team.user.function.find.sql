CREATE OR REPLACE FUNCTION team.user_find(
    _team_id uuid DEFAULT NULL,
	_user_id uuid DEFAULT NULL,
    _date_added TIMESTAMP DEFAULT NULL,
    _permissions JSONB DEFAULT NULL -- user_ids and team_ids
)
RETURNS TABLE (
    team_id uuid,
	user_id uuid,
	date_added TIMESTAMP
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
    WITH permitted_users AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'USER'::permission.category
        AND permission_tenant = 'TEAM'::permission.tenant
    ),
    permitted_teams AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'TEAM'::permission.category
        AND permission_tenant = 'USER'::permission.tenant
    )
    SELECT i.team_id, i.user_id, i.date_added
    FROM team.user as i
    WHERE (EXISTS(SELECT 1 FROM permitted_users WHERE permission_tenant_id is null) OR i.user_id = ANY(SELECT * FROM permitted_users))
    AND (EXISTS(SELECT 1 FROM permitted_teams WHERE permission_tenant_id is null) OR i.team_id = ANY(SELECT * FROM permitted_teams))
    AND (_team_id IS NULL OR i.team_id = _team_id)
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
