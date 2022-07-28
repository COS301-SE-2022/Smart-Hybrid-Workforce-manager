----------------------------------
-- Admin User

-- Admin User
SELECT "user".identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid,
	'admin@admin.com', 
	'Admin', 
	'Admin', 
	'admin@admin.com', 
	'/picture',
    false,
    'STANDARD',
    0,
    '09:00',
    '17:00',
	'22222222-dc08-4a06-9983-8b374586e459' -- preferred desk
);

SELECT "user".credential_store(
	'local.admin@admin.com',
	'admin',
	'admin@admin.com'
);

------------ Booking Permissions
-- Permission Admin
SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'BOOKING'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'BOOKING'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'BOOKING'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);


SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'BOOKING'::permission.category,
	'MEETINGROOM'::permission.tenant,
	null::uuid -- All meetingroom bookings
);


SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'BOOKING'::permission.category,
	'MEETINGROOM'::permission.tenant,
	null::uuid -- All meetingroom bookings
);


------------ Permission Permissions
SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'PERMISSION'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'PERMISSION'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'PERMISSION'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

------------ Resource Permissions
SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'RESOURCE'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'RESOURCE'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'RESOURCE'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'RESOURCE'::permission.category,
	'ROOM'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'RESOURCE'::permission.category,
	'ROOM'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'RESOURCE'::permission.category,
	'ROOM'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'RESOURCE'::permission.category,
	'ROOMASSOCIATION'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'RESOURCE'::permission.category,
	'ROOMASSOCIATION'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'RESOURCE'::permission.category,
	'ROOMASSOCIATION'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'RESOURCE'::permission.category,
	'BUILDING'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'RESOURCE'::permission.category,
	'BUILDING'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'RESOURCE'::permission.category,
	'BUILDING'::permission.tenant,
	null::uuid -- All users
);

------------ Role Permissions
SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'ROLE'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'ROLE'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'ROLE'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'ROLE'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'ROLE'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'ROLE'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'USER'::permission.category,
	'ROLE'::permission.tenant,
	null::uuid -- John Doe
);

------------ Team Permissions
SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'TEAM'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'TEAM'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'TEAM'::permission.category,
	'IDENTIFIER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'TEAM'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'TEAM'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'USER'::permission.category,
	'TEAM'::permission.tenant,
	null::uuid -- John Doe
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'TEAM'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All users
);


SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'CREATE'::permission.type,
	'TEAM'::permission.category,
	'ASSOCIATION'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'TEAM'::permission.category,
	'ASSOCIATION'::permission.tenant,
	null::uuid -- All users
);

SELECT permission.identifier_store(
	'00000000-0000-0000-0000-000000000000'::uuid, -- Admin
	'USER'::permission.id_type,
	'DELETE'::permission.type,
	'TEAM'::permission.category,
	'ASSOCIATION'::permission.tenant,
	null::uuid -- All users
);
----------------------------------
-- User

-- User 01
SELECT "user".identifier_store(
	'11111111-1111-4a06-9983-8b374586e459'::uuid,
	'john.doe@gmail.com', 
	'John', 
	'Doe', 
	'john.doe@gmail.com', 
	'/johndoe.png',
    false,
    'STANDARD',
    0,
    '09:00',
    '17:00'
);

-- User 02
SELECT "user".identifier_store(
	'11111111-2222-4a06-9983-8b374586e459'::uuid,
	'jane.doe@icloud.com', 
	'Jane', 
	'Doe', 
	'jane.doe@icloud.com', 
	'/janedoe.jpeg',
    false,
    'STANDARD',
    0,
    '09:00',
    '17:00'
);

-- User 03
SELECT "user".identifier_store(
	'11111111-3333-4a06-9983-8b374586e459'::uuid,
	'steve@harrington.com', 
	'Steve', 
	'Harrington', 
	'steve@harrington.com', 
	'/steve.png',
    false,
    'STANDARD',
    0,
    '09:00',
    '17:00'
);

----------------------------------
-- Team

-- Team 01
SELECT team.identifier_store(
	'12121212-dc08-4a06-9983-8b374586e459'::uuid,
	'Android',
	'Android project', 
	5, 
	'picture',
	0,
	'11111111-3333-4a06-9983-8b374586e459'
);

-- Team 02
SELECT team.identifier_store(
	'47474747-dc08-4a06-9983-8b374586e459'::uuid,
	'Aerial Photography',
	'Project on Aerial Photography', 
	5, 
	'picture',
	1,
	'11111111-3333-4a06-9983-8b374586e459'
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
	'11111111-1111-4a06-9983-8b374586e459'::uuid -- User 01
);

----------------------------------
-- Resource

