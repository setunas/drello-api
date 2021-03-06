CREATE TABLE IF NOT EXISTS users(
    id INT NOT NULL AUTO_INCREMENT KEY,
    username VARCHAR(255),
    firebase_uid VARCHAR(255) NOT NULL UNIQUE,
    board_id INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (board_id) REFERENCES boards(id)
);