CREATE TABLE IF NOT EXISTS staffs
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    staff_code VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    birth_date date,
    gender integer,
    phone_number VARCHAR(20),
    address VARCHAR,
    unit_id integer NOT NULL,
    staff_type_id integer NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_unit_id FOREIGN KEY (unit_id) REFERENCES units(id),
    CONSTRAINT fk_staff_type_id FOREIGN KEY (staff_type_id) REFERENCES staff_types(id)
);
