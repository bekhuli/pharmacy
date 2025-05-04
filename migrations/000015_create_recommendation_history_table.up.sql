CREATE TABLE IF NOT EXISTS recommendation_history (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    medicine_id UUID NOT NULL,
    is_ordered BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT fk_rehistory_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_rehistory_medicine FOREIGN KEY (medicine_id) REFERENCES medicines(id) ON DELETE CASCADE
)