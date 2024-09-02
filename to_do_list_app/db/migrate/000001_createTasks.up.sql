CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date timestamp NULL DEFAULT NULL,
    created_at  timestamp NOT NULL DEFAULT Now(),
    updated_at timestamp NOT NULL DEFAULT Now()
);
