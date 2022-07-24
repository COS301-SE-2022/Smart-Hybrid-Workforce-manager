CREATE OR REPLACE FUNCTION "user".credential_store (
	_id VARCHAR(256),
	_secret VARCHAR(256),
	_identifier VARCHAR(256)
)
RETURNS BOOLEAN AS 
$$
BEGIN
    IF EXISTS(SELECT 1 FROM "user"."credential" WHERE id = _id) THEN
        UPDATE "security".credential
        SET secret = CRYPT(_secret, GEN_SALT('bf'))::VARCHAR(256),
            identifier = _identifier,
            active = TRUE,
            failed_attempts = 0,
            last_accessed = now() AT TIME ZONE 'uct'
        WHERE id = _id;
    ELSE
        INSERT INTO "user"."credential" (id, secret, identifier, active, failed_attempts)
        VALUES (_id, CRYPT(_secret, GEN_SALT('bf'))::VARCHAR(256), _identifier, TRUE, 0);
    END IF;
    RETURN TRUE;
END
$$ LANGUAGE plpgsql;
