-- SQL script to create the database schema for Kisaan Sathi application
-- This script creates tables for users, posts, mandi prices, services, govt schemes, and notifications.
-- Make sure to run this script in a PostgreSQL database
CREATE SCHEMA IF NOT EXISTS kisansathi;

-- MANDI LOCATION MASTER DATA
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
    CONSTRAINT uq_mandi_district UNIQUE (state_id, name)
);

CREATE TABLE IF NOT EXISTS kisansathi.crops (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    variety VARCHAR(100),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE INDEX IF NOT EXISTS idx_mandi_districts_state
    ON kisansathi.mandi_districts (state_id);
-- One row represents a market/mandi and its location hierarchy.
CREATE TABLE IF NOT EXISTS kisansathi.mandi_details (
    id SERIAL PRIMARY KEY,
    mandi_code VARCHAR(50) UNIQUE,
    mandi_name VARCHAR(150) NOT NULL,
    district VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    pincode VARCHAR(10),
    address TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT uq_mandi_location UNIQUE (mandi_name, district, state)
);

CREATE INDEX IF NOT EXISTS idx_mandi_details_state
    ON kisansathi.mandi_details (state);
CREATE INDEX IF NOT EXISTS idx_mandi_details_district
    ON kisansathi.mandi_details (state, district);
CREATE INDEX IF NOT EXISTS idx_mandi_details_name
    ON kisansathi.mandi_details (mandi_name);

-- MANDI MARKET PRICES TABLE
-- Keeps the original columns and adds the fields used by the mandi queries.
CREATE TABLE IF NOT EXISTS kisansathi.market_prices (
    id INTEGER,
    mandi_id INTEGER,
    crop VARCHAR(50),
    price_per_quintal INTEGER,
    price_date VARCHAR(50),
    variety VARCHAR(100),
    market VARCHAR(100),
    district VARCHAR(100),
    state VARCHAR(100),
    min_price NUMERIC(12, 2),
    max_price NUMERIC(12, 2),
    modal_price NUMERIC(12, 2),
    arrival_date DATE,
    change NUMERIC(12, 2),
    trend VARCHAR(20)
);

-- Extend an existing table without removing legacy data.
ALTER TABLE kisansathi.market_prices
    ADD COLUMN IF NOT EXISTS variety VARCHAR(100),
    ADD COLUMN IF NOT EXISTS market VARCHAR(100),
    ADD COLUMN IF NOT EXISTS district VARCHAR(100),
    ADD COLUMN IF NOT EXISTS state VARCHAR(100),
    ADD COLUMN IF NOT EXISTS min_price NUMERIC(12, 2),
    ADD COLUMN IF NOT EXISTS max_price NUMERIC(12, 2),
    ADD COLUMN IF NOT EXISTS modal_price NUMERIC(12, 2),
    ADD COLUMN IF NOT EXISTS arrival_date DATE,
    ADD COLUMN IF NOT EXISTS change NUMERIC(12, 2),
    ADD COLUMN IF NOT EXISTS trend VARCHAR(20);

CREATE INDEX IF NOT EXISTS idx_market_prices_crop
    ON kisansathi.market_prices (crop);
CREATE INDEX IF NOT EXISTS idx_market_prices_state
    ON kisansathi.market_prices (state);
CREATE INDEX IF NOT EXISTS idx_market_prices_district
    ON kisansathi.market_prices (district);
CREATE INDEX IF NOT EXISTS idx_market_prices_arrival_date
    ON kisansathi.market_prices (arrival_date);

-- CROP CATALOG AND RECOMMENDATIONS
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
    CONSTRAINT uq_crop_variety UNIQUE (crop_id, name)
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
    CONSTRAINT uq_crop_disease UNIQUE (crop_id, name)
);

