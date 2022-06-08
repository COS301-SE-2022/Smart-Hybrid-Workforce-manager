CREATE OR REPLACE FUNCTION "user".identifier_find(
    _id uuid DEFAULT NULL,
	_identifier VARCHAR(256) DEFAULT NULL,
    _first_name VARCHAR(256) DEFAULT NULL,
    _last_name VARCHAR(256) DEFAULT NULL,
    _email VARCHAR(256) DEFAULT NULL,
    _picture VARCHAR(256) DEFAULT NULL,
    _date_created TIMESTAMP DEFAULT NULL,
    _work_from_home BOOLEAN DEFAULT NULL,
    _parking parking.type DEFAULT NULL,
    _office_days INTEGER DEFAULT NULL,
    _preferred_start_time TIME WITHOUT TIME ZONE DEFAULT NULL,
    _preferred_end_time TIME WITHOUT TIME ZONE DEFAULT NULL
)
RETURNS TABLE (
    id uuid,
	identifier VARCHAR(256),
	first_name VARCHAR(256),
	last_name VARCHAR(256),
	email VARCHAR(256),
	picture VARCHAR(256),
    date_created TIMESTAMP,
    work_from_home BOOLEAN,
    parking parking.type,
    office_days INTEGER,
    preferred_start_time TIME WITHOUT TIME ZONE,
    preferred_end_time TIME WITHOUT TIME ZONE
) AS 
$$
BEGIN
    RETURN QUERY
    SELECT i.id, i.identifier, i.first_name, i.last_name, i.email, i.picture, i.date_created, i.work_from_home, i.parking, i.office_days, i.preferred_start_time, i.preferred_end_time
    FROM "user".identifier as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_identifier IS NULL OR i.identifier = _identifier)
    AND (_first_name IS NULL OR i.first_name = _first_name)
    AND (_last_name IS NULL OR i.last_name = _last_name)
    AND (_email IS NULL OR i.email = _email)
    AND (_picture IS NULL OR i.picture = _picture)
    AND (_date_created IS NULL OR i.date_created >= _date_created)
    AND (_work_from_home IS NULL OR i.work_from_home = _work_from_home)
    AND (_parking IS NULL OR i.parking = _parking)
    AND (_office_days IS NULL OR i.office_days = _office_days)
    AND (_preferred_start_time IS NULL OR i.preferred_start_time = _preferred_start_time)
    AND (_preferred_end_time IS NULL OR i.preferred_end_time = _preferred_end_time);
END
$$ LANGUAGE plpgsql;
