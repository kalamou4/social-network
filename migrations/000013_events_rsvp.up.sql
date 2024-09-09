CREATE TABLE event_rsvp (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                -- Unique identifier for the RSVP
    event_id INTEGER NOT NULL REFERENCES events(id), -- Foreign key to the event
    user_id INTEGER NOT NULL REFERENCES users(id),   -- Foreign key to the user
    will_attend BOOLEAN NOT NULL,        -- Indicates if the user will attend
    rsvp_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- When the RSVP was made
);
