-- MySQL-compatible schema for a chat room server

-- Users table: Stores all registered users
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Chat Rooms table: Stores all available chat rooms
CREATE TABLE chat_rooms (
    id CHAR(36) PRIMARY KEY NOT NULL,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    is_private BOOLEAN DEFAULT FALSE,
    created_by CHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL
);

-- Messages table: Stores messages in chat rooms
CREATE TABLE messages (
    id CHAR(36) PRIMARY KEY NOT NULL,
    chat_room_id CHAR(36) NOT NULL,
    sender_id CHAR(36),
    content TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE SET NULL
);

-- Private Messages table: Stores private messages between users
CREATE TABLE private_messages (
    id CHAR(36) PRIMARY KEY NOT NULL,
    sender_id CHAR(36),
    receiver_id CHAR(36),
    content TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE SET NULL,
    FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE SET NULL
);

-- User_Rooms table: Manages user membership in chat rooms
CREATE TABLE user_rooms (
    user_id CHAR(36) NOT NULL,
    chat_room_id CHAR(36) NOT NULL,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, chat_room_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (chat_room_id) REFERENCES chat_rooms(id) ON DELETE CASCADE
);

-- Indexing for performance
CREATE INDEX idx_messages_chat_room_id ON messages(chat_room_id);
CREATE INDEX idx_messages_sender_id ON messages(sender_id);
CREATE INDEX idx_private_messages_sender_receiver ON private_messages(sender_id, receiver_id);

