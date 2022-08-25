CREATE SCHEMA IF NOT EXISTS booking;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS booking.identifier (
    id uuid DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES "user".identifier(id) ON DELETE CASCADE,
    resource_type resource.type NOT NULL,
    resource_preference_id uuid REFERENCES resource.identifier(id) ON DELETE SET NULL,
    resource_id uuid REFERENCES resource.identifier(id) ON DELETE CASCADE,
    start TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "end" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    booked BOOLEAN DEFAULT(null),
    automated BOOLEAN DEFAULT(false),
    date_created TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
    dependent uuid REFERENCES booking.identifier(id) ON DELETE CASCADE,
    
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS booking.meeting_room (
    id uuid DEFAULT uuid_generate_v4(),
    booking_id uuid NOT NULL REFERENCES booking.identifier(id) ON DELETE CASCADE,
    team_id uuid DEFAULT NULL REFERENCES team.identifier(id) ON DELETE CASCADE,
    role_id uuid DEFAULT NULL REFERENCES role.identifier(id) ON DELETE CASCADE,
    additional_attendees INT DEFAULT 0,
    desks_attendees BOOLEAN DEFAULT(false),
    desks_additional_attendees BOOLEAN DEFAULT(false),
    
    PRIMARY KEY (id)
);