CREATE SCHEMA IF NOT EXISTS example;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS example.example (
    id uuid DEFAULT uuid_generate_v4 ()
);