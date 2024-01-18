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
	Banned    bool      `json:"banned"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminResponse struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Active      bool      `json:"active"`
	AccessLevel int       `json:"access_level"`
	CreatedAt   time.Time `json:"created_at"`
}

type AdResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	Image       string    `json:"image"`
	Username    string    `json:"username"`
	Admin       string    `json:"admin"`
	Categories  []string  `json:"categories"`
	CreatedAt   time.Time `json:"created_at"`
}

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
