package dto

type CategoryRequest struct {
	Name   string `json:"name" binding:"required"`
	UserId uint   `json:"user_id" binding:"required"`
}

type CategoryResponse struct {
	Name string `json:"name"`
	CategoryId uint `json:"category_id"`
}
