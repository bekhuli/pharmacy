CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    phone VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    is_admin BOOLEAN NOT NULL DEFAULT FALSE
);