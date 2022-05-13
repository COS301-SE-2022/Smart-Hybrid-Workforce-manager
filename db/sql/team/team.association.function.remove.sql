CREATE OR REPLACE FUNCTION team.association_remove(
    _team_id uuid,
    _team_id_association uuid
)
RETURNS BOOLEAN AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM team.association as b
        WHERE b.team_id = _team_id
        AND b.team_id_association = _team_id_association
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_team_association'
            USING HINT = 'Please check the provided team and association parameter';
    END IF;

    DELETE FROM team.association 
    WHERE team_id = _team_id
    AND team_id_association = _team_id_association;

    RETURN TRUE;

END
$$ LANGUAGE plpgsql;