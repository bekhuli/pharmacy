CREATE TABLE IF NOT EXISTS utm_sources (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    content_name TEXT,
    source TEXT NOT NULL CHECK (source in ('tg', 'insta', 'fb', 'other')),
    campaign TEXT,
    content_creator TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);