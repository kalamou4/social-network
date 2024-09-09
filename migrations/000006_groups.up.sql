CREATE TABLE groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                      -- Maps to `ID string`
    creator_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `CreatorID string`
    title VARCHAR(255) NOT NULL,                -- Maps to `Title string`
    description TEXT,                           -- Maps to `Description string`
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Maps to `CreatedAt time.Time`
);
