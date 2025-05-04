CREATE TABLE IF NOT EXISTS user_criteria (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    age INTEGER,
    job TEXT,
    gender TEXT CHECK (gender IN ('male', 'female', 'other')),
    is_married BOOLEAN NOT NULL DEFAULT FALSE
);