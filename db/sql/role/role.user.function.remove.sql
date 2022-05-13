CREATE OR REPLACE FUNCTION role.user_remove(
    _role_id uuid,
    _user_id uuid
)
RETURNS BOOLEAN AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM role.user as b
        WHERE b.role_id = _role_id
        AND b.user_id = _user_id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_user_role'
            USING HINT = 'Please check the provided role and user parameter';
    END IF;

    DELETE FROM role.user 
    WHERE role_id = _role_id
    AND user_id = _user_id;

    RETURN TRUE;

END
$$ LANGUAGE plpgsql;