CREATE TABLE IF NOT EXISTS medicines (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    manufacturer TEXT
)