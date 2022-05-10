CREATE OR REPLACE FUNCTION team.user_find(
    _team_id uuid DEFAULT NULL,
	_user_id uuid DEFAULT NULL,
    _date_added TIMESTAMP DEFAULT NULL
)
RETURNS TABLE (
    team_id uuid,
	user_id uuid,
	date_added TIMESTAMP
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.team_id, i.user_id, i.date_added
    FROM team.user as i
    WHERE (_team_id IS NULL OR i.team_id = _team_id)
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);
END
$$ LANGUAGE plpgsql;
