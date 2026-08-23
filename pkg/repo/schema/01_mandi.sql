CREATE SCHEMA IF NOT EXISTS kisansathi;

CREATE TABLE IF NOT EXISTS kisansathi.mandi_states (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(10) UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS kisansathi.mandi_districts (
    id SERIAL PRIMARY KEY,
    state_id INTEGER NOT NULL REFERENCES kisansathi.mandi_states(id),
    name VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    UNIQUE (state_id, name)
);

CREATE TABLE IF NOT EXISTS kisansathi.mandi_details (
    id SERIAL PRIMARY KEY,
    mandi_code VARCHAR(50) UNIQUE,
    mandi_name VARCHAR(150) NOT NULL,
    district VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    address TEXT,
    pincode VARCHAR(10),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (mandi_name, district, state)
);

CREATE TABLE IF NOT EXISTS kisansathi.market_prices (
    id SERIAL PRIMARY KEY,
    mandi_id INTEGER REFERENCES kisansathi.mandi_details(id) ON DELETE SET NULL,
    crop VARCHAR(100) NOT NULL,
    variety VARCHAR(100),
    min_price NUMERIC(12, 2),
    max_price NUMERIC(12, 2),
    modal_price NUMERIC(12, 2),
    price_per_quintal NUMERIC(12, 2),
    arrival_date DATE,
    price_date VARCHAR(50),
    change NUMERIC(12, 2),
    trend VARCHAR(20),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_mandi_details_state_district
    ON kisansathi.mandi_details (state, district);
CREATE INDEX IF NOT EXISTS idx_market_prices_crop_date
    ON kisansathi.market_prices (crop, arrival_date);
