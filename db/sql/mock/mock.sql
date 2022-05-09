----------------------------------
-- User

-- User 01
SELECT "user".identifier_store(
	'11111111-dc08-4a06-9983-8b374586e459'::uuid,
	'email@example.com', 
	'Test', 
	'Tester', 
	'email@example.com', 
	'/picture'
);

----------------------------------
-- Team


----------------------------------
-- Resource

-- Resource 01
SELECT resource.identifier_store(
	'22222222-dc08-4a06-9983-8b374586e459'::uuid,
	null::uuid,
	'ALocation', 
	null::uuid, 
	'DESK'::resource.type
);

----------------------------------
-- Booking

-- Booking 01
SELECT booking.identifier_store(
	'33333333-dc08-4a06-9983-8b374586e459'::uuid,
	'11111111-dc08-4a06-9983-8b374586e459'::uuid, -- User 01
	'DESK'::resource.type,
	'22222222-dc08-4a06-9983-8b374586e459'::uuid, -- Resource 01
	'2022-05-09 09:54:16.865562'::TIMESTAMP,
	'2022-05-09 13:54:16.865562'::TIMESTAMP
);
