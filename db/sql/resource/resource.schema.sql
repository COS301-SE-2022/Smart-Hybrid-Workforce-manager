CREATE SCHEMA IF NOT EXISTS resource;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE resource.type AS ENUM ('PARKING', 'DESK', 'MEETINGROOM');

CREATE TABLE IF NOT EXISTS resource.building (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR(256),
    location VARCHAR(256), -- TODO [KP]: CHANGE TO WHAT WE USE
    dimension VARCHAR(256) NOT NULL DEFAULT('5x5'),
	
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS resource.room (
    id uuid DEFAULT uuid_generate_v4(),
    building_id uuid REFERENCES resource.building(id) ON DELETE CASCADE,
    name VARCHAR(256),
    location VARCHAR(256), -- TODO [KP]: CHANGE TO WHAT WE USE
    dimension VARCHAR(256) NOT NULL DEFAULT('5x5'),
	
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS resource.room_association (
    room_id uuid NOT NULL REFERENCES resource.room(id) ON DELETE CASCADE,
    room_id_association uuid NOT NULL REFERENCES resource.room(id) ON DELETE CASCADE,
	
    PRIMARY KEY (room_id, room_id_association)
);

CREATE TABLE IF NOT EXISTS resource.identifier (
    id uuid DEFAULT uuid_generate_v4(),
    room_id uuid REFERENCES resource.room(id) ON DELETE CASCADE,
    name VARCHAR(256),
    location VARCHAR(256), -- TODO [KP]: CHANGE TO WHAT WE USE
    role_id uuid REFERENCES role.identifier(id) ON DELETE SET NULL,
    resource_type resource.type NOT NULL,
    date_created TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
    decorations JSON NOT NULL,
	
    PRIMARY KEY (id)
);