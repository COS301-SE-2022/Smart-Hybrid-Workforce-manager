CREATE OR REPLACE FUNCTION "user".identifier_find (
	_identifier varchar(256)
)
RETURNS TABLE (
    identifier VARCHAR(256),
	first_name VARCHAR(256),
	last_name VARCHAR(256),
	email VARCHAR(256),
	picture VARCHAR(256),
	date_created TIMESTAMP
) AS 
$$
BEGIN
	RETURN QUERY
    SELECT i.identifier, i.first_name, i.last_name, i.email, i.picture, i.date_created
    FROM "user".identifier as i
    WHERE i.identifier = _identifier;
END
$$ LANGUAGE plpgsql;
