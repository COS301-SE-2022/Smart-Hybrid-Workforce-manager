CREATE OR REPLACE FUNCTION team.identifier_find(
    _id uuid DEFAULT NULL,
	_name VARCHAR(256) DEFAULT NULL,
    _description VARCHAR(256) DEFAULT NULL,
    _capacity INT DEFAULT NULL,
    _picture VARCHAR(256) DEFAULT NULL,
    _priority INT DEFAULT NULL,
    _team_lead_id uuid DEFAULT NULL,
    _date_created TIMESTAMP DEFAULT NULL,
    _permissions JSONB DEFAULT NULL -- Must only contain role_ids not user_ids
)
RETURNS TABLE (
    id uuid,
	name VARCHAR(256),
	description VARCHAR(256),
	capacity INT,
	picture VARCHAR(256),
    priority INT,
	team_lead_id uuid,
    date_created TIMESTAMP
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
        AND permission_tenant = 'IDENTIFIER'::permission.tenant
    )
    SELECT i.id, i.name, i.description, i.capacity, i.picture, i.priority, i.team_lead_id, i.date_created
    FROM team.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_teams WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_teams))
    AND (_id IS NULL OR i.id = _id)
    AND (_name IS NULL OR i.name = _name)
    AND (_description IS NULL OR i.description = _description)
    AND (_capacity IS NULL OR i.capacity = _capacity)
    AND (_picture IS NULL OR i.picture = _picture)
    AND (_priority IS NULL OR i.priority = _priority)
    AND (_team_lead_id IS NULL OR i.team_lead_id = _team_lead_id)
    AND (_date_created IS NULL OR i.date_created >= _date_created);

    DROP TABLE _permissions_table;
END
$$ LANGUAGE plpgsql;
