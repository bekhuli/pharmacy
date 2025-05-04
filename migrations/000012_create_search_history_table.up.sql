CREATE TABLE IF NOT EXISTS search_history (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    query TEXT NOT NULL,
    searched_at TIMESTAMPTZ NOT NULL DEFAULT now()
)