CREATE SCHEMA IF NOT EXISTS permission;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE permission.type AS ENUM ('CREATE', 'DELETE', 'VIEW', 'EDIT');
CREATE TYPE permission.category AS ENUM ('USER', 'BOOKING', 'PERMISSION', 'ROLE', 'TEAM', 'RESOURCE');
CREATE TYPE permission.tenant AS ENUM ('ROLE', 'USER', 'TEAM', 'ASSOCIATION' , 'PERMISSION', 'BUILDING', 'ROOM', 'ROOMASSOCIATION', 'IDENTIFIER', 'NA');

CREATE TABLE IF NOT EXISTS permission.role (
    id uuid DEFAULT uuid_generate_v4(),
    role_id uuid NOT NULL REFERENCES role.identifier(id) ON DELETE CASCADE,
    permission_type permission.type NOT NULL,
    permission_category permission.category NOT NULL,
    permission_tenant permission.tenant NOT NULL,
    permission_tenant_id uuid, -- If null then it means all
    date_added TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX permission_role_1 ON permission.role (role_id, permission_type, permission_category, permission_tenant, permission_tenant_id) 
WHERE permission_tenant_id is not null;
-- CREATE UNIQUE INDEX permission_role_2 ON permission.role (role_id, permission_type, permission_category, permission_tenant) 
-- WHERE permission_tenant_id is null;

CREATE TABLE IF NOT EXISTS permission.user (
    id uuid DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES "user".identifier(id) ON DELETE CASCADE,
    permission_type permission.type NOT NULL,
    permission_category permission.category NOT NULL,
    permission_tenant permission.tenant NOT NULL,
    permission_tenant_id uuid, -- If null then it means all
    date_added TIMESTAMP WITHOUT TIME ZONE DEFAULT(now() AT TIME ZONE 'uct'),
	
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX permission_user_1 ON permission.user (user_id, permission_type, permission_category, permission_tenant, permission_tenant_id) 
WHERE permission_tenant_id is not null;
-- CREATE UNIQUE INDEX permission_role_2 ON permission.role (role_id, permission_type, permission_category, permission_tenant) 
-- WHERE permission_tenant_id is null;