CREATE TABLE IF NOT EXISTS kisansathi.crop_recommendation (
    id SERIAL PRIMARY KEY,
    crop_id INTEGER REFERENCES kisansathi.crop(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    soil_type VARCHAR(100) NOT NULL,
    district VARCHAR(100) NOT NULL,
    confidence INTEGER NOT NULL CHECK (confidence BETWEEN 0 AND 100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT uq_crop_recommendation UNIQUE (name, soil_type, district)
);

CREATE INDEX IF NOT EXISTS idx_crop_season
    ON kisansathi.crop (season);
CREATE INDEX IF NOT EXISTS idx_crop_category
    ON kisansathi.crop (category_id);
CREATE INDEX IF NOT EXISTS idx_crop_soil_type
    ON kisansathi.crop (soil_type);
CREATE INDEX IF NOT EXISTS idx_crop_varieties_crop
    ON kisansathi.crop_varieties (crop_id);
CREATE INDEX IF NOT EXISTS idx_crop_diseases_crop
    ON kisansathi.crop_diseases (crop_id);
CREATE INDEX IF NOT EXISTS idx_crop_recommendation_lookup
    ON kisansathi.crop_recommendation (soil_type, district);

-- USERS TABLE
CREATE TABLE IF NOT EXISTS kisansathi.users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    phone VARCHAR(20) UNIQUE,
    role VARCHAR(20) CHECK (role IN ('farmer', 'advisor', 'scientist')),
    language VARCHAR(20),
    soil_type VARCHAR(50),
    district VARCHAR(100),
    lat DOUBLE PRECISION,
    lng DOUBLE PRECISION,
    created_at TIMESTAMP DEFAULT now()
);

-- POSTS TABLE
CREATE TABLE IF NOT EXISTS kisansathi.posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES kisansathi.users(id),
    caption TEXT,
    media_url TEXT,
    crop_tag VARCHAR(50),
    likes INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT now()
);

-- MANDI PRICES TABLE
CREATE TABLE IF NOT EXISTS kisansathi.mandi_prices (
    id SERIAL PRIMARY KEY,
    crop VARCHAR(100),
    region VARCHAR(100),
    price INTEGER,
    recorded_on DATE DEFAULT CURRENT_DATE
);

-- SERVICES TABLE
CREATE TABLE IF NOT EXISTS kisansathi.services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    type VARCHAR(50),
    contact VARCHAR(50),
    address TEXT,
    lat DOUBLE PRECISION,
    lng DOUBLE PRECISION
);

-- GOVT SCHEMES TABLE
CREATE TABLE IF NOT EXISTS kisansathi.govt_schemes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200),
    description TEXT,
    eligibility TEXT,
    tags TEXT[],
    pdf_url TEXT,
    created_at TIMESTAMP DEFAULT now()
);

-- NOTIFICATIONS TABLE
CREATE TABLE IF NOT EXISTS kisansathi.notifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES kisansathi.users(id),
    message TEXT,
    type VARCHAR(50),
    read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now()
);


-- Insert sample users (farmers, advisors, scientists)
INSERT INTO users (name, phone, role, language, soil_type, district, lat, lng) VALUES
('Ravi Yadav', '9876543210', 'farmer', 'Hindi', 'Loamy', 'Barabanki', 26.9371, 81.1895),
('Suman Verma', '9123456789', 'advisor', 'Hindi', NULL, 'Gorakhpur', 26.7606, 83.3732),
('Dr. Patel', '9988776655', 'scientist', 'English', NULL, 'Lucknow', 26.8467, 80.9462);

-- Insert sample posts
INSERT INTO posts (user_id, caption, media_url, crop_tag, likes) VALUES
(1, 'My wheat crop after organic fertilizer use!', 'https://example.com/img/wheat1.jpg', 'wheat', 12),
(1, 'Need help identifying this pest on brinjal', 'https://example.com/img/brinjal_bug.jpg', 'brinjal', 5);

-- Insert sample mandi prices
INSERT INTO mandi_prices (crop, region, price) VALUES
('Wheat', 'Barabanki', 2250),
('Rice', 'Lucknow', 1850),
('Potato', 'Agra', 800);

-- Insert sample services
INSERT INTO services (name, type, contact, address, lat, lng) VALUES
('Krishi Mitra Vet Center', 'vet', '7523999912', 'Barabanki Road', 26.9368, 81.1900),
('Soil Testing Lab – DeHaat', 'soil', '7412589630', 'Gorakhpur Sector 3', 26.7610, 83.3735);

-- Insert sample govt schemes
INSERT INTO govt_schemes (title, description, eligibility, tags, pdf_url) VALUES
('PM-Kisan Yojana', 'Rs. 6000/year direct to farmers bank accounts', 'All small & marginal farmers', ARRAY['income', 'direct-benefit'], 'https://example.gov/pm-kisansathi.pdf'),
('Fasal Bima Yojana', 'Insurance cover for crop damage due to climate risks', 'All registered farmers', ARRAY['insurance', 'climate'], 'https://example.gov/fasal-bima.pdf');

-- Insert sample notifications
INSERT INTO notifications (user_id, message, type) VALUES
(1, 'Wheat price has increased to ₹2250 in Barabanki', 'price'),
(1, 'PM-Kisan scheme deadline extended to July 15', 'scheme');
