CREATE SCHEMA IF NOT EXISTS kisansathi;

CREATE TABLE IF NOT EXISTS kisansathi.users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    phone VARCHAR(20) UNIQUE,
    email VARCHAR(150) UNIQUE,
    password TEXT,
    role VARCHAR(20) NOT NULL DEFAULT 'farmer',
    language VARCHAR(20),
    soil_type VARCHAR(50),
    district VARCHAR(100),
    lat DOUBLE PRECISION,
    lng DOUBLE PRECISION,
    avatar_url TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_users_district
    ON kisansathi.users (district);
