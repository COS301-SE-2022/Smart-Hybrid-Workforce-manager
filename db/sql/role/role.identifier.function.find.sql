CREATE OR REPLACE FUNCTION role.identifier_find(
    _id uuid DEFAULT NULL,
	_role_name VarChar(256) DEFAULT NULL,
    _date_added TIMESTAMP DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	role_name VarChar(256),
	date_added TIMESTAMP
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.id, i.role_name, i.date_added
    FROM role.identifier as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_role_name IS NULL OR i.role_name = _role_name)
    AND (_date_added IS NULL OR i.date_added >= _date_added);
END
$$ LANGUAGE plpgsql;
