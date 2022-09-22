CREATE OR REPLACE FUNCTION "user".identifier_store(
    _id uuid, -- NULLABLE, If supplied try update else insert
    _identifier VARCHAR(256),
	_first_name VARCHAR(256),
	_last_name VARCHAR(256),
	_email VARCHAR(256),
	_picture VARCHAR(256),
    _work_from_home BOOLEAN,
    _parking parking.type,
    _office_days INTEGER,
    _preferred_start_time TIME WITHOUT TIME ZONE,
    _preferred_end_time TIME WITHOUT TIME ZONE,
    _preferred_desk uuid DEFAULT NULL,
    _building_id uuid DEFAULT NULL
)
RETURNS uuid AS 
$$
DECLARE
	__id uuid;
BEGIN
    IF EXISTS(SELECT 1 FROM "user".identifier WHERE identifier = _identifier AND id = _id) THEN
        UPDATE "user".identifier
        SET first_name = _first_name,
            last_name = _last_name,
            email = _email,
            picture = _picture,
            work_from_home = _work_from_home,
            parking = _parking,
            office_days = _office_days,
            preferred_start_time = _preferred_start_time,
            preferred_end_time = _preferred_end_time,
            preferred_desk = _preferred_desk,
            building_id = _building_id
        WHERE identifier = _identifier
        RETURNING identifier.id INTO __id;
    ELSE
        INSERT INTO "user".identifier (id, identifier, first_name, last_name, email, picture, work_from_home, parking, office_days, preferred_start_time, preferred_end_time, preferred_desk, building_id)
        VALUES (COALESCE(_id, uuid_generate_v4()), _identifier, _first_name, _last_name, _email, _picture, _work_from_home, _parking, _office_days, _preferred_start_time, _preferred_end_time, _preferred_desk, _building_id)
        RETURNING identifier.id INTO __id;
    END IF;
    RETURN __id;
END
$$ LANGUAGE plpgsql;
