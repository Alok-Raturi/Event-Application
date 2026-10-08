CREATE TABLE event_attendees (
    user_id UUID NOT NULL REFERENCES users(id),
    event_id UUID NOT NULL REFERENCES events(id),
    PRIMARY KEY (user_id, event_id)
)