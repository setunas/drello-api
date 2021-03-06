CREATE TABLE IF NOT EXISTS cards(
    id INT NOT NULL AUTO_INCREMENT KEY,
    title VARCHAR(255),
    description TEXT,
    column_id INT NOT NULL,
    position DOUBLE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (column_id) REFERENCES columns(id) ON DELETE CASCADE
);