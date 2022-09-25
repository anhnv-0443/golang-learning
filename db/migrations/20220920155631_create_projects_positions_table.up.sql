CREATE TABLE IF NOT EXISTS projects_positions
(
    id BIGSERIAL PRIMARY KEY,
    project_id BIGSERIAL NOT NULL,
    position_id BIGSERIAL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_position_id FOREIGN KEY (position_id) REFERENCES positions(id),
    CONSTRAINT fk_project_id FOREIGN KEY (project_id) REFERENCES projects(id)
);
