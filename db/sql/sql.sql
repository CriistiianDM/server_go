-- Author: Cristian Machado (cristian.machado@correounivalle.edu.co)

-- Create Database
SELECT 'CREATE DATABASE server_go'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'server_go')\gexec

-- Sequences
CREATE SEQUENCE IF NOT EXISTS routes_company_id_seq
INCREMENT 1
MINVALUE 1
MAXVALUE 1000000
START 1
CACHE 1;

-- Tables
CREATE TABLE IF NOT EXISTS routes_company (
    id BIGINT PRIMARY KEY DEFAULT nextval('routes_company_id_seq'),
    route VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS routes_company_route_idx ON routes_company (route);


-- Insert Data
INSERT INTO routes_company (route, created_at , updated_at) VALUES ('cristiank', NOW(), NOW());