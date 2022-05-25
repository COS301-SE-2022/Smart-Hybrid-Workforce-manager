CREATE TABLE users (
  user_id integer unique,
  name    varchar(40),
  email   varchar(40),
  PRIMARY KEY(user_id)
);

INSERT INTO "users" (user_id, name, email) VALUES 
(123, 'Chris', 'Hadfield');