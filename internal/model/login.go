package model

type Login struct {
	Username string `json:"username" example:"student" binding:"required"`
	Password string `json:"password" example:"not empty" binding:"required"`
}
