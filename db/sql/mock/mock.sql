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

-- Team 01
SELECT team.identifier_store(
	'12121212-dc08-4a06-9983-8b374586e459'::uuid,
	'aTeam',
	'a description', 
	5, 
	'picture'
);

-- Team 02
SELECT team.identifier_store(
	'47474747-dc08-4a06-9983-8b374586e459'::uuid,
	'anotherTeam',
	'a description', 
	5, 
	'picture'
);

-- Team Association 01
SELECT team.association_store(
	'12121212-dc08-4a06-9983-8b374586e459'::uuid, -- Team 01
	'47474747-dc08-4a06-9983-8b374586e459'::uuid -- Team 02
);

-- Team User 01
-- Team Association 01
SELECT team.user_store(
	'12121212-dc08-4a06-9983-8b374586e459'::uuid, -- Team 01
	'11111111-dc08-4a06-9983-8b374586e459'::uuid -- User 01
);

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

----------------------------------
-- Role

-- Role 01
SELECT role.identifier_store('45454545-dc08-4a06-9983-8b374586e459'::uuid, 'aRole');

-- User Role 01
SELECT role.user_store('45454545-dc08-4a06-9983-8b374586e459'::uuid ,'11111111-dc08-4a06-9983-8b374586e459'::uuid); -- Role 01, User 01

----------------------------------
-- Permissions

-- Permission User 01
SELECT permission.user_store(
	'11111111-dc08-4a06-9983-8b374586e459'::uuid, -- User 01
	'VIEW'::permission.type,
	'BOOKING'::permission.category,
	'USER'::permission.tenant,
	'11111111-dc08-4a06-9983-8b374586e459'::uuid -- User 01
);

-- Permission Role 01
SELECT permission.role_store(
	'45454545-dc08-4a06-9983-8b374586e459'::uuid, -- Role 01
	'VIEW'::permission.type,
	'BOOKING'::permission.category,
	'ROLE'::permission.tenant,
	'45454545-dc08-4a06-9983-8b374586e459'::uuid -- Role 01
);