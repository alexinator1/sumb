DROP TABLE IF EXISTS employee CASCADE;

-- Drop enum types created for employee
DO $$
BEGIN
	IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'employee_role') THEN
		DROP TYPE employee_role;
	END IF;
	IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'employee_status') THEN
		DROP TYPE employee_status;
	END IF;
END$$;
