package dto

type AuthRequest struct {
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
