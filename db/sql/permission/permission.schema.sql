CREATE SCHEMA IF NOT EXISTS permission;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE permission.type AS ENUM ('CREATE', 'DELETE', 'VIEW', 'EDIT');
CREATE TYPE permission.category AS ENUM ('USER', 'BOOKING', 'PERMISSION');
CREATE TYPE permission.tenant AS ENUM ('ROLE', 'USER', 'TEAM'); -- If null then it means all

CREATE TABLE IF NOT EXISTS permission.role (
    role_id uuid NOT NULL REFERENCES role.identifier(id) ON DELETE CASCADE,
    permission_type permission.type NOT NULL,
    permission_category permission.category NOT NULL,
    permission_tenant permission.tenant NOT NULL,
    permission_tenant_id uuid,
    date_added TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (role_id, permission_type, permission_category, permission_tenant, permission_tenant_id)
);

CREATE TABLE IF NOT EXISTS permission.user (
    user_id uuid NOT NULL REFERENCES "user".identifier(id) ON DELETE CASCADE,
    permission_type permission.type NOT NULL,
    permission_category permission.category NOT NULL,
    permission_tenant permission.tenant NOT NULL,
    permission_tenant_id uuid,
    date_added TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (user_id, permission_type, permission_category, permission_tenant, permission_tenant_id)
);