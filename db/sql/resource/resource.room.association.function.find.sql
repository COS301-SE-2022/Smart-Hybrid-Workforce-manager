CREATE OR REPLACE FUNCTION resource.room_association_find(
    _room_id uuid DEFAULT NULL,
	_room_id_association uuid DEFAULT NULL
)
RETURNS TABLE (
    room_id uuid,
	room_id_association uuid
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.room_id, i.room_id_association
    FROM resource.room_association as i
    WHERE (_room_id IS NULL OR i.room_id = _room_id)
    AND (_room_id_association IS NULL OR i.room_id_association = _room_id_association);
END
$$ LANGUAGE plpgsql;
