CREATE OR REPLACE FUNCTION team.user_remove(
    _team_id uuid,
    _user_id uuid
)
RETURNS TABLE (
    team_id uuid,
	user_id uuid,
	date_added TIMESTAMP
) AS 
$$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM team.user as b
        WHERE b.team_id = _team_id
        AND b.user_id = _user_id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_team_user'
            USING HINT = 'Please check the provided team and user parameter';
    END IF;

    RETURN QUERY
    DELETE FROM team.user as a
    WHERE a.team_id = _team_id
    AND a.user_id = _user_id
    RETURNING *;

END
$$ LANGUAGE plpgsql;