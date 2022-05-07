-----------------
-- User

SELECT "user".identifier_store(
	'email@example.com', 
	'Test', 
	'Tester', 
	'email@example.com', 
	'/picture'
);

-----------------
-- Resource

SELECT resource.identifier_create(
	'aLocation', 
	null, 
	'DESK'::resource.type
);
