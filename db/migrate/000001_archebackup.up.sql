CREATE SCHEMA booking;


ALTER SCHEMA booking OWNER TO admin;

--
-- TOC entry 13 (class 2615 OID 16559)
-- Name: permission; Type: SCHEMA; Schema: -; Owner: admin
--

CREATE SCHEMA permission;


ALTER SCHEMA permission OWNER TO admin;

--
-- TOC entry 12 (class 2615 OID 16637)
-- Name: resource; Type: SCHEMA; Schema: -; Owner: admin
--

CREATE SCHEMA resource;


ALTER SCHEMA resource OWNER TO admin;

--
-- TOC entry 7 (class 2615 OID 16526)
-- Name: role; Type: SCHEMA; Schema: -; Owner: admin
--

CREATE SCHEMA role;


ALTER SCHEMA role OWNER TO admin;

--
-- TOC entry 10 (class 2615 OID 16476)
-- Name: team; Type: SCHEMA; Schema: -; Owner: admin
--

CREATE SCHEMA team;


ALTER SCHEMA team OWNER TO admin;

--
-- TOC entry 11 (class 2615 OID 16385)
-- Name: user; Type: SCHEMA; Schema: -; Owner: admin
--

CREATE SCHEMA "user";


ALTER SCHEMA "user" OWNER TO admin;

--
-- TOC entry 2 (class 3079 OID 16386)
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- TOC entry 3568 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- TOC entry 3 (class 3079 OID 16423)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 3569 (class 0 OID 0)
-- Dependencies: 3
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- TOC entry 958 (class 1247 OID 16570)
-- Name: category; Type: TYPE; Schema: permission; Owner: admin
--

CREATE TYPE permission.category AS ENUM (
    'USER',
    'BOOKING',
    'PERMISSION',
    'ROLE',
    'TEAM',
    'RESOURCE'
);


ALTER TYPE permission.category OWNER TO admin;

--
-- TOC entry 961 (class 1247 OID 16584)
-- Name: tenant; Type: TYPE; Schema: permission; Owner: admin
--

CREATE TYPE permission.tenant AS ENUM (
    'ROLE',
    'USER',
    'TEAM',
    'ASSOCIATION',
    'PERMISSION',
    'BUILDING',
    'ROOM',
    'ROOMASSOCIATION',
    'IDENTIFIER',
    'NA'
);


ALTER TYPE permission.tenant OWNER TO admin;

--
-- TOC entry 955 (class 1247 OID 16561)
-- Name: type; Type: TYPE; Schema: permission; Owner: admin
--

CREATE TYPE permission.type AS ENUM (
    'CREATE',
    'DELETE',
    'VIEW',
    'EDIT'
);


ALTER TYPE permission.type OWNER TO admin;

--
-- TOC entry 970 (class 1247 OID 16639)
-- Name: type; Type: TYPE; Schema: resource; Owner: admin
--

CREATE TYPE resource.type AS ENUM (
    'PARKING',
    'DESK',
    'MEETINGROOM'
);


ALTER TYPE resource.type OWNER TO admin;

--
-- TOC entry 934 (class 1247 OID 16450)
-- Name: credential_type; Type: TYPE; Schema: user; Owner: admin
--

CREATE TYPE "user".credential_type AS ENUM (
    'federated',
    'local'
);


ALTER TYPE "user".credential_type OWNER TO admin;

--
-- TOC entry 327 (class 1255 OID 16738)
-- Name: identifier_find(uuid, uuid, resource.type, uuid, uuid, timestamp without time zone, timestamp without time zone, boolean, timestamp without time zone, jsonb); Type: FUNCTION; Schema: booking; Owner: admin
--

CREATE FUNCTION booking.identifier_find(_id uuid DEFAULT NULL::uuid, _user_id uuid DEFAULT NULL::uuid, _resource_type resource.type DEFAULT NULL::resource.type, _resource_preference_id uuid DEFAULT NULL::uuid, _resource_id uuid DEFAULT NULL::uuid, _start timestamp without time zone DEFAULT NULL::timestamp without time zone, _end timestamp without time zone DEFAULT NULL::timestamp without time zone, _booked boolean DEFAULT NULL::boolean, _date_created timestamp without time zone DEFAULT NULL::timestamp without time zone, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(id uuid, user_id uuid, resource_type resource.type, resource_preference_id uuid, resource_id uuid, start timestamp without time zone, "end" timestamp without time zone, booked boolean, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_users AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'BOOKING'::permission.category
        AND permission_tenant = 'USER'::permission.tenant
    )
    SELECT i.id, i.user_id, i.resource_type, i.resource_preference_id, i.resource_id, i.start, i."end", i.booked, i.date_created
    FROM booking.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_users WHERE permission_tenant_id is null) OR i.user_id = ANY(SELECT * FROM permitted_users))
    AND (_id IS NULL OR i.id = _id)
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_resource_type IS NULL OR i.resource_type = _resource_type)
    AND (_resource_preference_id IS NULL OR i.resource_preference_id = _resource_preference_id)
    AND (_resource_id IS NULL OR i.resource_id = _resource_id)
    AND (_start IS NULL OR i.start >= _start)
    AND (_end IS NULL OR i."end" <= _end)
    AND (_booked IS NULL OR i.booked = _booked)
    AND (_date_created IS NULL OR i.date_created >= _date_created);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION booking.identifier_find(_id uuid, _user_id uuid, _resource_type resource.type, _resource_preference_id uuid, _resource_id uuid, _start timestamp without time zone, _end timestamp without time zone, _booked boolean, _date_created timestamp without time zone, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 328 (class 1255 OID 16739)
-- Name: identifier_remove(uuid); Type: FUNCTION; Schema: booking; Owner: admin
--

CREATE FUNCTION booking.identifier_remove(_id uuid) RETURNS TABLE(id uuid, user_id uuid, resource_type resource.type, resource_preference_id uuid, resource_id uuid, start timestamp without time zone, "end" timestamp without time zone, booked boolean, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM booking.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_booking'
            USING HINT = 'Please check the provided booking id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM booking.identifier as a WHERE a.id = _id 
    RETURNING *;

END
$$;


ALTER FUNCTION booking.identifier_remove(_id uuid) OWNER TO admin;

--
-- TOC entry 329 (class 1255 OID 16740)
-- Name: identifier_store(uuid, uuid, resource.type, uuid, uuid, timestamp without time zone, timestamp without time zone, boolean); Type: FUNCTION; Schema: booking; Owner: admin
--

CREATE FUNCTION booking.identifier_store(_id uuid, _user_id uuid, _resource_type resource.type, _resource_preference_id uuid, _resource_id uuid, _start timestamp without time zone, _end timestamp without time zone, _booked boolean DEFAULT NULL::boolean) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
DECLARE
	__id uuid;
BEGIN
	IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM booking.identifier WHERE id = _id)) THEN
        UPDATE booking.identifier
        SET user_id = _user_id,
            resource_type = _resource_type,
            resource_preference_id = _resource_preference_id,
            resource_id = _resource_id,
            start = _start,
            "end" = _end,
            booked = _booked
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO booking.identifier(id, user_id, resource_type, resource_preference_id, start, "end")
    	VALUES (COALESCE(_id, uuid_generate_v4()), _user_id, _resource_type, _resource_preference_id, _start, _end)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$;


ALTER FUNCTION booking.identifier_store(_id uuid, _user_id uuid, _resource_type resource.type, _resource_preference_id uuid, _resource_id uuid, _start timestamp without time zone, _end timestamp without time zone, _booked boolean) OWNER TO admin;

--
-- TOC entry 303 (class 1255 OID 16631)
-- Name: role_find(uuid, permission.type, permission.category, permission.tenant, uuid, timestamp without time zone, jsonb); Type: FUNCTION; Schema: permission; Owner: admin
--

