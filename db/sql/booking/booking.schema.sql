CREATE SCHEMA IF NOT EXISTS booking;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS booking.identifier (
    id uuid DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES "user".identifier(id) ON DELETE CASCADE,
    resource_type resource.type NOT NULL, -- TODO [KP]: ENUM
    resource_preference_id uuid REFERENCES resource.identifier(id) ON DELETE SET NULL,
    start TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    "end" TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    booked BOOLEAN DEFAULT(false),
    date_created TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);