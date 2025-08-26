
-- Section1
   -- Create events table
CREATE TABLE events (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    date DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Section2
   -- Create event_user table
CREATE TABLE event_user (
    user_id BIGINT UNSIGNED,
    event_id BIGINT UNSIGNED,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (event_id) REFERENCES events(id)
);