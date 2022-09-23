CREATE OR REPLACE FUNCTION role.identifier_store(
	_id uuid, -- NULLABLE, If supplied try update else insert
	_name VARCHAR(256),
    _color VARCHAR(256),
	_lead_id uuid DEFAULT NULL
)
RETURNS uuid AS
$$
DECLARE
	__id uuid;
BEGIN
	IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM role.identifier WHERE id = _id)) THEN
        UPDATE role.identifier
        SET name = _name,
            color = _color,
			lead_id = _lead_id
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO role.identifier(id, name, color, lead_id)
    	VALUES (COALESCE(_id, uuid_generate_v4()), _name, _color, _lead_id)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
