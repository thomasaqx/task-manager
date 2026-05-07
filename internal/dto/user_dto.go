package dto

type UserRequest struct{
	Name string `json:"name" binding:"required"`
	Email uint `json:"email" binding:"required"`
	Cpf int64 `json:"cpf" binding:"required"`
}

type UserResponse struct {
	Id uint `json:"id"`
}
