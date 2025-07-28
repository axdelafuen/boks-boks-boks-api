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

type LabelResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

type CreateBoxRequest struct {
	Title string `json:"title" binding:"required,min=1,max=100"`
}

type ItemResponse struct {
	Id     string          `json:"id"`
	Title  string          `json:"title"`
	Amount int             `json:"amount"`
	Labels []LabelResponse `json:"labels"`
}

type CreateItemRequest struct {
	Title  string          `json:"title" binding:"required,min=1,max=100"`
	Amount int             `json:"amount" binding:"required,min=1"`
	Labels []LabelResponse `json:"labels"`
}

type UpdateBoxRequest struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type UpdateItemRequest struct {
	Id     string          `json:"id"`
	Title  string          `json:"title"`
	Amount int             `json:"amount"`
	Labels []LabelResponse `json:"labels"`
}

type CreateLabelRequest struct {
	Title       string `json:"title"`
	Color       string `json:"color"`
	Description string `json:"description"`
}
