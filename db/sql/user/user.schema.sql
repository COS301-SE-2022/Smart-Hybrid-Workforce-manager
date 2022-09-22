CREATE SCHEMA IF NOT EXISTS "user";
CREATE SCHEMA IF NOT EXISTS parking;
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE parking.type AS ENUM ('NONE', 'STANDARD', 'DISABLED');

CREATE TABLE IF NOT EXISTS "user".identifier (
    id uuid DEFAULT uuid_generate_v4(),
    identifier VARCHAR(256) UNIQUE NOT NULL,
    first_name VARCHAR(256) CHECK(first_name <> ''),
    last_name VARCHAR(256) CHECK(last_name <> ''),
    email VARCHAR(256) CHECK(email <> ''),
    picture VARCHAR(256) CHECK(picture <> ''),
    date_created TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
    work_from_home BOOLEAN NOT NULL DEFAULT false,
    parking parking.type NOT NULL DEFAULT 'STANDARD',
    office_days INTEGER NOT NULL DEFAULT 0,
    preferred_start_time TIME WITHOUT TIME ZONE DEFAULT NULL,
    preferred_end_time TIME WITHOUT TIME ZONE DEFAULT NULL,
    preferred_desk uuid DEFAULT NULL,
    building_id uuid REFERENCES resource.building(id) ON DELETE CASCADE DEFAULT NULL,
	
    PRIMARY KEY (id)
);

CREATE TYPE "user".credential_type AS ENUM ('federated', 'local');

CREATE TABLE IF NOT EXISTS "user".credential (
    id VARCHAR(256),
    secret VARCHAR(256),
    identifier VARCHAR(256) NOT NULL REFERENCES "user".identifier(identifier) ON DELETE CASCADE,
    "type"  "user".credential_type generated always as (
        CASE
        WHEN secret IS NULL AND id NOT ILIKE 'local.%' THEN 'federated'::"user".credential_type
        ELSE 'local'::"user".credential_type
        END
    ) stored,
    active BOOLEAN NOT NULL,
    failed_attempts INT NOT NULL DEFAULT(0),
    last_accessed TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);