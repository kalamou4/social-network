CREATE TABLE sessions (
    session_id STRING PRIMARY KEY,              -- Maps to `SessionID string`
    user_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `UserID string`
    expires_at TIMESTAMP NOT NULL               -- Maps to `ExpiresAt time.Time`
);
