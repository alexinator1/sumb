CREATE TABLE IF NOT EXISTS business (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    owner_first_name VARCHAR(100) NOT NULL,
    owner_last_name VARCHAR(100) NOT NULL,
    owner_middle_name VARCHAR(100),
    owner_email VARCHAR(255) NOT NULL,
    owner_phone VARCHAR(30) NOT NULL,
    logo_id VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_working BOOLEAN NOT NULL DEFAULT true,
    deleted_at TIMESTAMPTZ,
    owner_id BIGINT
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_business_name ON business (name);