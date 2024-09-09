CREATE TABLE group_members (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                      -- Maps to `ID string`
    group_id INTEGER NOT NULL REFERENCES groups(id), -- Maps to `GroupID string`
    user_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `UserID string`
    role VARCHAR(50) DEFAULT 'member',          -- Maps to `Role string`
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Maps to `JoinedAt time.Time`
);
