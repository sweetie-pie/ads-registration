package http

import "time"

type TokenResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type AdResponse struct {
	ID          uint         `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      int          `json:"status"`
	Image       string       `json:"image"`
	Categories  []string     `json:"categories"`
	CreatedAt   time.Time    `json:"created_at"`
	User        UserResponse `json:"username"`
}
