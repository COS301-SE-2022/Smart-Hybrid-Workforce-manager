CREATE OR REPLACE FUNCTION team.identifier_remove(
    _id uuid
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

    IF NOT EXISTS (
        SELECT 1
        FROM team.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_team'
            USING HINT = 'Please check the provided team id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM team.identifier as a WHERE a.id = _id
    RETURNING *;

END
$$ LANGUAGE plpgsql;