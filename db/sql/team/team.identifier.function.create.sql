CREATE OR REPLACE FUNCTION team.identifier_create(
	_name VARCHAR(256),
	_description VARCHAR(256),
	_capacity INT,
	_picture VARCHAR(256)
)
RETURNS BOOLEAN AS 
$$
BEGIN
    INSERT INTO team.identifier(name, description, capacity, picture)
    VALUES (_name, _description, _capacity, _picture);
    RETURN TRUE;
END
$$ LANGUAGE plpgsql;
