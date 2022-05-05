CREATE SCHEMA IF NOT EXISTS team;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS team.identifier (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR(256) CHECK(name <> ''),
    description VARCHAR(256) CHECK(description <> ''),
    capacity INT,
    picture VARCHAR(256) CHECK(picture <> ''),
    date_created TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);