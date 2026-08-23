CREATE SCHEMA IF NOT EXISTS kisansathi;

CREATE TABLE IF NOT EXISTS kisansathi.metadata_soil_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS kisansathi.soil_reports (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES kisansathi.users(id) ON DELETE SET NULL,
    soil_type VARCHAR(100) NOT NULL,
    ph NUMERIC(4, 2),
    nitrogen NUMERIC(10, 2),
    phosphorus NUMERIC(10, 2),
    potassium NUMERIC(10, 2),
    organic_carbon NUMERIC(10, 2),
    electrical_ec NUMERIC(10, 2),
    moisture NUMERIC(10, 2),
    status VARCHAR(30) NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS kisansathi.soil_recommendations (
    id SERIAL PRIMARY KEY,
    crop VARCHAR(100) NOT NULL,
    soil_type VARCHAR(100) NOT NULL,
    recommendation TEXT,
    fertilizer TEXT,
    irrigation TEXT,
    UNIQUE (crop, soil_type)
);
