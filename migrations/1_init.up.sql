Create table IF NOT EXISTS machines (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    media jsonbB DEFAULT '[]'::jsonb,
    created_at TIMESTAMP,
)

CREATE INDEX IF NOT EXISTS idx_machines_name ON machines (name);

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    third_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);

Create table if not exists flows (
    id serial PRIMARY KEY,
    problemname VARCHAR(255) NOT NULL,
    steps jsonb DEFAULT '[]'::jsonb,
    report TEXT,
);