-- Building 01
SELECT resource.building_store(
	'98989898-dc08-4a06-9983-8b374586e459'::uuid,
	'Durban Office',
	'63 South Street Drive',
	'10x5'
);

-- Room 01
SELECT resource.room_store(
	'14141414-dc08-4a06-9983-8b374586e459'::uuid,
	'98989898-dc08-4a06-9983-8b374586e459'::uuid,
	'Office Block B',
	'B4',
	'5x5'
);


-- Room 02
SELECT resource.room_store(
	'15151515-dc08-4a06-9983-8b374586e459'::uuid,
	'98989898-dc08-4a06-9983-8b374586e459'::uuid,
	'Secretary',
	'B9',
	'5x5'
);

-- Room Association 01
SELECT resource.room_association_store(
	'15151515-dc08-4a06-9983-8b374586e459'::uuid, -- Room 02
	'14141414-dc08-4a06-9983-8b374586e459'::uuid -- Room 01

);

-- Resource Desk 01
SELECT resource.identifier_store(
	'22222222-dc08-4a06-9983-8b374586e459'::uuid,
	'14141414-dc08-4a06-9983-8b374586e459'::uuid, -- Room 01
	'115', 
	'10',
    '10',
    '60',
    '55',
    '0', 
	null::uuid, 
	'DESK'::resource.type,
	'{}'
);

----------------------------------
-- Booking

-- Booking 01
SELECT booking.identifier_store(
	'33333333-1111-4a06-9983-8b374586e459'::uuid,
	'11111111-1111-4a06-9983-8b374586e459'::uuid, -- User 01
	'DESK'::resource.type,
	'22222222-dc08-4a06-9983-8b374586e459'::uuid, -- Resource Desk 01
	null::uuid,
	'2022-05-09 09:54:16.865562'::TIMESTAMP,
	'2022-05-09 13:54:16.865562'::TIMESTAMP,
	null,
	null,
	null
);

-- Booking 02
SELECT booking.identifier_store(
	'33333333-2222-4a06-9983-8b374586e459'::uuid,
	'11111111-2222-4a06-9983-8b374586e459'::uuid, -- User 02
	'DESK'::resource.type,
	'22222222-dc08-4a06-9983-8b374586e459'::uuid, -- Resource Desk 01
	null::uuid,
	'2022-05-09 09:54:16.865562'::TIMESTAMP,
	'2022-05-09 13:54:16.865562'::TIMESTAMP,
	null,
	null,
	null
);

----------------------------------
-- Role

-- Role 01
SELECT role.identifier_store('45454545-1111-4a06-9983-8b374586e459'::uuid, 'Executives', '11111111-1111-4a06-9983-8b374586e459');
-- Role 02
SELECT role.identifier_store('45454545-2222-4a06-9983-8b374586e459'::uuid, 'Secretary', '11111111-2222-4a06-9983-8b374586e459');

-- User Role 01
SELECT role.user_store('45454545-1111-4a06-9983-8b374586e459'::uuid ,'11111111-1111-4a06-9983-8b374586e459'::uuid); -- Role 01, User 01

-- User Role 02
SELECT role.user_store('45454545-2222-4a06-9983-8b374586e459'::uuid ,'11111111-2222-4a06-9983-8b374586e459'::uuid); -- Role 02, User 02

----------------------------------
-- Permissions

-- Permission User 01
SELECT permission.identifier_store(
	'11111111-3333-4a06-9983-8b374586e459'::uuid, -- User 03
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'BOOKING'::permission.category,
	'USER'::permission.tenant,
	null::uuid -- All
);

-- Permission User 02
SELECT permission.identifier_store(
	'11111111-2222-4a06-9983-8b374586e459'::uuid, -- User 02
	'USER'::permission.id_type,
	'VIEW'::permission.type,
	'BOOKING'::permission.category,
	'USER'::permission.tenant,
	'11111111-2222-4a06-9983-8b374586e459'::uuid -- User 02
);

-- Permission Role 01
SELECT permission.identifier_store(
	'45454545-1111-4a06-9983-8b374586e459'::uuid, -- Role 01
	'ROLE'::permission.id_type,
	'VIEW'::permission.type,
	'BOOKING'::permission.category,
	'ROLE'::permission.tenant,
	'45454545-1111-4a06-9983-8b374586e459'::uuid -- Role 01
);

-- Permission Role 02
SELECT permission.identifier_store(
	'45454545-1111-4a06-9983-8b374586e459'::uuid, -- Role 01
	'ROLE'::permission.id_type,
	'VIEW'::permission.type,
	'BOOKING'::permission.category,
	'ROLE'::permission.tenant,
	'45454545-2222-4a06-9983-8b374586e459'::uuid -- Role 02
);