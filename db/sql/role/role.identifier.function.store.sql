CREATE OR REPLACE FUNCTION role.identifier_store(
	_id uuid, -- NULLABLE, If supplied try update else insert
	_role_name VarChar(256),
	_role_lead_id uuid DEFAULT NULL
)
RETURNS uuid AS
$$
DECLARE
	__id uuid;
BEGIN
	IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM role.identifier WHERE id = _id)) THEN
        UPDATE role.identifier
        SET role_name = _role_name,
			role_lead_id = _role_lead_id
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO role.identifier(id, role_name, role_lead_id)
    	VALUES (COALESCE(_id, uuid_generate_v4()), _role_name, _role_lead_id)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$ LANGUAGE plpgsql;
