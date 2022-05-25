CREATE SCHEMA IF NOT EXISTS "pets";

CREATE TABLE "pets".dog (
    id SERIAL PRIMARY KEY,
    "name" VARCHAR(255),
    breed VARCHAR(255)
);
