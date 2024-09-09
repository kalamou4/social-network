CREATE TABLE notifications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                      -- Maps to `ID string`
    user_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `UserID string`
    message TEXT NOT NULL,                      -- Maps to `Message string`
    link VARCHAR(255),                          -- Maps to `Link string`
    type VARCHAR(50) NOT NULL,                  -- Maps to `Type string`
    read BOOLEAN DEFAULT FALSE,                 -- Maps to `Read bool`
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Maps to `CreatedAt time.Time`
);
