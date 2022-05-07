CREATE SCHEMA IF NOT EXISTS resource;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE resource.type AS ENUM ('PARKING', 'DESK', 'MEETINGROOM');

CREATE TABLE IF NOT EXISTS resource.identifier (
    id uuid DEFAULT uuid_generate_v4(),
    room_id uuid NOT NULL DEFAULT uuid_generate_v4(), -- TODO [KP]: REFERENCE
    location VARCHAR(256), -- TODO [KP]: CHANGE TO WHAT WE USE
    role_id uuid REFERENCES role.identifier(id) ON DELETE SET NULL,
    resource_type resource.type NOT NULL,
    date_created TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);