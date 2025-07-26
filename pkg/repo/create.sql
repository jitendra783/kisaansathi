-- SQL script to create the database schema for Kisaan Sathi application
-- This script creates tables for users, posts, mandi prices, services, govt schemes, and notifications.
-- Make sure to run this script in a PostgreSQL database
-- SCHEMA CREATION
CREATE SCHEMA IF NOT EXISTS kisan;

-- USERS TABLE
CREATE TABLE IF NOT EXISTS kisan.users (
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
CREATE TABLE IF NOT EXISTS kisan.posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES kisan.users(id),
    caption TEXT,
    media_url TEXT,
    crop_tag VARCHAR(50),
    likes INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT now()
);

-- MANDI PRICES TABLE
CREATE TABLE IF NOT EXISTS kisan.mandi_prices (
    id SERIAL PRIMARY KEY,
    crop VARCHAR(100),
    region VARCHAR(100),
    price INTEGER,
    recorded_on DATE DEFAULT CURRENT_DATE
);

-- SERVICES TABLE
CREATE TABLE IF NOT EXISTS kisan.services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    type VARCHAR(50),
    contact VARCHAR(50),
    address TEXT,
    lat DOUBLE PRECISION,
    lng DOUBLE PRECISION
);

-- GOVT SCHEMES TABLE
CREATE TABLE IF NOT EXISTS kisan.govt_schemes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200),
    description TEXT,
    eligibility TEXT,
    tags TEXT[],
    pdf_url TEXT,
    created_at TIMESTAMP DEFAULT now()
);

-- NOTIFICATIONS TABLE
CREATE TABLE IF NOT EXISTS kisan.notifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES kisan.users(id),
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
('PM-Kisan Yojana', 'Rs. 6000/year direct to farmers bank accounts', 'All small & marginal farmers', ARRAY['income', 'direct-benefit'], 'https://example.gov/pm-kisan.pdf'),
('Fasal Bima Yojana', 'Insurance cover for crop damage due to climate risks', 'All registered farmers', ARRAY['insurance', 'climate'], 'https://example.gov/fasal-bima.pdf');

-- Insert sample notifications
INSERT INTO notifications (user_id, message, type) VALUES
(1, 'Wheat price has increased to ₹2250 in Barabanki', 'price'),
(1, 'PM-Kisan scheme deadline extended to July 15', 'scheme');
