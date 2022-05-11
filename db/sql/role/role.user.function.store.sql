CREATE OR REPLACE FUNCTION role.user_store(
	_role_id uuid,
	_user_id uuid
)
RETURNS BOOLEAN AS
$$
BEGIN
    INSERT INTO role.user(role_id, user_id)
    VALUES (_role_id, _user_id);
	RETURN true;
END
$$ LANGUAGE plpgsql;
