CREATE OR REPLACE FUNCTION "user".identifier_store(
	_first_name VARCHAR(256),
	_last_name VARCHAR(256),
	_email VARCHAR(256),
	_picture VARCHAR(256)
)
RETURNS BOOLEAN AS $$
BEGIN
    IF EXISTS(SELECT 1 FROM "user".identifier WHERE email = _email) THEN
        UPDATE "user".identifier
        SET first_name = _first_name,
            last_name = _last_name,
            picture = _picture
        WHERE email = _email;
    ELSE
        INSERT INTO "user".identifier (first_name, last_name, email, picture)
        VALUES (_first_name, _last_name, _email, _picture);
    END IF;
    RETURN TRUE;
END
$$ LANGUAGE plpgsql;
