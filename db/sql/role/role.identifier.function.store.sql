CREATE OR REPLACE FUNCTION role.identifier_store(
	_id uuid, -- NULLABLE, If supplied try update else insert
	_role_name VarChar(256)
)
RETURNS uuid AS
$$
DECLARE
	__id uuid;
BEGIN
	IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM role.identifier WHERE id = _id)) THEN
        UPDATE role.identifier
        SET role_name = _role_name
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO role.identifier(id, role_name)
    	VALUES (COALESCE(_id, uuid_generate_v4()), _role_name)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
