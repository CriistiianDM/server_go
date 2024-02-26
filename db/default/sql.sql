-- Postgresql

-- @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>

-- Create Database
SELECT 'CREATE DATABASE server_go'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'server_go')\gexec

-- Connect to Database
\connect server_go

-- Sequences
CREATE SEQUENCE IF NOT EXISTS ip_address_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

CREATE SEQUENCE IF NOT EXISTS routes_defautls_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;


-- Ip Address Access
CREATE TABLE IF NOT EXISTS ip_address_access (
    id BIGINT PRIMARY KEY DEFAULT nextval('ip_address_id_seq'),
    ip_address VARCHAR(50) NOT NULL UNIQUE,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE  IF NOT EXISTS routes_defautls (
    id BIGINT PRIMARY KEY DEFAULT nextval('routes_defautls_id_seq'),
    route VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create Indexes


-- PROCEDURES

-- FUNCTIONS - QUERYS

-- Insert Rows Values

INSERT INTO routes_defautls (route) 
values ('my-server-go');

INSERT INTO ip_address_access (ip_address, active)
VALUES ('127.0.0.1', TRUE);