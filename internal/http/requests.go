package http

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserClaims struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	AccessLevel int    `json:"access_level"`
}
