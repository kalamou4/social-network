CREATE TABLE private_messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                      -- Maps to `ID string`
    sender_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `SenderID string`
    recipient_id INTEGER NOT NULL REFERENCES users(id), -- Maps to `RecipientID string`
    message TEXT NOT NULL,                      -- Maps to `Message string`
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Maps to `SentAt time.Time`
);
