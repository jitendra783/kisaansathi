CREATE SCHEMA IF NOT EXISTS kisansathi;

CREATE TABLE IF NOT EXISTS kisansathi.fertilizers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    type VARCHAR(100),
    npk_ratio VARCHAR(50),
    description TEXT,
    application_method TEXT,
    dosage VARCHAR(100),
    suitable_crops TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (name, type)
);

CREATE TABLE IF NOT EXISTS kisansathi.crop_fertilizers (
    id SERIAL PRIMARY KEY,
    crop_id INTEGER NOT NULL REFERENCES kisansathi.crop(id) ON DELETE CASCADE,
    fertilizer_id INTEGER NOT NULL REFERENCES kisansathi.fertilizers(id) ON DELETE CASCADE,
    growth_stage VARCHAR(100),
    recommended_dosage VARCHAR(100),
    notes TEXT,
    UNIQUE (crop_id, fertilizer_id, growth_stage)
);

CREATE INDEX IF NOT EXISTS idx_crop_fertilizers_crop
    ON kisansathi.crop_fertilizers (crop_id);
