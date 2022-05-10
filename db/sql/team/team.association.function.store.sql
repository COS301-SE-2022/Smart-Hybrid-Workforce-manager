CREATE OR REPLACE FUNCTION team.association_store(
	_team_id uuid, -- NOT NULLABLE
	_team_id_association uuid
)
RETURNS BOOLEAN AS
$$
DECLARE
	__id uuid;
BEGIN
	IF (_team_id IS NOT NULL AND EXISTS(SELECT 1 FROM team.association WHERE team_id = _team_id)) THEN
        UPDATE team.association
        SET team_id_association = _team_id_association
        WHERE team_id = _team_id;
    ELSE
    	INSERT INTO team.association(team_id, team_id_association)
    	VALUES (_team_id, _team_id_association);
    END IF;
	RETURN true;
END
$$ LANGUAGE plpgsql;
