CREATE SCHEMA IF NOT EXISTS team;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS team.identifier (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR(256),
    description VARCHAR(256),
    capacity INT,
    picture VARCHAR(256),
    date_created TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS team.user (
    team_id uuid NOT NULL REFERENCES team.identifier(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES "user".identifier(id) ON DELETE CASCADE,
    date_added TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (team_id, user_id)
);

CREATE TABLE IF NOT EXISTS team.association (
    team_id uuid NOT NULL REFERENCES team.identifier(id) ON DELETE CASCADE,
    team_id_association uuid NOT NULL REFERENCES team.identifier(id) ON DELETE CASCADE,
	
    PRIMARY KEY (team_id, team_id_association)
);