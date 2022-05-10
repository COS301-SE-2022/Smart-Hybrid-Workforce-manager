CREATE OR REPLACE FUNCTION team.identifier_find(
    _id uuid DEFAULT NULL,
	_name VARCHAR(256) DEFAULT NULL,
    _description VARCHAR(256) DEFAULT NULL,
    _capacity INT DEFAULT NULL,
    _picture VARCHAR(256) DEFAULT NULL,
    _date_created TIMESTAMP DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	name VARCHAR(256),
	description VARCHAR(256),
	capacity INT,
	picture VARCHAR(256),
    date_created TIMESTAMP
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.id, i.name, i.description, i.capacity, i.picture, i.date_created
    FROM team.identifier as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_name IS NULL OR i.name = _name)
    AND (_description IS NULL OR i.description = _description)
    AND (_capacity IS NULL OR i.capacity = _capacity)
    AND (_picture IS NULL OR i.picture >= _picture)
    AND (_date_created IS NULL OR i.date_created >= _date_created);
END
$$ LANGUAGE plpgsql;
