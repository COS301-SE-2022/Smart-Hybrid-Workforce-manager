CREATE OR REPLACE FUNCTION team.association_find(
    _team_id uuid DEFAULT NULL,
	_team_id_association uuid DEFAULT NULL
)
RETURNS TABLE (
    team_id uuid,
	team_id_association uuid
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.team_id, i.team_id_association
    FROM team.association as i
    WHERE (_team_id IS NULL OR i.team_id = _team_id)
    AND (_team_id_association IS NULL OR i.team_id_association = _team_id_association);
END
$$ LANGUAGE plpgsql;
