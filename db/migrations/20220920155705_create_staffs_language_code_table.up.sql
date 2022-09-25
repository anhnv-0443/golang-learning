CREATE TABLE IF NOT EXISTS staffs_language_codes
(
    id BIGSERIAL PRIMARY KEY,
    staff_id BIGSERIAL NOT NULL,
    language_code_id BIGSERIAL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_staff_id FOREIGN KEY (staff_id) REFERENCES staffs(id),
    CONSTRAINT fk_language_code_id FOREIGN KEY (language_code_id) REFERENCES language_codes(id)
);
