package entity

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type AuthResponse struct {
    UserID int    `json:"user_id"`
    Email  string `json:"email"`
    Token  string `json:"token"`
}