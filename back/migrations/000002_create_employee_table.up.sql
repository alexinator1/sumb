DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'employee_role') THEN
        CREATE TYPE employee_role AS ENUM ('regular','admin','owner');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'employee_status') THEN
        CREATE TYPE employee_status AS ENUM ('onboarding','active','vacation','ill');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS employee (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    password_hash VARCHAR(255) NOT NULL,
    phone VARCHAR(30),
    position VARCHAR(100),
    role employee_role NOT NULL DEFAULT 'regular',
    birth_date DATE,
    hired_at TIMESTAMPTZ,
    fired_at TIMESTAMPTZ,
    added_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    email VARCHAR(255) NOT NULL UNIQUE,
    status employee_status NOT NULL DEFAULT 'onboarding',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_by BIGINT,
    avatar_url VARCHAR(255),
    business_id BIGINT,
    CONSTRAINT fk_employee_created_by FOREIGN KEY (created_by) REFERENCES employee(id) ON DELETE SET NULL,
    CONSTRAINT fk_employee_business FOREIGN KEY (business_id) REFERENCES business(id) ON DELETE SET NULL
);

-- Indexes
CREATE UNIQUE INDEX IF NOT EXISTS idx_employee_email ON employee (email);
CREATE INDEX IF NOT EXISTS idx_employee_phone ON employee (phone);
CREATE INDEX IF NOT EXISTS idx_employee_business_id ON employee (business_id);