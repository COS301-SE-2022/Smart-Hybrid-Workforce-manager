CREATE SCHEMA IF NOT EXISTS role;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS role.identifier (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR(256) UNIQUE NOT NULL CHECK(name <> ''),
    color VARCHAR(256),
    lead_id uuid REFERENCES "user".identifier(id) ON DELETE SET NULL,
    date_added TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS role.user (
    role_id uuid NOT NULL REFERENCES role.identifier(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES "user".identifier(id) ON DELETE CASCADE,
    date_added TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (role_id, user_id)
);