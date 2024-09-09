CREATE TABLE comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                      -- Maps to `ID string`
    post_id INTEGER NOT NULL REFERENCES posts(id), -- Maps to `PostID string`
    user_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `UserID string`
    content TEXT NOT NULL,                      -- Maps to `Content string`
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Maps to `CreatedAt time.Time`
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Maps to `UpdatedAt time.Time`
);
