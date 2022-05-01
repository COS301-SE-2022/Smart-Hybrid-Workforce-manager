CREATE OR REPLACE FUNCTION "user".identifier_store(
    _identifier VARCHAR(256),
	_first_name VARCHAR(256),
	_last_name VARCHAR(256),
	_email VARCHAR(256),
	_picture VARCHAR(256)
)
RETURNS BOOLEAN AS 
$$
BEGIN
    IF EXISTS(SELECT 1 FROM "user".identifier WHERE identifier = _identifier) THEN
        UPDATE "user".identifier
        SET first_name = _first_name,
            last_name = _last_name,
            email = _email,
            picture = _picture
        WHERE identifier = _identifier;
    ELSE
        INSERT INTO "user".identifier (identifier, first_name, last_name, email, picture)
        VALUES (_identifier, _first_name, _last_name, _email, _picture);
    END IF;
    RETURN TRUE;
END
$$ LANGUAGE plpgsql;
