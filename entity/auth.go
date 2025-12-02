package entity

type LoginRequest struct {
	Email    string
	Password string
}

type AuthResponse struct {
	UserID int
	Email  string
	Token  string
}
