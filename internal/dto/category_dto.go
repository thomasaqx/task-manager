package dto

type CategoryRequest struct {
	Name   string `json:"name" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
}

type CategoryResponse struct {
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}
