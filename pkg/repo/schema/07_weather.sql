CREATE SCHEMA IF NOT EXISTS kisansathi;

CREATE TABLE IF NOT EXISTS kisansathi.weather_observations (
    id BIGSERIAL PRIMARY KEY,
    location VARCHAR(150) NOT NULL,
    district VARCHAR(100),
    state VARCHAR(100),
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    observed_at TIMESTAMPTZ NOT NULL,
    temperature NUMERIC(6, 2),
    feels_like NUMERIC(6, 2),
    humidity NUMERIC(5, 2),
    rainfall NUMERIC(8, 2),
    wind_speed NUMERIC(8, 2),
    weather_condition VARCHAR(100),
    raw_data JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_weather_location_time
    ON kisansathi.weather_observations (location, observed_at DESC);
