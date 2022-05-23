CREATE OR REPLACE FUNCTION "user".identifier_store(
    _id uuid, -- NULLABLE, If supplied try update else insert
    _identifier VARCHAR(256),
	_first_name VARCHAR(256),
	_last_name VARCHAR(256),
	_email VARCHAR(256),
	_picture VARCHAR(256)
)
RETURNS uuid AS 
$$
DECLARE
	__id uuid;
BEGIN
    IF EXISTS(SELECT 1 FROM "user".identifier WHERE identifier = _identifier AND id = _id) THEN
        UPDATE "user".identifier
        SET first_name = _first_name,
            last_name = _last_name,
            email = _email,
            picture = _picture
        WHERE identifier = _identifier
        RETURNING identifier.id INTO __id;
    ELSE
        INSERT INTO "user".identifier (id, identifier, first_name, last_name, email, picture)
        VALUES (COALESCE(_id, uuid_generate_v4()), _identifier, _first_name, _last_name, _email, _picture)
        RETURNING identifier.id INTO __id;
    END IF;
    RETURN __id;
END
$$ LANGUAGE plpgsql;
