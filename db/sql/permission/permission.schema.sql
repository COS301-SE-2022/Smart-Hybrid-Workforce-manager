CREATE SCHEMA IF NOT EXISTS permission;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE permission.id_type AS ENUM ('USER', 'ROLE', 'TEAM');
CREATE TYPE permission.type AS ENUM ('CREATE', 'DELETE', 'VIEW', 'EDIT');
CREATE TYPE permission.category AS ENUM ('USER', 'BOOKING', 'PERMISSION', 'ROLE', 'TEAM', 'RESOURCE');
CREATE TYPE permission.tenant AS ENUM ('ROLE', 'USER', 'TEAM', 'ASSOCIATION' , 'PERMISSION', 'BUILDING', 'ROOM', 'ROOMASSOCIATION', 'IDENTIFIER', 'NA');

CREATE TABLE IF NOT EXISTS permission.identifier (
    id uuid DEFAULT uuid_generate_v4(),
    permission_id uuid NOT NULL,
    permission_id_type permission.id_type NOT NULL,
    permission_type permission.type NOT NULL,
    permission_category permission.category NOT NULL,
    permission_tenant permission.tenant NOT NULL,
    permission_tenant_id uuid, -- If null then it means all
    date_added TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX permission_identifier_1 ON permission.identifier (permission_id, permission_id_type, permission_type, permission_category, permission_tenant, permission_tenant_id) 
WHERE permission_tenant_id is not null;