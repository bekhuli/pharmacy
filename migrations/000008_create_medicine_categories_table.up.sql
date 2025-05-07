CREATE TABLE IF NOT EXISTS medicine_categories (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP
)