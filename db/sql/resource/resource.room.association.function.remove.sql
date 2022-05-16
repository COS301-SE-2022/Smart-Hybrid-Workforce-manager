CREATE OR REPLACE FUNCTION resource.room_association_remove(
    _room_id uuid,
    _room_id_association uuid
)
RETURNS TABLE (
    room_id uuid,
	room_id_association uuid
) AS 
$$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM resource.room_association as b
        WHERE b.room_id = _room_id
        AND b.room_id_association = _room_id_association
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_room_association'
            USING HINT = 'Please check the provided room association parameters';
    END IF;

    RETURN QUERY
    DELETE FROM resource.room_association as a 
    WHERE a.room_id = _room_id
    AND a.room_id_association = _room_id_association
    RETURNING *;

END
$$ LANGUAGE plpgsql;