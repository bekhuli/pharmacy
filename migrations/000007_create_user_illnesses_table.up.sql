CREATE TABLE IF NOT EXISTS user_illnesses (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    illness_id UUID NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_illness FOREIGN KEY (illness_id) REFERENCES illnesses(id) ON DELETE CASCADE
);