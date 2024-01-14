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
    json_data JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Functions
CREATE OR REPLACE FUNCTION products_per_company(route_p TEXT)
RETURNS TABLE (
    id BIGINT,
    route  VARCHAR(50),
    json_data JSONB,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)
AS $$
BEGIN
    RETURN QUERY
    SELECT * 
    FROM products_routes_company as t1
    WHERE t1.route = route_p;
END;
$$
LANGUAGE plpgsql;

-- Se me olvido ponerle el UNIQUE a la columna route :3 demasiaaado tarde
ALTER TABLE routes_company ADD CONSTRAINT routes_company_route_key UNIQUE (route);

-- Se elimna el campo product_id
-- ALTER TABLE products_routes_company DROP COLUMN product_id;

-- Indexes
CREATE INDEX IF NOT EXISTS routes_company_route_idx ON routes_company (route);
CREATE INDEX IF NOT EXISTS products_routes_company_route_idx ON products_routes_company (route);


CREATE OR REPLACE PROCEDURE insert_new_product_per_company(route_p TEXT, json_ JSONB)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO products_routes_company (route, json_data, created_at , updated_at) 
    VALUES (route_p, json_, NOW(), NOW()); 
END;
$$;

-- Retorna todos los productos de una ruta
CREATE OR REPLACE PROCEDURE get_product_per_company(route_p TEXT,OUT id BIGINT, OUT route TEXT, OUT json_data JSONB, OUT created_at TIMESTAMP, OUT updated_at TIMESTAMP)
LANGUAGE plpgsql
AS $$
BEGIN
    SELECT * 
    INTO id, route, json_data, created_at, updated_at
    FROM products_routes_company
    WHERE route = route_p;
END;
$$;


-- Insert Data
INSERT INTO routes_company (route, created_at , updated_at) VALUES ('cristiank', NOW(), NOW());
INSERT INTO routes_company (route, created_at , updated_at) VALUES ('fosk', NOW(), NOW());


-- Insert Data
INSERT INTO products_routes_company (route, json_data, created_at , updated_at) 
VALUES ('cristiank','{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('cristiank','{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('cristiank','{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('cristiank','{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('cristiank','{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

-- whit fosk

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('fosk', '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('fosk', '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('fosk','{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('fosk', '{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());

INSERT INTO products_routes_company (route, json_data, created_at , updated_at)
VALUES ('fosk','{"nameProduct": "tv", "price": "27.08", "store": "3", "qty": "2" , "discount": "3" , "idItem": "4" , "region": "CO"}', NOW(), NOW());