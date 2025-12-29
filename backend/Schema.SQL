-- Clients
CREATE TABLE clients (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Rooms
CREATE TABLE rooms (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Client-Room mapping
CREATE TABLE client_rooms (
    client_id TEXT REFERENCES clients(id),
    room_id TEXT REFERENCES rooms(id),
    PRIMARY KEY (client_id, room_id)
);

-- Messages
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    from_client TEXT REFERENCES clients(id),
    room_id TEXT REFERENCES rooms(id),
    body TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT NOW()
);

-- Offline messages
CREATE TABLE offline_messages (
    id SERIAL PRIMARY KEY,
    client_id TEXT REFERENCES clients(id),
    message_id INT REFERENCES messages(id),
    delivered BOOLEAN DEFAULT FALSE
);
