package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID           string `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
}

type ChatRoom struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	IsPrivate   bool    `json:"is_private" db:"is_private"`
	CreatedBy   *string `json:"created_by,omitempty" db:"created_by"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
}

type Message struct {
	ID         string  `json:"id" db:"id"`
	ChatRoomID string  `json:"chat_room_id" db:"chat_room_id"`
	SenderID   *string `json:"sender_id,omitempty" db:"sender_id"`
	Content    string  `json:"content" db:"content"`
	SentAt     string  `json:"sent_at" db:"sent_at"`
}

type PrivateMessage struct {
	ID         string  `json:"id" db:"id"`
	SenderID   *string `json:"sender_id,omitempty" db:"sender_id"`
	ReceiverID *string `json:"receiver_id,omitempty" db:"receiver_id"`
	Content    string  `json:"content" db:"content"`
	SentAt     string  `json:"sent_at" db:"sent_at"`
}

type UserRoom struct {
	UserID     string `json:"user_id" db:"user_id"`
	ChatRoomID string `json:"chat_room_id" db:"chat_room_id"`
	JoinedAt   string `json:"joined_at" db:"joined_at"`
}

type MessageIndex struct {
	ChatRoomID string `json:"chat_room_id" db:"chat_room_id"`
	SenderID   string `json:"sender_id" db:"sender_id"`
}

type PrivateMessageIndex struct {
	SenderID   string `json:"sender_id" db:"sender_id"`
	ReceiverID string `json:"receiver_id" db:"receiver_id"`
}

// Struct for credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// JWT Claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
