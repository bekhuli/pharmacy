CREATE TABLE IF NOT EXISTS illnesses (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id UUID REFERENCES illness_categories(id) ON DELETE SET NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP
);