CREATE TABLE IF NOT EXISTS medicine_illness_links (
    medicine_id UUID NOT NULL,
    illness_id UUID NOT NULL,
    PRIMARY KEY (medicine_id, illness_id),
    CONSTRAINT fk_medicine FOREIGN KEY (medicine_id) REFERENCES medicines(id) ON DELETE CASCADE,
    CONSTRAINT fk_illness FOREIGN KEY (illness_id) REFERENCES illnesses(id) ON DELETE CASCADE
)