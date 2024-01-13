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

CREATE SEQUENCE IF NOT EXISTS products_routes_company_id_seq
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


-- Tabla de productos x route
CREATE TABLE IF NOT EXISTS products_routes_company (
    id BIGINT PRIMARY KEY DEFAULT nextval('products_routes_company_id_seq'),
    route VARCHAR(50) NOT NULL,
    product_id BIGINT NOT NULL,
    json_data JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Se me olvido ponerle el UNIQUE a la columna route :3 demasiaaado tarde
ALTER TABLE routes_company ADD CONSTRAINT routes_company_route_key UNIQUE (route);

-- Indexes
CREATE INDEX IF NOT EXISTS routes_company_route_idx ON routes_company (route);
CREATE INDEX IF NOT EXISTS products_routes_company_route_idx ON products_routes_company (route);


-- Insert Data
INSERT INTO routes_company (route, created_at , updated_at) VALUES ('cristiank', NOW(), NOW());
INSERT INTO routes_company (route, created_at , updated_at) VALUES ('fosk', NOW(), NOW());


-- Insert Data
INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at) 
VALUES ('cristiank', 4, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('cristiank', 5, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('cristiank', 6, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('cristiank', 7, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('cristiank', 8, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

-- whit fosk

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('fosk', 4, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('fosk', 5, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('fosk', 6, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('fosk', 7, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, product_id, json_data, created_at , updated_at)
VALUES ('fosk', 8, '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());