CREATE FUNCTION permission.role_find(_role_id uuid DEFAULT NULL::uuid, _permission_type permission.type DEFAULT NULL::permission.type, _permission_category permission.category DEFAULT NULL::permission.category, _permission_tenant permission.tenant DEFAULT NULL::permission.tenant, _permission_tenant_id uuid DEFAULT NULL::uuid, _date_added timestamp without time zone DEFAULT NULL::timestamp without time zone, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(role_id uuid, permission_type permission.type, permission_category permission.category, permission_tenant permission.tenant, permission_tenant_id uuid, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_roles AS (
        SELECT a.permission_tenant_id FROM _permissions_table as a
        WHERE a.permission_type = 'VIEW'::permission.type
        AND a.permission_category = 'PERMISSION'::permission.category
        AND a.permission_tenant = 'ROLE'::permission.tenant
    )
    SELECT i.role_id, i.permission_type, i.permission_category, i.permission_tenant, i.permission_tenant_id, i.date_added
    FROM permission.role as i
    WHERE (EXISTS(SELECT 1 FROM permitted_roles as b WHERE b.permission_tenant_id is null) OR i.role_id = ANY(SELECT * FROM permitted_roles))
    AND (_role_id IS NULL OR i.role_id = _role_id)
    AND (_permission_type IS NULL OR i.permission_type = _permission_type)
    AND (_permission_category IS NULL OR i.permission_category = _permission_category)
    AND (_permission_tenant IS NULL OR i.permission_tenant = _permission_tenant)
    AND (_permission_tenant_id IS NULL OR i.permission_tenant_id = _permission_tenant_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION permission.role_find(_role_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid, _date_added timestamp without time zone, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 312 (class 1255 OID 16632)
-- Name: role_remove(uuid, permission.type, permission.category, permission.tenant, uuid); Type: FUNCTION; Schema: permission; Owner: admin
--

CREATE FUNCTION permission.role_remove(_role_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid) RETURNS TABLE(role_id uuid, permission_type permission.type, permission_category permission.category, permission_tenant permission.tenant, permission_tenant_id uuid, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM permission.role as b
        WHERE b.role_id = _role_id
        AND b.permission_type = _permission_type
        AND b.permission_category = _permission_category
        AND b.permission_tenant = _permission_tenant
        AND b.permission_tenant_id = _permission_tenant_id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_role'
            USING HINT = 'Please check the provided role parameters';
    END IF;

    RETURN QUERY
    DELETE FROM permission.role as a
    WHERE a.role_id = _role_id
    AND a.permission_type = _permission_type
    AND a.permission_category = _permission_category
    AND a.permission_tenant = _permission_tenant
    AND a.permission_tenant_id = _permission_tenant_id
    RETURNING a.role_id, a.permission_type, a.permission_category, a.permission_tenant, a.permission_tenant_id, a.date_added;

END
$$;


ALTER FUNCTION permission.role_remove(_role_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid) OWNER TO admin;

--
-- TOC entry 311 (class 1255 OID 16633)
-- Name: role_store(uuid, permission.type, permission.category, permission.tenant, uuid); Type: FUNCTION; Schema: permission; Owner: admin
--

CREATE FUNCTION permission.role_store(_role_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN

    INSERT INTO permission.role(role_id, permission_type, permission_category, permission_tenant, permission_tenant_id)
    VALUES (_role_id, _permission_type, _permission_category, _permission_tenant, _permission_tenant_id);

	RETURN true;
END
$$;


ALTER FUNCTION permission.role_store(_role_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid) OWNER TO admin;

--
-- TOC entry 313 (class 1255 OID 16634)
-- Name: user_find(uuid, permission.type, permission.category, permission.tenant, uuid, timestamp without time zone, jsonb); Type: FUNCTION; Schema: permission; Owner: admin
--

CREATE FUNCTION permission.user_find(_user_id uuid DEFAULT NULL::uuid, _permission_type permission.type DEFAULT NULL::permission.type, _permission_category permission.category DEFAULT NULL::permission.category, _permission_tenant permission.tenant DEFAULT NULL::permission.tenant, _permission_tenant_id uuid DEFAULT NULL::uuid, _date_added timestamp without time zone DEFAULT NULL::timestamp without time zone, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(user_id uuid, permission_type permission.type, permission_category permission.category, permission_tenant permission.tenant, permission_tenant_id uuid, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_users AS (
        SELECT a.permission_tenant_id FROM _permissions_table as a
        WHERE a.permission_type = 'VIEW'::permission.type
        AND a.permission_category = 'PERMISSION'::permission.category
        AND a.permission_tenant = 'USER'::permission.tenant
    )
    SELECT i.user_id, i.permission_type, i.permission_category, i.permission_tenant, i.permission_tenant_id, i.date_added
    FROM permission.user as i
    WHERE (EXISTS(SELECT 1 FROM permitted_users as b WHERE b.permission_tenant_id is null) OR i.user_id = ANY(SELECT * FROM permitted_users))
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_permission_type IS NULL OR i.permission_type = _permission_type)
    AND (_permission_category IS NULL OR i.permission_category = _permission_category)
    AND (_permission_tenant IS NULL OR i.permission_tenant = _permission_tenant)
    AND (_permission_tenant_id IS NULL OR i.permission_tenant_id = _permission_tenant_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION permission.user_find(_user_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid, _date_added timestamp without time zone, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 314 (class 1255 OID 16635)
-- Name: user_remove(uuid, permission.type, permission.category, permission.tenant, uuid); Type: FUNCTION; Schema: permission; Owner: admin
--

CREATE FUNCTION permission.user_remove(_user_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid) RETURNS TABLE(user_id uuid, permission_type permission.type, permission_category permission.category, permission_tenant permission.tenant, permission_tenant_id uuid, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM permission.user as b
        WHERE b.user_id = _user_id
        AND b.permission_type = _permission_type
        AND b.permission_category = _permission_category
        AND b.permission_tenant = _permission_tenant
        AND b.permission_tenant_id = _permission_tenant_id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_user'
            USING HINT = 'Please check the provided user parameters';
    END IF;

    RETURN QUERY
    DELETE FROM permission.user as a
    WHERE a.user_id = _user_id
    AND a.permission_type = _permission_type
    AND a.permission_category = _permission_category
    AND a.permission_tenant = _permission_tenant
    AND a.permission_tenant_id = _permission_tenant_id
    RETURNING a.user_id, a.permission_type, a.permission_category, a.permission_tenant, a.permission_tenant_id, a.date_added;

END
$$;


ALTER FUNCTION permission.user_remove(_user_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid) OWNER TO admin;

--
-- TOC entry 277 (class 1255 OID 16636)
-- Name: user_store(uuid, permission.type, permission.category, permission.tenant, uuid); Type: FUNCTION; Schema: permission; Owner: admin
--

CREATE FUNCTION permission.user_store(_user_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN

    INSERT INTO permission.user(user_id, permission_type, permission_category, permission_tenant, permission_tenant_id)
    VALUES (_user_id, _permission_type, _permission_category, _permission_tenant, _permission_tenant_id);

	RETURN true;
END
$$;


ALTER FUNCTION permission.user_store(_user_id uuid, _permission_type permission.type, _permission_category permission.category, _permission_tenant permission.tenant, _permission_tenant_id uuid) OWNER TO admin;

--
-- TOC entry 315 (class 1255 OID 16702)
-- Name: building_find(uuid, character varying, character varying, character varying, jsonb); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.building_find(_id uuid DEFAULT NULL::uuid, _name character varying DEFAULT NULL::character varying, _location character varying DEFAULT NULL::character varying, _dimension character varying DEFAULT NULL::character varying, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(id uuid, name character varying, location character varying, dimension character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_buildings AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'RESOURCE'::permission.category
        AND permission_tenant = 'BUILDING'::permission.tenant
    )
    SELECT i.id, i.name, i.location, i.dimension
    FROM resource.building as i
    WHERE (EXISTS(SELECT 1 FROM permitted_buildings WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_buildings))
    AND (_id IS NULL OR i.id = _id)
    AND (_name IS NULL OR i.name = _name)
    AND (_location IS NULL OR i.location = _location)
    AND (_dimension IS NULL OR i.dimension = _dimension);
    
    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION resource.building_find(_id uuid, _name character varying, _location character varying, _dimension character varying, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 316 (class 1255 OID 16703)
-- Name: building_remove(uuid); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.building_remove(_id uuid) RETURNS TABLE(id uuid, name character varying, location character varying, dimension character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM resource.building as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_building'
            USING HINT = 'Please check the provided building id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM resource.building as a WHERE a.id = _id
    RETURNING *;

END
$$;


ALTER FUNCTION resource.building_remove(_id uuid) OWNER TO admin;

--
-- TOC entry 317 (class 1255 OID 16704)
-- Name: building_store(uuid, character varying, character varying, character varying); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.building_store(_id uuid, _name character varying, _location character varying, _dimension character varying) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
DECLARE
	__id uuid;
BEGIN
    IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.building WHERE id = _id)) THEN
        UPDATE resource.building
        SET name = _name,
            location = _location,
            dimension = _dimension
        WHERE id = _id
		RETURNING building.id INTO __id;
    ELSE
    	INSERT INTO resource.building(id, name, location, dimension)
        VALUES (COALESCE(_id, uuid_generate_v4()), _name, _location, _dimension)
		RETURNING building.id INTO __id;
    END IF;
	RETURN __id;
END
$$;


ALTER FUNCTION resource.building_store(_id uuid, _name character varying, _location character varying, _dimension character varying) OWNER TO admin;

--
-- TOC entry 319 (class 1255 OID 16705)
-- Name: identifier_find(uuid, uuid, character varying, character varying, uuid, resource.type, timestamp without time zone, jsonb); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.identifier_find(_id uuid DEFAULT NULL::uuid, _room_id uuid DEFAULT NULL::uuid, _name character varying DEFAULT NULL::character varying, _location character varying DEFAULT NULL::character varying, _role_id uuid DEFAULT NULL::uuid, _resource_type resource.type DEFAULT NULL::resource.type, _date_created timestamp without time zone DEFAULT NULL::timestamp without time zone, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(id uuid, room_id uuid, name character varying, location character varying, role_id uuid, resource_type resource.type, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_identifiers AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'RESOURCE'::permission.category
        AND permission_tenant = 'IDENTIFIER'::permission.tenant
    )
    SELECT i.id, i.room_id, i.name, i.location, i.role_id, i.resource_type, i.date_created
    FROM resource.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_identifiers WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_identifiers))
    AND (_id IS NULL OR i.id = _id)
    AND (_room_id IS NULL OR i.room_id = _room_id)
    AND (_name IS NULL OR i.name = _name)
    AND (_location IS NULL OR i.location = _location)
    AND (_role_id IS NULL OR i.role_id = _role_id)
    AND (_resource_type IS NULL OR i.resource_type = _resource_type)
    AND (_date_created IS NULL OR i.date_created >= _date_created);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION resource.identifier_find(_id uuid, _room_id uuid, _name character varying, _location character varying, _role_id uuid, _resource_type resource.type, _date_created timestamp without time zone, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 320 (class 1255 OID 16706)
-- Name: identifier_remove(uuid); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.identifier_remove(_id uuid) RETURNS TABLE(id uuid, room_id uuid, name character varying, location character varying, role_id uuid, resource_type resource.type, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM resource.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_resource'
            USING HINT = 'Please check the provided resource id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM resource.identifier as a WHERE a.id = _id
    RETURNING *;

END
$$;


ALTER FUNCTION resource.identifier_remove(_id uuid) OWNER TO admin;

--
-- TOC entry 321 (class 1255 OID 16707)
-- Name: identifier_store(uuid, uuid, character varying, character varying, uuid, resource.type); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.identifier_store(_id uuid, _room_id uuid, _name character varying, _location character varying, _role_id uuid, _resource_type resource.type) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
DECLARE
	__id uuid;
BEGIN
    IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.identifier WHERE id = _id)) THEN
        UPDATE resource.identifier
        SET room_id = _room_id,
            name = _name,
            location = _location,            
            role_id = _role_id,
            resource_type = _resource_type
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO resource.identifier(id, room_id, name, location, role_id, resource_type)
        VALUES (COALESCE(_id, uuid_generate_v4()), _room_id, _name, _location, _role_id, _resource_type)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$;


ALTER FUNCTION resource.identifier_store(_id uuid, _room_id uuid, _name character varying, _location character varying, _role_id uuid, _resource_type resource.type) OWNER TO admin;

--
-- TOC entry 322 (class 1255 OID 16708)
-- Name: room_association_find(uuid, uuid, jsonb); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.room_association_find(_room_id uuid DEFAULT NULL::uuid, _room_id_association uuid DEFAULT NULL::uuid, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(room_id uuid, room_id_association uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_associations AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'RESOURCE'::permission.category
        AND permission_tenant = 'ROOMASSOCIATION'::permission.tenant
    )
    SELECT i.room_id, i.room_id_association
    FROM resource.room_association as i
    WHERE (EXISTS(SELECT 1 FROM permitted_associations WHERE permission_tenant_id is null) OR i.room_id = ANY(SELECT * FROM permitted_associations))
    AND (_room_id IS NULL OR i.room_id = _room_id)
    AND (_room_id_association IS NULL OR i.room_id_association = _room_id_association);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION resource.room_association_find(_room_id uuid, _room_id_association uuid, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 323 (class 1255 OID 16709)
-- Name: room_association_remove(uuid, uuid); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.room_association_remove(_room_id uuid, _room_id_association uuid) RETURNS TABLE(room_id uuid, room_id_association uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM resource.room_association as b
        WHERE b.room_id = _room_id
        AND b.room_id_association = _room_id_association
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_room_association'
            USING HINT = 'Please check the provided room association parameters';
    END IF;

    RETURN QUERY
    DELETE FROM resource.room_association as a 
    WHERE a.room_id = _room_id
    AND a.room_id_association = _room_id_association
    RETURNING *;

END
$$;


ALTER FUNCTION resource.room_association_remove(_room_id uuid, _room_id_association uuid) OWNER TO admin;

--
-- TOC entry 324 (class 1255 OID 16710)
-- Name: room_association_store(uuid, uuid); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.room_association_store(_room_id uuid, _room_id_association uuid) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (_room_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.room_association WHERE room_id = _room_id)) THEN
        UPDATE resource.room_association
        SET room_id_association = _room_id_association
        WHERE room_id = _room_id;
    ELSE
    	INSERT INTO resource.room_association(room_id, room_id_association)
    	VALUES (_room_id, _room_id_association);
    END IF;
	RETURN true;
END
$$;


ALTER FUNCTION resource.room_association_store(_room_id uuid, _room_id_association uuid) OWNER TO admin;

--
-- TOC entry 325 (class 1255 OID 16711)
-- Name: room_find(uuid, uuid, character varying, character varying, character varying, jsonb); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.room_find(_id uuid DEFAULT NULL::uuid, _building_id uuid DEFAULT NULL::uuid, _name character varying DEFAULT NULL::character varying, _location character varying DEFAULT NULL::character varying, _dimension character varying DEFAULT NULL::character varying, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(id uuid, building_id uuid, name character varying, location character varying, dimension character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_rooms AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'RESOURCE'::permission.category
        AND permission_tenant = 'ROOM'::permission.tenant
    )
    SELECT i.id, i.building_id, i.name, i.location, i.dimension
    FROM resource.room as i
    WHERE (EXISTS(SELECT 1 FROM permitted_rooms WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_rooms))
    AND (_id IS NULL OR i.id = _id)
    AND (_building_id IS NULL OR i.building_id = _building_id)
    AND (_name IS NULL OR i.name = _name)
    AND (_location IS NULL OR i.location = _location)
    AND (_dimension IS NULL OR i.dimension = _dimension);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION resource.room_find(_id uuid, _building_id uuid, _name character varying, _location character varying, _dimension character varying, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 318 (class 1255 OID 16712)
-- Name: room_remove(uuid); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.room_remove(_id uuid) RETURNS TABLE(id uuid, building_id uuid, name character varying, location character varying, dimension character varying)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM resource.room as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_resource'
            USING HINT = 'Please check the provided resource id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM resource.room as a WHERE a.id = _id
    RETURNING *;

END
$$;


ALTER FUNCTION resource.room_remove(_id uuid) OWNER TO admin;

--
-- TOC entry 326 (class 1255 OID 16713)
-- Name: room_store(uuid, uuid, character varying, character varying, character varying); Type: FUNCTION; Schema: resource; Owner: admin
--

CREATE FUNCTION resource.room_store(_id uuid, _building_id uuid, _name character varying, _location character varying, _dimension character varying) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
DECLARE
	__id uuid;
BEGIN
    IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM resource.room WHERE id = _id)) THEN
        UPDATE resource.room
        SET name = _name,
            building_id = _building_id,
            location = _location,
            dimension = _dimension
        WHERE id = _id
		RETURNING room.id INTO __id;
    ELSE
    	INSERT INTO resource.room(id, building_id, name, location, dimension)
        VALUES (COALESCE(_id, uuid_generate_v4()), _building_id, _name, _location, _dimension)
		RETURNING room.id INTO __id;
    END IF;
	RETURN __id;
END
$$;


ALTER FUNCTION resource.room_store(_id uuid, _building_id uuid, _name character varying, _location character varying, _dimension character varying) OWNER TO admin;

--
-- TOC entry 306 (class 1255 OID 16553)
-- Name: identifier_find(uuid, character varying, timestamp without time zone, jsonb); Type: FUNCTION; Schema: role; Owner: admin
--

CREATE FUNCTION role.identifier_find(_id uuid DEFAULT NULL::uuid, _role_name character varying DEFAULT NULL::character varying, _date_added timestamp without time zone DEFAULT NULL::timestamp without time zone, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(id uuid, role_name character varying, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_roles AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'ROLE'::permission.category
        AND permission_tenant = 'IDENTIFIER'::permission.tenant
    )
    SELECT i.id, i.role_name, i.date_added
    FROM role.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_roles WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_roles))
    AND (_id IS NULL OR i.id = _id)
    AND (_role_name IS NULL OR i.role_name = _role_name)
    AND (_date_added IS NULL OR i.date_added >= _date_added);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION role.identifier_find(_id uuid, _role_name character varying, _date_added timestamp without time zone, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 307 (class 1255 OID 16554)
-- Name: identifier_remove(uuid); Type: FUNCTION; Schema: role; Owner: admin
--

CREATE FUNCTION role.identifier_remove(_id uuid) RETURNS TABLE(id uuid, role_name character varying, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM role.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_role'
            USING HINT = 'Please check the provided role id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM role.identifier as a WHERE a.id = _id
    RETURNING *;

END
$$;


ALTER FUNCTION role.identifier_remove(_id uuid) OWNER TO admin;

--
-- TOC entry 308 (class 1255 OID 16555)
-- Name: identifier_store(uuid, character varying); Type: FUNCTION; Schema: role; Owner: admin
--

CREATE FUNCTION role.identifier_store(_id uuid, _role_name character varying) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
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
$$;


ALTER FUNCTION role.identifier_store(_id uuid, _role_name character varying) OWNER TO admin;

--
-- TOC entry 309 (class 1255 OID 16556)
-- Name: user_find(uuid, uuid, timestamp without time zone, jsonb); Type: FUNCTION; Schema: role; Owner: admin
--

CREATE FUNCTION role.user_find(_role_id uuid DEFAULT NULL::uuid, _user_id uuid DEFAULT NULL::uuid, _date_added timestamp without time zone DEFAULT NULL::timestamp without time zone, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(role_id uuid, user_id uuid, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_users AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'USER'::permission.category
        AND permission_tenant = 'ROLE'::permission.tenant
    ),
    permitted_roles AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'ROLE'::permission.category
        AND permission_tenant = 'USER'::permission.tenant
    )
    SELECT i.role_id, i.user_id, i.date_added
    FROM role.user as i
    WHERE (EXISTS(SELECT 1 FROM permitted_users WHERE permission_tenant_id is null) OR i.user_id = ANY(SELECT * FROM permitted_users))
    AND (EXISTS(SELECT 1 FROM permitted_roles WHERE permission_tenant_id is null) OR i.role_id = ANY(SELECT * FROM permitted_roles))
    AND (_role_id IS NULL OR i.role_id = _role_id)
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION role.user_find(_role_id uuid, _user_id uuid, _date_added timestamp without time zone, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 310 (class 1255 OID 16557)
-- Name: user_remove(uuid, uuid); Type: FUNCTION; Schema: role; Owner: admin
--

CREATE FUNCTION role.user_remove(_role_id uuid, _user_id uuid) RETURNS TABLE(role_id uuid, user_id uuid, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
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

    RETURN QUERY
    DELETE FROM role.user as a
    WHERE a.role_id = _role_id
    AND a.user_id = _user_id
    RETURNING *;

END
$$;


ALTER FUNCTION role.user_remove(_role_id uuid, _user_id uuid) OWNER TO admin;

--
-- TOC entry 292 (class 1255 OID 16558)
-- Name: user_store(uuid, uuid); Type: FUNCTION; Schema: role; Owner: admin
--

CREATE FUNCTION role.user_store(_role_id uuid, _user_id uuid) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO role.user(role_id, user_id)
    VALUES (_role_id, _user_id);
	RETURN true;
END
$$;


ALTER FUNCTION role.user_store(_role_id uuid, _user_id uuid) OWNER TO admin;

--
-- TOC entry 298 (class 1255 OID 16517)
-- Name: association_find(uuid, uuid, jsonb); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.association_find(_team_id uuid DEFAULT NULL::uuid, _team_id_association uuid DEFAULT NULL::uuid, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(team_id uuid, team_id_association uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_teams AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'TEAM'::permission.category
        AND permission_tenant = 'ASSOCIATION'::permission.tenant
    )
    SELECT i.team_id, i.team_id_association
    FROM team.association as i
    WHERE (EXISTS(SELECT 1 FROM permitted_teams WHERE permission_tenant_id is null) OR i.team_id = ANY(SELECT * FROM permitted_teams))
    AND (EXISTS(SELECT 1 FROM permitted_teams WHERE permission_tenant_id is null) OR i._team_id_association = ANY(SELECT * FROM permitted_teams))
    AND (_team_id IS NULL OR i.team_id = _team_id)
    AND (_team_id_association IS NULL OR i.team_id_association = _team_id_association);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION team.association_find(_team_id uuid, _team_id_association uuid, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 299 (class 1255 OID 16518)
-- Name: association_remove(uuid, uuid); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.association_remove(_team_id uuid, _team_id_association uuid) RETURNS TABLE(team_id uuid, team_id_association uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM team.association as b
        WHERE b.team_id = _team_id
        AND b.team_id_association = _team_id_association
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_team_association'
            USING HINT = 'Please check the provided team and association parameter';
    END IF;

    RETURN QUERY
    DELETE FROM team.association as a
    WHERE a.team_id = _team_id
    AND a.team_id_association = _team_id_association
    RETURNING *;

END
$$;


ALTER FUNCTION team.association_remove(_team_id uuid, _team_id_association uuid) OWNER TO admin;

--
-- TOC entry 278 (class 1255 OID 16519)
-- Name: association_store(uuid, uuid); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.association_store(_team_id uuid, _team_id_association uuid) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
	IF (_team_id IS NOT NULL AND EXISTS(SELECT 1 FROM team.association WHERE team_id = _team_id)) THEN
        UPDATE team.association
        SET team_id_association = _team_id_association
        WHERE team_id = _team_id;
    ELSE
    	INSERT INTO team.association(team_id, team_id_association)
    	VALUES (_team_id, _team_id_association);
    END IF;
	RETURN true;
END
$$;


ALTER FUNCTION team.association_store(_team_id uuid, _team_id_association uuid) OWNER TO admin;

--
-- TOC entry 300 (class 1255 OID 16520)
-- Name: identifier_find(uuid, character varying, character varying, integer, character varying, timestamp without time zone, jsonb); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.identifier_find(_id uuid DEFAULT NULL::uuid, _name character varying DEFAULT NULL::character varying, _description character varying DEFAULT NULL::character varying, _capacity integer DEFAULT NULL::integer, _picture character varying DEFAULT NULL::character varying, _date_created timestamp without time zone DEFAULT NULL::timestamp without time zone, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(id uuid, name character varying, description character varying, capacity integer, picture character varying, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_teams AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'TEAM'::permission.category
        AND permission_tenant = 'IDENTIFIER'::permission.tenant
    )
    SELECT i.id, i.name, i.description, i.capacity, i.picture, i.date_created
    FROM team.identifier as i
    WHERE (EXISTS(SELECT 1 FROM permitted_teams WHERE permission_tenant_id is null) OR i.id = ANY(SELECT * FROM permitted_teams))
    AND (_id IS NULL OR i.id = _id)
    AND (_name IS NULL OR i.name = _name)
    AND (_description IS NULL OR i.description = _description)
    AND (_capacity IS NULL OR i.capacity = _capacity)
    AND (_picture IS NULL OR i.picture = _picture)
    AND (_date_created IS NULL OR i.date_created >= _date_created);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION team.identifier_find(_id uuid, _name character varying, _description character varying, _capacity integer, _picture character varying, _date_created timestamp without time zone, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 301 (class 1255 OID 16521)
-- Name: identifier_remove(uuid); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.identifier_remove(_id uuid) RETURNS TABLE(id uuid, name character varying, description character varying, capacity integer, picture character varying, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM team.identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_team'
            USING HINT = 'Please check the provided team id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM team.identifier as a WHERE a.id = _id
    RETURNING *;

END
$$;


ALTER FUNCTION team.identifier_remove(_id uuid) OWNER TO admin;

--
-- TOC entry 302 (class 1255 OID 16522)
-- Name: identifier_store(uuid, character varying, character varying, integer, character varying); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.identifier_store(_id uuid, _name character varying, _description character varying, _capacity integer, _picture character varying) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
DECLARE
	__id uuid;
BEGIN
	IF (_id IS NOT NULL AND EXISTS(SELECT 1 FROM team.identifier WHERE id = _id)) THEN
        UPDATE team.identifier
        SET name = _name,
            description = _description,
            capacity = _capacity,
            picture = _picture
        WHERE id = _id
		RETURNING identifier.id INTO __id;
    ELSE
    	INSERT INTO team.identifier(id, name, description, capacity, picture)
    	VALUES (COALESCE(_id, uuid_generate_v4()), _name, _description, _capacity, _picture)
		RETURNING identifier.id INTO __id;
    END IF;
	RETURN __id;
END
$$;


ALTER FUNCTION team.identifier_store(_id uuid, _name character varying, _description character varying, _capacity integer, _picture character varying) OWNER TO admin;

--
-- TOC entry 304 (class 1255 OID 16523)
-- Name: user_find(uuid, uuid, timestamp without time zone, jsonb); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.user_find(_team_id uuid DEFAULT NULL::uuid, _user_id uuid DEFAULT NULL::uuid, _date_added timestamp without time zone DEFAULT NULL::timestamp without time zone, _permissions jsonb DEFAULT NULL::jsonb) RETURNS TABLE(team_id uuid, user_id uuid, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    CREATE TEMP TABLE _permissions_table (
        permission_type permission.type,
        permission_category permission.category,
        permission_tenant permission.tenant,
        permission_tenant_id uuid
    );

    INSERT INTO _permissions_table (
    SELECT
        (jsonb_array_elements(_permissions)->>'permission_type')::permission.type,
        (jsonb_array_elements(_permissions)->>'permission_category')::permission.category,
        (jsonb_array_elements(_permissions)->>'permission_tenant')::permission.tenant,
        (jsonb_array_elements(_permissions)->>'permission_tenant_id')::uuid
    );

    RETURN QUERY
    WITH permitted_users AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'USER'::permission.category
        AND permission_tenant = 'TEAM'::permission.tenant
    ),
    permitted_teams AS (
        SELECT permission_tenant_id FROM _permissions_table 
        WHERE permission_type = 'VIEW'::permission.type
        AND permission_category = 'TEAM'::permission.category
        AND permission_tenant = 'USER'::permission.tenant
    )
    SELECT i.team_id, i.user_id, i.date_added
    FROM team.user as i
    WHERE (EXISTS(SELECT 1 FROM permitted_users WHERE permission_tenant_id is null) OR i.user_id = ANY(SELECT * FROM permitted_users))
    AND (EXISTS(SELECT 1 FROM permitted_teams WHERE permission_tenant_id is null) OR i.team_id = ANY(SELECT * FROM permitted_teams))
    AND (_team_id IS NULL OR i.team_id = _team_id)
    AND (_user_id IS NULL OR i.user_id = _user_id)
    AND (_date_added IS NULL OR i.date_added >= _date_added);

    DROP TABLE _permissions_table;
END
$$;


ALTER FUNCTION team.user_find(_team_id uuid, _user_id uuid, _date_added timestamp without time zone, _permissions jsonb) OWNER TO admin;

--
-- TOC entry 305 (class 1255 OID 16524)
-- Name: user_remove(uuid, uuid); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.user_remove(_team_id uuid, _user_id uuid) RETURNS TABLE(team_id uuid, user_id uuid, date_added timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM team.user as b
        WHERE b.team_id = _team_id
        AND b.user_id = _user_id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_team_user'
            USING HINT = 'Please check the provided team and user parameter';
    END IF;

    RETURN QUERY
    DELETE FROM team.user as a
    WHERE a.team_id = _team_id
    AND a.user_id = _user_id
    RETURNING *;

END
$$;


ALTER FUNCTION team.user_remove(_team_id uuid, _user_id uuid) OWNER TO admin;

--
-- TOC entry 279 (class 1255 OID 16525)
-- Name: user_store(uuid, uuid); Type: FUNCTION; Schema: team; Owner: admin
--

CREATE FUNCTION team.user_store(_team_id uuid, _user_id uuid) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO team.user(team_id, user_id)
    VALUES (_team_id, _user_id);
	RETURN true;
END
$$;


ALTER FUNCTION team.user_store(_team_id uuid, _user_id uuid) OWNER TO admin;

--
-- TOC entry 295 (class 1255 OID 16470)
-- Name: credential_find(character varying, character varying, character varying, "user".credential_type, boolean, integer, timestamp without time zone); Type: FUNCTION; Schema: user; Owner: admin
--

CREATE FUNCTION "user".credential_find(_id character varying DEFAULT NULL::character varying, _secret character varying DEFAULT NULL::character varying, _identifier character varying DEFAULT NULL::character varying, _type "user".credential_type DEFAULT NULL::"user".credential_type, _active boolean DEFAULT NULL::boolean, _failed_attempts integer DEFAULT NULL::integer, _last_accessed timestamp without time zone DEFAULT NULL::timestamp without time zone) RETURNS TABLE(id character varying, secret character varying, identifier character varying, type "user".credential_type, active boolean, failed_attempts integer, last_accessed timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT i.id, i.secret, i.identifier, i."type", i.active, i.failed_attempts, i.last_accessed
    FROM "user".credential as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_secret IS NULL OR i.secret = _secret)
    AND (_identifier IS NULL OR i.identifier = _identifier)
    AND (_type IS NULL OR i."type" = _type)
    AND (_active IS NULL OR i.active = _active)
    AND (_failed_attempts IS NULL OR i.failed_attempts = _failed_attempts)
    AND (_last_accessed IS NULL OR i.last_accessed >= _last_accessed);
END
$$;


ALTER FUNCTION "user".credential_find(_id character varying, _secret character varying, _identifier character varying, _type "user".credential_type, _active boolean, _failed_attempts integer, _last_accessed timestamp without time zone) OWNER TO admin;

--
-- TOC entry 280 (class 1255 OID 16471)
-- Name: credential_remove(uuid); Type: FUNCTION; Schema: user; Owner: admin
--

CREATE FUNCTION "user".credential_remove(_id uuid) RETURNS TABLE(id character varying, secret character varying, identifier character varying, type "user".credential_type, active boolean, failed_attempts integer, last_accessed timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM "user".credential as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_user'
            USING HINT = 'Please check the provided user id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM "user".credential as a WHERE a.id = _id
    RETURNING *;

END
$$;


ALTER FUNCTION "user".credential_remove(_id uuid) OWNER TO admin;

--
-- TOC entry 284 (class 1255 OID 16472)
-- Name: credential_store(character varying, character varying, character varying); Type: FUNCTION; Schema: user; Owner: admin
--

CREATE FUNCTION "user".credential_store(_id character varying, _secret character varying, _identifier character varying) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF EXISTS(SELECT 1 FROM "user"."credential" WHERE id = _id) THEN
        UPDATE "user".credential
        SET secret = CRYPT(_secret, GEN_SALT('bf', 13))::VARCHAR(256),
            identifier = _identifier,
            active = TRUE,
            failed_attempts = 0,
            last_accessed = now() AT TIME ZONE 'uct'
        WHERE id = _id;
    ELSE
        INSERT INTO "user"."credential" (id, secret, identifier, active, failed_attempts)
        VALUES (_id, CRYPT(_secret, GEN_SALT('bf', 13))::VARCHAR(256), _identifier, TRUE, 0);
    END IF;
    RETURN TRUE;
END
$$;


ALTER FUNCTION "user".credential_store(_id character varying, _secret character varying, _identifier character varying) OWNER TO admin;

--
-- TOC entry 296 (class 1255 OID 16473)
-- Name: identifier_find(uuid, character varying, character varying, character varying, character varying, character varying, timestamp without time zone); Type: FUNCTION; Schema: user; Owner: admin
--

CREATE FUNCTION "user".identifier_find(_id uuid DEFAULT NULL::uuid, _identifier character varying DEFAULT NULL::character varying, _first_name character varying DEFAULT NULL::character varying, _last_name character varying DEFAULT NULL::character varying, _email character varying DEFAULT NULL::character varying, _picture character varying DEFAULT NULL::character varying, _date_created timestamp without time zone DEFAULT NULL::timestamp without time zone) RETURNS TABLE(id uuid, identifier character varying, first_name character varying, last_name character varying, email character varying, picture character varying, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN
    RETURN QUERY
    SELECT i.id, i.identifier, i.first_name, i.last_name, i.email, i.picture, i.date_created
    FROM "user".identifier as i
    WHERE (_id IS NULL OR i.id = _id)
    AND (_identifier IS NULL OR i.identifier = _identifier)
    AND (_first_name IS NULL OR i.first_name = _first_name)
    AND (_last_name IS NULL OR i.last_name = _last_name)
    AND (_email IS NULL OR i.email = _email)
    AND (_picture IS NULL OR i.picture = _picture)
    AND (_date_created IS NULL OR i.date_created >= _date_created);
END
$$;


ALTER FUNCTION "user".identifier_find(_id uuid, _identifier character varying, _first_name character varying, _last_name character varying, _email character varying, _picture character varying, _date_created timestamp without time zone) OWNER TO admin;

--
-- TOC entry 285 (class 1255 OID 16474)
-- Name: identifier_remove(uuid); Type: FUNCTION; Schema: user; Owner: admin
--

CREATE FUNCTION "user".identifier_remove(_id uuid) RETURNS TABLE(id uuid, identifier character varying, first_name character varying, last_name character varying, email character varying, picture character varying, date_created timestamp without time zone)
    LANGUAGE plpgsql
    AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM "user".identifier as b
        WHERE b.id = _id
        FOR UPDATE
    ) THEN
        RAISE EXCEPTION 'invalid_user'
            USING HINT = 'Please check the provided user id parameter';
    END IF;

    RETURN QUERY
    DELETE FROM "user".identifier as a WHERE a.id = _id
    RETURNING *;

END
$$;


ALTER FUNCTION "user".identifier_remove(_id uuid) OWNER TO admin;

--
-- TOC entry 297 (class 1255 OID 16475)
-- Name: identifier_store(uuid, character varying, character varying, character varying, character varying, character varying); Type: FUNCTION; Schema: user; Owner: admin
--

CREATE FUNCTION "user".identifier_store(_id uuid, _identifier character varying, _first_name character varying, _last_name character varying, _email character varying, _picture character varying) RETURNS uuid
    LANGUAGE plpgsql
    AS $$
DECLARE
	__id uuid;
BEGIN
    IF EXISTS(SELECT 1 FROM "user".identifier WHERE identifier = _identifier AND id = _id) THEN
        UPDATE "user".identifier
        SET first_name = _first_name,
            last_name = _last_name,
            email = _email,
            picture = _picture
        WHERE identifier = _identifier
        RETURNING identifier.id INTO __id;
    ELSE
        INSERT INTO "user".identifier (id, identifier, first_name, last_name, email, picture)
        VALUES (COALESCE(_id, uuid_generate_v4()), _identifier, _first_name, _last_name, _email, _picture)
        RETURNING identifier.id INTO __id;
    END IF;
    RETURN __id;
END
$$;


ALTER FUNCTION "user".identifier_store(_id uuid, _identifier character varying, _first_name character varying, _last_name character varying, _email character varying, _picture character varying) OWNER TO admin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 230 (class 1259 OID 16715)
-- Name: identifier; Type: TABLE; Schema: booking; Owner: admin
--

CREATE TABLE booking.identifier (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    resource_type resource.type NOT NULL,
    resource_preference_id uuid,
    resource_id uuid,
    start timestamp without time zone NOT NULL,
    "end" timestamp without time zone NOT NULL,
    booked boolean DEFAULT false,
    date_created timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text)
);


ALTER TABLE booking.identifier OWNER TO admin;

--
-- TOC entry 224 (class 1259 OID 16605)
-- Name: role; Type: TABLE; Schema: permission; Owner: admin
--

CREATE TABLE permission.role (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    role_id uuid NOT NULL,
    permission_type permission.type NOT NULL,
    permission_category permission.category NOT NULL,
    permission_tenant permission.tenant NOT NULL,
    permission_tenant_id uuid,
    date_added timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text)
);


ALTER TABLE permission.role OWNER TO admin;

--
-- TOC entry 225 (class 1259 OID 16618)
-- Name: user; Type: TABLE; Schema: permission; Owner: admin
--

CREATE TABLE permission."user" (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    permission_type permission.type NOT NULL,
    permission_category permission.category NOT NULL,
    permission_tenant permission.tenant NOT NULL,
    permission_tenant_id uuid,
    date_added timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text)
);


ALTER TABLE permission."user" OWNER TO admin;

--
-- TOC entry 226 (class 1259 OID 16645)
-- Name: building; Type: TABLE; Schema: resource; Owner: admin
--

CREATE TABLE resource.building (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(256),
    location character varying(256),
    dimension character varying(256) DEFAULT '5x5'::character varying NOT NULL
);


ALTER TABLE resource.building OWNER TO admin;

--
-- TOC entry 229 (class 1259 OID 16683)
-- Name: identifier; Type: TABLE; Schema: resource; Owner: admin
--

CREATE TABLE resource.identifier (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    room_id uuid,
    name character varying(256),
    location character varying(256),
    role_id uuid,
    resource_type resource.type NOT NULL,
    date_created timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text)
);


ALTER TABLE resource.identifier OWNER TO admin;

--
-- TOC entry 227 (class 1259 OID 16654)
-- Name: room; Type: TABLE; Schema: resource; Owner: admin
--

CREATE TABLE resource.room (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    building_id uuid,
    name character varying(256),
    location character varying(256),
    dimension character varying(256) DEFAULT '5x5'::character varying NOT NULL
);


ALTER TABLE resource.room OWNER TO admin;

--
-- TOC entry 228 (class 1259 OID 16668)
-- Name: room_association; Type: TABLE; Schema: resource; Owner: admin
--

CREATE TABLE resource.room_association (
    room_id uuid NOT NULL,
    room_id_association uuid NOT NULL
);


ALTER TABLE resource.room_association OWNER TO admin;

--
-- TOC entry 222 (class 1259 OID 16527)
-- Name: identifier; Type: TABLE; Schema: role; Owner: admin
--

CREATE TABLE role.identifier (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    role_name character varying(256) NOT NULL,
    date_added timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text),
    CONSTRAINT identifier_role_name_check CHECK (((role_name)::text <> ''::text))
);


ALTER TABLE role.identifier OWNER TO admin;

--
-- TOC entry 223 (class 1259 OID 16537)
-- Name: user; Type: TABLE; Schema: role; Owner: admin
--

CREATE TABLE role."user" (
    role_id uuid NOT NULL,
    user_id uuid NOT NULL,
    date_added timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text)
);


ALTER TABLE role."user" OWNER TO admin;

--
-- TOC entry 221 (class 1259 OID 16502)
-- Name: association; Type: TABLE; Schema: team; Owner: admin
--

CREATE TABLE team.association (
    team_id uuid NOT NULL,
    team_id_association uuid NOT NULL
);


ALTER TABLE team.association OWNER TO admin;

--
-- TOC entry 219 (class 1259 OID 16477)
-- Name: identifier; Type: TABLE; Schema: team; Owner: admin
--

CREATE TABLE team.identifier (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(256),
    description character varying(256),
    capacity integer,
    picture character varying(256),
    date_created timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text)
);


ALTER TABLE team.identifier OWNER TO admin;

--
-- TOC entry 220 (class 1259 OID 16486)
-- Name: user; Type: TABLE; Schema: team; Owner: admin
--

CREATE TABLE team."user" (
    team_id uuid NOT NULL,
    user_id uuid NOT NULL,
    date_added timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text)
);


ALTER TABLE team."user" OWNER TO admin;

--
-- TOC entry 218 (class 1259 OID 16455)
-- Name: credential; Type: TABLE; Schema: user; Owner: admin
--

CREATE TABLE "user".credential (
    id character varying(256) NOT NULL,
    secret character varying(256),
    identifier character varying(256) NOT NULL,
    type "user".credential_type GENERATED ALWAYS AS (
CASE
    WHEN ((secret IS NULL) AND ((id)::text !~~* 'local.%'::text)) THEN 'federated'::"user".credential_type
    ELSE 'local'::"user".credential_type
END) STORED,
    active boolean NOT NULL,
    failed_attempts integer DEFAULT 0 NOT NULL,
    last_accessed timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text)
);


ALTER TABLE "user".credential OWNER TO admin;

--
-- TOC entry 217 (class 1259 OID 16434)
-- Name: identifier; Type: TABLE; Schema: user; Owner: admin
--

CREATE TABLE "user".identifier (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    identifier character varying(256) NOT NULL,
    first_name character varying(256),
    last_name character varying(256),
    email character varying(256),
    picture character varying(256),
    date_created timestamp without time zone DEFAULT (now() AT TIME ZONE 'uct'::text),
    CONSTRAINT identifier_email_check CHECK (((email)::text <> ''::text)),
    CONSTRAINT identifier_first_name_check CHECK (((first_name)::text <> ''::text)),
    CONSTRAINT identifier_last_name_check CHECK (((last_name)::text <> ''::text)),
    CONSTRAINT identifier_picture_check CHECK (((picture)::text <> ''::text))
);


ALTER TABLE "user".identifier OWNER TO admin;

--
-- TOC entry 3561 (class 0 OID 16715)
-- Dependencies: 230
-- Data for Name: identifier; Type: TABLE DATA; Schema: booking; Owner: admin
--

INSERT INTO booking.identifier VALUES ('33333333-1111-4a06-9983-8b374586e459', '11111111-1111-4a06-9983-8b374586e459', 'DESK', '22222222-dc08-4a06-9983-8b374586e459', NULL, '2022-05-09 09:54:16.865562', '2022-05-09 13:54:16.865562', false, '2022-05-24 20:14:40.224673');
INSERT INTO booking.identifier VALUES ('33333333-2222-4a06-9983-8b374586e459', '11111111-2222-4a06-9983-8b374586e459', 'DESK', '22222222-dc08-4a06-9983-8b374586e459', NULL, '2022-05-09 09:54:16.865562', '2022-05-09 13:54:16.865562', false, '2022-05-24 20:14:40.228339');


--
-- TOC entry 3555 (class 0 OID 16605)
-- Dependencies: 224
-- Data for Name: role; Type: TABLE DATA; Schema: permission; Owner: admin
--

INSERT INTO permission.role VALUES ('d10a1637-7cf8-4d9b-ae9f-5ca9e96b51da', '45454545-1111-4a06-9983-8b374586e459', 'VIEW', 'BOOKING', 'ROLE', '45454545-1111-4a06-9983-8b374586e459', '2022-05-24 20:14:40.243897');
INSERT INTO permission.role VALUES ('21af45a7-87fe-43b9-a191-0dcd80e88cd4', '45454545-1111-4a06-9983-8b374586e459', 'VIEW', 'BOOKING', 'ROLE', '45454545-2222-4a06-9983-8b374586e459', '2022-05-24 20:14:40.24871');


--
-- TOC entry 3556 (class 0 OID 16618)
-- Dependencies: 225
-- Data for Name: user; Type: TABLE DATA; Schema: permission; Owner: admin
--

INSERT INTO permission."user" VALUES ('1401a9da-23b7-4a9c-bed3-22343a1d4642', '00000000-0000-0000-0000-000000000000', 'VIEW', 'BOOKING', 'USER', NULL, '2022-05-24 20:14:40.144529');
INSERT INTO permission."user" VALUES ('1dee94c2-d2ee-4b93-904e-8a5100eb2bcd', '00000000-0000-0000-0000-000000000000', 'CREATE', 'BOOKING', 'USER', NULL, '2022-05-24 20:14:40.148413');
INSERT INTO permission."user" VALUES ('edebf417-8528-4a5c-a8ef-3be9c1b2c73d', '00000000-0000-0000-0000-000000000000', 'DELETE', 'BOOKING', 'USER', NULL, '2022-05-24 20:14:40.149619');
INSERT INTO permission."user" VALUES ('2fc9d3d1-7627-43cd-985a-d5ecf18358f2', '00000000-0000-0000-0000-000000000000', 'CREATE', 'PERMISSION', 'USER', NULL, '2022-05-24 20:14:40.150944');
INSERT INTO permission."user" VALUES ('5e6a0b23-b504-4a99-a108-8dcbd1ed2ff5', '00000000-0000-0000-0000-000000000000', 'VIEW', 'PERMISSION', 'USER', NULL, '2022-05-24 20:14:40.152115');
INSERT INTO permission."user" VALUES ('5df549c5-00b8-45c3-a354-c74ae68801e1', '00000000-0000-0000-0000-000000000000', 'DELETE', 'PERMISSION', 'USER', NULL, '2022-05-24 20:14:40.153523');
INSERT INTO permission."user" VALUES ('2cdafd7e-1e6a-4228-a30a-9e7234bec165', '00000000-0000-0000-0000-000000000000', 'CREATE', 'PERMISSION', 'ROLE', NULL, '2022-05-24 20:14:40.154833');
INSERT INTO permission."user" VALUES ('aec5443c-2da2-4674-b71d-1169ef59ae07', '00000000-0000-0000-0000-000000000000', 'VIEW', 'PERMISSION', 'ROLE', NULL, '2022-05-24 20:14:40.155973');
INSERT INTO permission."user" VALUES ('a423e5c5-8386-4035-acae-e3eedcf03041', '00000000-0000-0000-0000-000000000000', 'DELETE', 'PERMISSION', 'ROLE', NULL, '2022-05-24 20:14:40.157287');
INSERT INTO permission."user" VALUES ('19b472ef-02b5-4dcc-bfec-5b6e651dc229', '00000000-0000-0000-0000-000000000000', 'CREATE', 'RESOURCE', 'IDENTIFIER', NULL, '2022-05-24 20:14:40.158442');
INSERT INTO permission."user" VALUES ('f6614443-3dce-44ae-80e4-cd353fd02dd6', '00000000-0000-0000-0000-000000000000', 'VIEW', 'RESOURCE', 'IDENTIFIER', NULL, '2022-05-24 20:14:40.159728');
INSERT INTO permission."user" VALUES ('89fe5276-a475-4f4b-8a6f-d7f933b7b63a', '00000000-0000-0000-0000-000000000000', 'DELETE', 'RESOURCE', 'IDENTIFIER', NULL, '2022-05-24 20:14:40.16085');
INSERT INTO permission."user" VALUES ('2f969b74-bb6b-4df5-aee2-964bffe1bf43', '00000000-0000-0000-0000-000000000000', 'CREATE', 'RESOURCE', 'ROOM', NULL, '2022-05-24 20:14:40.161955');
INSERT INTO permission."user" VALUES ('10889b55-9561-49a8-a552-caff8b102730', '00000000-0000-0000-0000-000000000000', 'VIEW', 'RESOURCE', 'ROOM', NULL, '2022-05-24 20:14:40.163175');
INSERT INTO permission."user" VALUES ('beec8a55-10ef-4aec-bffa-870118fcfc7b', '00000000-0000-0000-0000-000000000000', 'DELETE', 'RESOURCE', 'ROOM', NULL, '2022-05-24 20:14:40.164253');
INSERT INTO permission."user" VALUES ('b8c19f27-4c2f-4c76-94c8-f092c110f331', '00000000-0000-0000-0000-000000000000', 'CREATE', 'RESOURCE', 'ROOMASSOCIATION', NULL, '2022-05-24 20:14:40.165481');
INSERT INTO permission."user" VALUES ('3e0fa228-2731-4c35-a2c7-1d6d84781b28', '00000000-0000-0000-0000-000000000000', 'VIEW', 'RESOURCE', 'ROOMASSOCIATION', NULL, '2022-05-24 20:14:40.166616');
INSERT INTO permission."user" VALUES ('cf6077cb-7f3f-4ae7-97ee-0501331635fa', '00000000-0000-0000-0000-000000000000', 'DELETE', 'RESOURCE', 'ROOMASSOCIATION', NULL, '2022-05-24 20:14:40.167901');
INSERT INTO permission."user" VALUES ('4362937f-f569-498a-b1de-8866c3027697', '00000000-0000-0000-0000-000000000000', 'CREATE', 'RESOURCE', 'BUILDING', NULL, '2022-05-24 20:14:40.169258');
INSERT INTO permission."user" VALUES ('31c11a56-53b2-45ba-a3e3-d895cb194838', '00000000-0000-0000-0000-000000000000', 'VIEW', 'RESOURCE', 'BUILDING', NULL, '2022-05-24 20:14:40.170465');
INSERT INTO permission."user" VALUES ('7660c4a1-bc53-4f79-8a56-c9d686a424df', '00000000-0000-0000-0000-000000000000', 'DELETE', 'RESOURCE', 'BUILDING', NULL, '2022-05-24 20:14:40.171715');
INSERT INTO permission."user" VALUES ('af86090d-739f-41f3-998d-18bf2b2397bc', '00000000-0000-0000-0000-000000000000', 'CREATE', 'ROLE', 'IDENTIFIER', NULL, '2022-05-24 20:14:40.172953');
INSERT INTO permission."user" VALUES ('0601e7a6-5f5f-48fe-a692-095504b6ef06', '00000000-0000-0000-0000-000000000000', 'VIEW', 'ROLE', 'IDENTIFIER', NULL, '2022-05-24 20:14:40.174275');
INSERT INTO permission."user" VALUES ('4dceaba6-9671-49de-904a-ca54ca485a2d', '00000000-0000-0000-0000-000000000000', 'DELETE', 'ROLE', 'IDENTIFIER', NULL, '2022-05-24 20:14:40.175525');
INSERT INTO permission."user" VALUES ('5bf3384f-9335-4f9f-84ab-ec0492aa568a', '00000000-0000-0000-0000-000000000000', 'CREATE', 'ROLE', 'USER', NULL, '2022-05-24 20:14:40.176824');
INSERT INTO permission."user" VALUES ('e97b44de-a42f-4a00-a86b-b2ad4eb7e7e2', '00000000-0000-0000-0000-000000000000', 'VIEW', 'ROLE', 'USER', NULL, '2022-05-24 20:14:40.178077');
INSERT INTO permission."user" VALUES ('efc79932-f9b0-4839-9372-938b9717cf35', '00000000-0000-0000-0000-000000000000', 'DELETE', 'ROLE', 'USER', NULL, '2022-05-24 20:14:40.179168');
INSERT INTO permission."user" VALUES ('9f0d6e73-6d19-49c0-8111-6a79531c5932', '00000000-0000-0000-0000-000000000000', 'CREATE', 'TEAM', 'IDENTIFIER', NULL, '2022-05-24 20:14:40.1807');
INSERT INTO permission."user" VALUES ('52fe43ef-69a7-43d5-a31c-86af944dd97f', '00000000-0000-0000-0000-000000000000', 'VIEW', 'TEAM', 'IDENTIFIER', NULL, '2022-05-24 20:14:40.181869');
INSERT INTO permission."user" VALUES ('d298350a-dc76-445b-a793-423c2e2712ec', '00000000-0000-0000-0000-000000000000', 'CREATE', 'TEAM', 'USER', NULL, '2022-05-24 20:14:40.183302');
INSERT INTO permission."user" VALUES ('1a824708-bf78-414e-b56d-bfc527430b44', '00000000-0000-0000-0000-000000000000', 'VIEW', 'TEAM', 'USER', NULL, '2022-05-24 20:14:40.184497');
INSERT INTO permission."user" VALUES ('7db817e6-8292-451b-8291-18c69d329755', '00000000-0000-0000-0000-000000000000', 'CREATE', 'TEAM', 'ASSOCIATION', NULL, '2022-05-24 20:14:40.186043');
INSERT INTO permission."user" VALUES ('9fa92c86-e9f9-4481-bce4-9b4421521e8d', '00000000-0000-0000-0000-000000000000', 'VIEW', 'TEAM', 'ASSOCIATION', NULL, '2022-05-24 20:14:40.187213');
INSERT INTO permission."user" VALUES ('49f311f9-9481-4ad4-9d4c-1c855885c30f', '11111111-3333-4a06-9983-8b374586e459', 'VIEW', 'BOOKING', 'USER', NULL, '2022-05-24 20:14:40.240414');
INSERT INTO permission."user" VALUES ('4dd4bb88-a248-4ffd-a5b1-5693513dd6eb', '11111111-2222-4a06-9983-8b374586e459', 'VIEW', 'BOOKING', 'USER', '11111111-2222-4a06-9983-8b374586e459', '2022-05-24 20:14:40.241742');


--
-- TOC entry 3557 (class 0 OID 16645)
-- Dependencies: 226
-- Data for Name: building; Type: TABLE DATA; Schema: resource; Owner: admin
--

INSERT INTO resource.building VALUES ('98989898-dc08-4a06-9983-8b374586e459', 'aName', 'ALocation', '5x5');


--
-- TOC entry 3560 (class 0 OID 16683)
-- Dependencies: 229
-- Data for Name: identifier; Type: TABLE DATA; Schema: resource; Owner: admin
--

INSERT INTO resource.identifier VALUES ('22222222-dc08-4a06-9983-8b374586e459', '14141414-dc08-4a06-9983-8b374586e459', 'name', 'ALocation', NULL, 'DESK', '2022-05-24 20:14:40.220707');


--
-- TOC entry 3558 (class 0 OID 16654)
-- Dependencies: 227
-- Data for Name: room; Type: TABLE DATA; Schema: resource; Owner: admin
--

INSERT INTO resource.room VALUES ('14141414-dc08-4a06-9983-8b374586e459', '98989898-dc08-4a06-9983-8b374586e459', 'aName', 'ALocation', '5x5');
INSERT INTO resource.room VALUES ('15151515-dc08-4a06-9983-8b374586e459', '98989898-dc08-4a06-9983-8b374586e459', 'aName', 'ALocation', '5x5');


--
-- TOC entry 3559 (class 0 OID 16668)
-- Dependencies: 228
-- Data for Name: room_association; Type: TABLE DATA; Schema: resource; Owner: admin
--

INSERT INTO resource.room_association VALUES ('15151515-dc08-4a06-9983-8b374586e459', '14141414-dc08-4a06-9983-8b374586e459');


--
-- TOC entry 3553 (class 0 OID 16527)
-- Dependencies: 222
-- Data for Name: identifier; Type: TABLE DATA; Schema: role; Owner: admin
--

INSERT INTO role.identifier VALUES ('45454545-1111-4a06-9983-8b374586e459', 'aRole', '2022-05-24 20:14:40.229881');
INSERT INTO role.identifier VALUES ('45454545-2222-4a06-9983-8b374586e459', 'anotherRole', '2022-05-24 20:14:40.23427');


--
-- TOC entry 3554 (class 0 OID 16537)
-- Dependencies: 223
-- Data for Name: user; Type: TABLE DATA; Schema: role; Owner: admin
--

INSERT INTO role."user" VALUES ('45454545-1111-4a06-9983-8b374586e459', '11111111-1111-4a06-9983-8b374586e459', '2022-05-24 20:14:40.235491');
INSERT INTO role."user" VALUES ('45454545-2222-4a06-9983-8b374586e459', '11111111-2222-4a06-9983-8b374586e459', '2022-05-24 20:14:40.239051');


--
-- TOC entry 3552 (class 0 OID 16502)
-- Dependencies: 221
-- Data for Name: association; Type: TABLE DATA; Schema: team; Owner: admin
--

INSERT INTO team.association VALUES ('12121212-dc08-4a06-9983-8b374586e459', '47474747-dc08-4a06-9983-8b374586e459');


--
-- TOC entry 3550 (class 0 OID 16477)
-- Dependencies: 219
-- Data for Name: identifier; Type: TABLE DATA; Schema: team; Owner: admin
--

INSERT INTO team.identifier VALUES ('12121212-dc08-4a06-9983-8b374586e459', 'aTeam', 'a description', 5, 'picture', '2022-05-24 20:14:40.193113');
INSERT INTO team.identifier VALUES ('47474747-dc08-4a06-9983-8b374586e459', 'anotherTeam', 'a description', 5, 'picture', '2022-05-24 20:14:40.197778');


--
-- TOC entry 3551 (class 0 OID 16486)
-- Dependencies: 220
-- Data for Name: user; Type: TABLE DATA; Schema: team; Owner: admin
--

INSERT INTO team."user" VALUES ('12121212-dc08-4a06-9983-8b374586e459', '11111111-1111-4a06-9983-8b374586e459', '2022-05-24 20:14:40.203489');


--
-- TOC entry 3549 (class 0 OID 16455)
-- Dependencies: 218
-- Data for Name: credential; Type: TABLE DATA; Schema: user; Owner: admin
--



--
-- TOC entry 3548 (class 0 OID 16434)
-- Dependencies: 217
-- Data for Name: identifier; Type: TABLE DATA; Schema: user; Owner: admin
--

INSERT INTO "user".identifier VALUES ('00000000-0000-0000-0000-000000000000', 'admin@example.com', 'Admin', 'Admin', 'admin@example.com', '/picture', '2022-05-24 20:14:40.132478');
INSERT INTO "user".identifier VALUES ('11111111-1111-4a06-9983-8b374586e459', 'email@example.com', 'Test', 'Tester', 'email@example.com', '/picture', '2022-05-24 20:14:40.188771');
INSERT INTO "user".identifier VALUES ('11111111-2222-4a06-9983-8b374586e459', 'anemail@example.com', 'Test', 'Tester', 'anemail@example.com', '/picture', '2022-05-24 20:14:40.190209');
INSERT INTO "user".identifier VALUES ('11111111-3333-4a06-9983-8b374586e459', 'anotheremail@example.com', 'Test', 'Tester', 'anotheremail@example.com', '/picture', '2022-05-24 20:14:40.19168');


--
-- TOC entry 3391 (class 2606 OID 16722)
-- Name: identifier identifier_pkey; Type: CONSTRAINT; Schema: booking; Owner: admin
--

ALTER TABLE ONLY booking.identifier
    ADD CONSTRAINT identifier_pkey PRIMARY KEY (id);


--
-- TOC entry 3378 (class 2606 OID 16611)
-- Name: role role_pkey; Type: CONSTRAINT; Schema: permission; Owner: admin
--

ALTER TABLE ONLY permission.role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);


--
-- TOC entry 3381 (class 2606 OID 16624)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: permission; Owner: admin
--

ALTER TABLE ONLY permission."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 3383 (class 2606 OID 16653)
-- Name: building building_pkey; Type: CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.building
    ADD CONSTRAINT building_pkey PRIMARY KEY (id);


--
-- TOC entry 3389 (class 2606 OID 16691)
-- Name: identifier identifier_pkey; Type: CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.identifier
    ADD CONSTRAINT identifier_pkey PRIMARY KEY (id);


--
-- TOC entry 3387 (class 2606 OID 16672)
-- Name: room_association room_association_pkey; Type: CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.room_association
    ADD CONSTRAINT room_association_pkey PRIMARY KEY (room_id, room_id_association);


--
-- TOC entry 3385 (class 2606 OID 16662)
-- Name: room room_pkey; Type: CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.room
    ADD CONSTRAINT room_pkey PRIMARY KEY (id);


--
-- TOC entry 3371 (class 2606 OID 16534)
-- Name: identifier identifier_pkey; Type: CONSTRAINT; Schema: role; Owner: admin
--

ALTER TABLE ONLY role.identifier
    ADD CONSTRAINT identifier_pkey PRIMARY KEY (id);


--
-- TOC entry 3373 (class 2606 OID 16536)
-- Name: identifier identifier_role_name_key; Type: CONSTRAINT; Schema: role; Owner: admin
--

ALTER TABLE ONLY role.identifier
    ADD CONSTRAINT identifier_role_name_key UNIQUE (role_name);


--
-- TOC entry 3375 (class 2606 OID 16542)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: role; Owner: admin
--

ALTER TABLE ONLY role."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (role_id, user_id);


--
-- TOC entry 3369 (class 2606 OID 16506)
-- Name: association association_pkey; Type: CONSTRAINT; Schema: team; Owner: admin
--

ALTER TABLE ONLY team.association
    ADD CONSTRAINT association_pkey PRIMARY KEY (team_id, team_id_association);


--
-- TOC entry 3365 (class 2606 OID 16485)
-- Name: identifier identifier_pkey; Type: CONSTRAINT; Schema: team; Owner: admin
--

ALTER TABLE ONLY team.identifier
    ADD CONSTRAINT identifier_pkey PRIMARY KEY (id);


--
-- TOC entry 3367 (class 2606 OID 16491)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: team; Owner: admin
--

ALTER TABLE ONLY team."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (team_id, user_id);


--
-- TOC entry 3363 (class 2606 OID 16464)
-- Name: credential credential_pkey; Type: CONSTRAINT; Schema: user; Owner: admin
--

ALTER TABLE ONLY "user".credential
    ADD CONSTRAINT credential_pkey PRIMARY KEY (id);


--
-- TOC entry 3359 (class 2606 OID 16448)
-- Name: identifier identifier_identifier_key; Type: CONSTRAINT; Schema: user; Owner: admin
--

ALTER TABLE ONLY "user".identifier
    ADD CONSTRAINT identifier_identifier_key UNIQUE (identifier);


--
-- TOC entry 3361 (class 2606 OID 16446)
-- Name: identifier identifier_pkey; Type: CONSTRAINT; Schema: user; Owner: admin
--

ALTER TABLE ONLY "user".identifier
    ADD CONSTRAINT identifier_pkey PRIMARY KEY (id);


--
-- TOC entry 3376 (class 1259 OID 16617)
-- Name: permission_role_1; Type: INDEX; Schema: permission; Owner: admin
--

CREATE UNIQUE INDEX permission_role_1 ON permission.role USING btree (role_id, permission_type, permission_category, permission_tenant, permission_tenant_id) WHERE (permission_tenant_id IS NOT NULL);


--
-- TOC entry 3379 (class 1259 OID 16630)
-- Name: permission_user_1; Type: INDEX; Schema: permission; Owner: admin
--

CREATE UNIQUE INDEX permission_user_1 ON permission."user" USING btree (user_id, permission_type, permission_category, permission_tenant, permission_tenant_id) WHERE (permission_tenant_id IS NOT NULL);


--
-- TOC entry 3408 (class 2606 OID 16733)
-- Name: identifier identifier_resource_id_fkey; Type: FK CONSTRAINT; Schema: booking; Owner: admin
--

ALTER TABLE ONLY booking.identifier
    ADD CONSTRAINT identifier_resource_id_fkey FOREIGN KEY (resource_id) REFERENCES resource.identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3407 (class 2606 OID 16728)
-- Name: identifier identifier_resource_preference_id_fkey; Type: FK CONSTRAINT; Schema: booking; Owner: admin
--

ALTER TABLE ONLY booking.identifier
    ADD CONSTRAINT identifier_resource_preference_id_fkey FOREIGN KEY (resource_preference_id) REFERENCES resource.identifier(id) ON DELETE SET NULL;


--
-- TOC entry 3406 (class 2606 OID 16723)
-- Name: identifier identifier_user_id_fkey; Type: FK CONSTRAINT; Schema: booking; Owner: admin
--

ALTER TABLE ONLY booking.identifier
    ADD CONSTRAINT identifier_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user".identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3399 (class 2606 OID 16612)
-- Name: role role_role_id_fkey; Type: FK CONSTRAINT; Schema: permission; Owner: admin
--

ALTER TABLE ONLY permission.role
    ADD CONSTRAINT role_role_id_fkey FOREIGN KEY (role_id) REFERENCES role.identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3400 (class 2606 OID 16625)
-- Name: user user_user_id_fkey; Type: FK CONSTRAINT; Schema: permission; Owner: admin
--

ALTER TABLE ONLY permission."user"
    ADD CONSTRAINT user_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user".identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3405 (class 2606 OID 16697)
-- Name: identifier identifier_role_id_fkey; Type: FK CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.identifier
    ADD CONSTRAINT identifier_role_id_fkey FOREIGN KEY (role_id) REFERENCES role.identifier(id) ON DELETE SET NULL;


--
-- TOC entry 3404 (class 2606 OID 16692)
-- Name: identifier identifier_room_id_fkey; Type: FK CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.identifier
    ADD CONSTRAINT identifier_room_id_fkey FOREIGN KEY (room_id) REFERENCES resource.room(id) ON DELETE CASCADE;


--
-- TOC entry 3403 (class 2606 OID 16678)
-- Name: room_association room_association_room_id_association_fkey; Type: FK CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.room_association
    ADD CONSTRAINT room_association_room_id_association_fkey FOREIGN KEY (room_id_association) REFERENCES resource.room(id) ON DELETE CASCADE;


--
-- TOC entry 3402 (class 2606 OID 16673)
-- Name: room_association room_association_room_id_fkey; Type: FK CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.room_association
    ADD CONSTRAINT room_association_room_id_fkey FOREIGN KEY (room_id) REFERENCES resource.room(id) ON DELETE CASCADE;


--
-- TOC entry 3401 (class 2606 OID 16663)
-- Name: room room_building_id_fkey; Type: FK CONSTRAINT; Schema: resource; Owner: admin
--

ALTER TABLE ONLY resource.room
    ADD CONSTRAINT room_building_id_fkey FOREIGN KEY (building_id) REFERENCES resource.building(id) ON DELETE CASCADE;


--
-- TOC entry 3397 (class 2606 OID 16543)
-- Name: user user_role_id_fkey; Type: FK CONSTRAINT; Schema: role; Owner: admin
--

ALTER TABLE ONLY role."user"
    ADD CONSTRAINT user_role_id_fkey FOREIGN KEY (role_id) REFERENCES role.identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3398 (class 2606 OID 16548)
-- Name: user user_user_id_fkey; Type: FK CONSTRAINT; Schema: role; Owner: admin
--

ALTER TABLE ONLY role."user"
    ADD CONSTRAINT user_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user".identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3396 (class 2606 OID 16512)
-- Name: association association_team_id_association_fkey; Type: FK CONSTRAINT; Schema: team; Owner: admin
--

ALTER TABLE ONLY team.association
    ADD CONSTRAINT association_team_id_association_fkey FOREIGN KEY (team_id_association) REFERENCES team.identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3395 (class 2606 OID 16507)
-- Name: association association_team_id_fkey; Type: FK CONSTRAINT; Schema: team; Owner: admin
--

ALTER TABLE ONLY team.association
    ADD CONSTRAINT association_team_id_fkey FOREIGN KEY (team_id) REFERENCES team.identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3393 (class 2606 OID 16492)
-- Name: user user_team_id_fkey; Type: FK CONSTRAINT; Schema: team; Owner: admin
--

ALTER TABLE ONLY team."user"
    ADD CONSTRAINT user_team_id_fkey FOREIGN KEY (team_id) REFERENCES team.identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3394 (class 2606 OID 16497)
-- Name: user user_user_id_fkey; Type: FK CONSTRAINT; Schema: team; Owner: admin
--

ALTER TABLE ONLY team."user"
    ADD CONSTRAINT user_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user".identifier(id) ON DELETE CASCADE;


--
-- TOC entry 3392 (class 2606 OID 16465)
-- Name: credential credential_identifier_fkey; Type: FK CONSTRAINT; Schema: user; Owner: admin
--

ALTER TABLE ONLY "user".credential
    ADD CONSTRAINT credential_identifier_fkey FOREIGN KEY (identifier) REFERENCES "user".identifier(identifier) ON DELETE CASCADE;


-- Completed on 2022-05-25 02:24:21

--
-- PostgreSQL database dump complete
--

