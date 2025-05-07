CREATE TABLE IF NOT EXISTS utm_sources (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    content_name VARCHAR(255),
    source TEXT NOT NULL CHECK (source in ('tg', 'insta', 'fb', 'other')),
    campaign VARCHAR(255),
    content_creator VARCHAR(255),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);