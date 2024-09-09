CREATE TABLE followers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                      -- Maps to `ID string`
    follower_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `FollowerID string`
    followed_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `FollowedID string`
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Maps to `CreatedAt time.Time`
    status VARCHAR(50) DEFAULT 'pending'        -- Maps to `Status string`
);
