CREATE OR REPLACE FUNCTION "user".identifier_find(
    _id uuid DEFAULT NULL,
	_identifier VARCHAR(256) DEFAULT NULL,
    _first_name VARCHAR(256) DEFAULT NULL,
    _last_name VARCHAR(256) DEFAULT NULL,
    _email VARCHAR(256) DEFAULT NULL,
    _picture VARCHAR(256) DEFAULT NULL,
    _date_created TIMESTAMP DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
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
    SELECT i.id, i.identifier, i.first_name, i.last_name, i.email, i.picture, i.date_created
    FROM "user".identifier as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_identifier IS NULL OR i.identifier = _identifier)
    AND (_first_name IS NULL OR i.first_name = _first_name)
    AND (_last_name IS NULL OR i.last_name = _last_name)
    AND (_email IS NULL OR i.email = _email)
    AND (_picture IS NULL OR i.picture = _picture)
    AND (_date_created IS NULL OR i.date_created >= _date_created);
END
$$ LANGUAGE plpgsql;
