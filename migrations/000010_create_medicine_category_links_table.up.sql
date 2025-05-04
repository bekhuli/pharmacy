CREATE TABLE IF NOT EXISTS medicine_category_links (
    medicine_id UUID NOT NULL,
    category_id UUID NOT NULL,
    PRIMARY KEY (medicine_id, category_id),
    CONSTRAINT fk_medicine FOREIGN KEY (medicine_id) REFERENCES medicines(id) ON DELETE CASCADE,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES medicine_categories(id) ON DELETE CASCADE
);