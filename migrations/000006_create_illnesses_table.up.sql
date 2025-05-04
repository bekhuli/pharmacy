CREATE TABLE IF NOT EXISTS illnesses (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    category_id UUID REFERENCES illness_categories(id) ON DELETE SET NULL
);