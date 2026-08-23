CREATE SCHEMA IF NOT EXISTS kisansathi;

CREATE TABLE IF NOT EXISTS kisansathi.equipment_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS kisansathi.equipments (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES kisansathi.equipment_categories(id) ON DELETE SET NULL,
    name VARCHAR(150) NOT NULL,
    brand VARCHAR(100),
    model VARCHAR(100),
    description TEXT,
    rental_price NUMERIC(12, 2),
    purchase_price NUMERIC(12, 2),
    availability_status VARCHAR(30) NOT NULL DEFAULT 'AVAILABLE',
    owner_id INTEGER REFERENCES kisansathi.users(id) ON DELETE SET NULL,
    location VARCHAR(150),
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_equipments_category_location
    ON kisansathi.equipments (category_id, location);
