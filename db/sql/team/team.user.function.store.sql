CREATE OR REPLACE FUNCTION team.user_store(
	_team_id uuid, -- NOT NULLABLE
	_user_id uuid
)
RETURNS BOOLEAN AS
$$
BEGIN
    INSERT INTO team.user(team_id, user_id)
    VALUES (_team_id, _user_id);
	RETURN true;
END
$$ LANGUAGE plpgsql;
