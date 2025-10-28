-- Migration: create business table
-- Creates table: business

CREATE TABLE IF NOT EXISTS business (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    owner_first_name VARCHAR(100),
    owner_second_name VARCHAR(100),
    owner_middle_name VARCHAR(100),
    owner_email VARCHAR(255),
    owner_phone VARCHAR(30),
    logo_id VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_working SMALLINT NOT NULL DEFAULT 1,
    deleted_at TIMESTAMPTZ,
    owner_id BIGINT,
    CONSTRAINT fk_owner_employee FOREIGN KEY (owner_id) REFERENCES employee(id) ON DELETE SET NULL
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_business_name ON business (name);
CREATE INDEX IF NOT EXISTS idx_business_owner_id ON business (owner_id);
