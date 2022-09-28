CREATE SCHEMA IF NOT EXISTS booking;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- TRIGGER USED TO LOWER CHANCES OF DOUBLE BOOKINGS BEING MADE

DROP TRIGGER IF EXISTS checkForConflictingBookings ON "booking".identifier;

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


CREATE OR REPLACE FUNCTION booking.timeInInterval(_check timestamp, _start timestamp, _end timestamp)
RETURNS boolean AS
$$
  BEGIN
    RETURN (_check >= _start AND _check <= _end);
  END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION booking.conflictingBooking(_booking booking.identifier)
RETURNS boolean AS
$$
  DECLARE
    numMatches int;
  BEGIN
    SELECT COUNT(*) INTO numMatches FROM booking.identifier AS i
	  WHERE (_booking).user_id=i.user_id
	  AND (booking.timeInInterval((_booking).start, i.start, i.end) OR booking.timeInInterval((_booking).end, i.start, i.end))
	  AND (_booking).resource_type='DESK';
	RETURN numMatches > 1;
  END;
$$ LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION booking.checkForNoConflicts()
RETURNS TRIGGER AS
$$
  BEGIN
    IF booking.conflictingBooking(new) THEN
	  RAISE EXCEPTION 'conflicting_bookings';
	END IF;
	RETURN new;
  END
$$ LANGUAGE plpgsql;



CREATE CONSTRAINT TRIGGER checkForConflictingBookings
AFTER INSERT OR UPDATE ON booking.identifier
DEFERRABLE INITIALLY DEFERRED
FOR EACH ROW EXECUTE PROCEDURE booking.checkForNoConflicts();