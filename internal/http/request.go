package http

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserMetaRequest struct {
	Banned      bool `json:"banned"`
	Active      bool `json:"active"`
	AccessLevel int  `json:"access_level"`
}

type UserClaims struct {
	Username    string `json:"username"`
	IsAdmin     bool   `json:"is_admin"`
	Banned      bool   `json:"banned"`
	Active      bool   `json:"active"`
	AccessLevel int    `json:"access_level"`
}

type CategoryRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
