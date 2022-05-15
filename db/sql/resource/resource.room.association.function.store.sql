CREATE OR REPLACE FUNCTION resource.room_association_store(
    _room_id uuid,
    _room_id_association uuid
)
RETURNS BOOLEAN AS
$$
BEGIN
    IF (_room_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.room_association WHERE room_id = _room_id)) THEN
        UPDATE resource.room_association
        SET room_id_association = _room_id_association
        WHERE room_id = _room_id;
    ELSE
    	INSERT INTO resource.room_association(room_id, room_id_association)
    	VALUES (_room_id, _room_id_association);
    END IF;
	RETURN true;
END
$$ LANGUAGE plpgsql;
