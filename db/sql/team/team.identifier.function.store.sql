CREATE OR REPLACE FUNCTION team.identifier_store(
	_id uuid, -- NULLABLE, If supplied try update else insert
	_name VARCHAR(256),
	_description VARCHAR(256),
	_capacity INT,
	_picture VARCHAR(256)
)
RETURNS uuid AS
$$
DECLARE
	__id uuid;
BEGIN
	IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM team.identifier WHERE id = _id)) THEN
        UPDATE team.identifier
        SET name = _name,
            description = _description,
            capacity = _capacity,
            picture = _picture
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO team.identifier(id, name, description, capacity, picture)
    	VALUES (COALESCE(_id, uuid_generate_v4()), _name, _description, _capacity, _picture)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;