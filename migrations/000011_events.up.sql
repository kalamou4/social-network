CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                      -- Maps to `ID string`
    group_id INTEGER NOT NULL REFERENCES groups(id), -- Maps to `GroupID string`
    title VARCHAR(255) NOT NULL,                -- Maps to `Title string`
    description TEXT,                           -- Maps to `Description string`
    date_time TIMESTAMP NOT NULL,               -- Maps to `DateTime time.Time`
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Maps to `CreatedAt time.Time`
);
