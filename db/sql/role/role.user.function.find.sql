CREATE OR REPLACE FUNCTION role.user_find(
    _role_id uuid DEFAULT NULL,
	_user_id uuid DEFAULT NULL,
    _date_added TIMESTAMP DEFAULT NULL
)
RETURNS TABLE (
    role_id uuid,
	user_id uuid,
	date_added TIMESTAMP
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.role_id, i.user_id, i.date_added
    FROM role.user as i
    WHERE (_role_id IS NULL OR i.role_id = _role_id)
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);
END
$$ LANGUAGE plpgsql;
