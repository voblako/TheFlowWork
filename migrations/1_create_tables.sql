Create table IF NOT EXISTS company.assets (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    media jsonb,
    created_at TIMESTAMP,
)

CREATE INDEX IF NOT EXISTS idx_assets_name ON company.assets (name);

CREATE TABLE IF NOT EXISTS app.users (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    third_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    birthdate DATE
);

CREATE INDEX IF NOT EXISTS idx_users_email ON app.users (email);

CREATE IF NOT EXISTS TABLE company.problems (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    priority_level INT,
    created_at TIMESTAMP
)

Create table if not exists company.flows (
    id serial PRIMARY KEY,
    startdate TIMESTAMP,
    problem_id INT REFERENCES problems(id),
    machine_id INT REFERENCES machines(id),
    user_id INT REFERENCES users(id),
    desision_steps jsonb,
    created_at TIMESTAMP,
    report TEXT,
);