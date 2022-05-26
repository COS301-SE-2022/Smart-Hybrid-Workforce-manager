CREATE SCHEMA IF NOT EXISTS "pets";

CREATE TABLE pets.person (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    age INT
);