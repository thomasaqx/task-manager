package model

type Category struct{
	Id uint `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	UserId uint `json:"user_id" db:"user_id"`
}