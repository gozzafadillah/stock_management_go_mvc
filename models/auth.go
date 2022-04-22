package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
}
