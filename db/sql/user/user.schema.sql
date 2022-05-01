CREATE SCHEMA IF NOT EXISTS "user";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "user".identifier (
    identifier VARCHAR(256),
    first_name VARCHAR(256) CHECK(first_name <> ''),
    last_name VARCHAR(256) CHECK(last_name <> ''),
    email VARCHAR(256) CHECK(email <> ''),
    picture VARCHAR(256) CHECK(picture <> ''),
    date_created TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (identifier)
);

CREATE TYPE "user".credential_type AS ENUM ('federated', 'local');

CREATE TABLE IF NOT EXISTS "user".credential (
    id VARCHAR(256),
    secret VARCHAR(256),
    identifier VARCHAR(256) NOT NULL,
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