CREATE TABLE IF NOT EXISTS cards(
    id INT NOT NULL AUTO_INCREMENT KEY,
    title VARCHAR(255),
    description TEXT,
    PRIMARY KEY (id),
    FOREIGN KEY (column_id) REFERENCES columns(id)
);