CREATE SCHEMA IF NOT EXISTS kisansathi;

CREATE TABLE IF NOT EXISTS kisansathi.crop_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS kisansathi.crop (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES kisansathi.crop_categories(id) ON DELETE SET NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    season VARCHAR(50) NOT NULL,
    soil_type VARCHAR(100) NOT NULL,
    duration INTEGER NOT NULL CHECK (duration > 0),
    image_url TEXT,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS kisansathi.crop_varieties (
    id SERIAL PRIMARY KEY,
    crop_id INTEGER NOT NULL REFERENCES kisansathi.crop(id) ON DELETE CASCADE,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    maturity_days INTEGER CHECK (maturity_days > 0),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (crop_id, name)
);

CREATE TABLE IF NOT EXISTS kisansathi.crop_diseases (
    id SERIAL PRIMARY KEY,
    crop_id INTEGER NOT NULL REFERENCES kisansathi.crop(id) ON DELETE CASCADE,
    name VARCHAR(150) NOT NULL,
    symptoms TEXT,
    causes TEXT,
    prevention TEXT,
    treatment TEXT,
    image_url TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (crop_id, name)
);

CREATE TABLE IF NOT EXISTS kisansathi.crop_recommendation (
    id SERIAL PRIMARY KEY,
    crop_id INTEGER REFERENCES kisansathi.crop(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    soil_type VARCHAR(100) NOT NULL,
    district VARCHAR(100) NOT NULL,
    confidence INTEGER NOT NULL CHECK (confidence BETWEEN 0 AND 100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (name, soil_type, district)
);
