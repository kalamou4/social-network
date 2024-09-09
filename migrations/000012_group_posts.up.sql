CREATE TABLE group_posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                      -- Maps to `ID string`
    group_id INTEGER NOT NULL REFERENCES groups(id), -- Maps to `GroupID string`
    user_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `UserID string`
    content TEXT NOT NULL,                      -- Maps to `Content string`
    image_url VARCHAR(255),                     -- Maps to `ImageURL string`
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Maps to `CreatedAt time.Time`
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Maps to `UpdatedAt time.Time`
);
