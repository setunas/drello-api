CREATE TABLE IF NOT EXISTS columns(
    id INT NOT NULL AUTO_INCREMENT KEY,
    title VARCHAR(255),
    board_id INT NOT NULL,
    FOREIGN KEY (board_id) REFERENCES boards(id)
);