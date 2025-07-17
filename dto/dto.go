package dto

type AuthRequest struct {
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type BoxResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type CreateBoxRequest struct {
	Title string `json:"title" binding:"required,min=1,max=100"`
}
