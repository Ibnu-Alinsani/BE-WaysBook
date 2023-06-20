package models

import "time"

type User struct {
	Id             int           `json:"id"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	Password       string        `json:"password"`
	Gender         string        `json:"gender"`
	Phone          string        `json:"phone"`
	Address        string        `json:"address"`
	Avatar         string        `json:"avatar"`
	Role           string        `json:"role"`
	CartItem       []Cart        `json:"cart_item"`
	Transaction    []Transaction `json:"transaction"`
	PublicIDAvatar string        `json:"-"`
	CreatedAt      time.Time     `json:"-"`
	UpdatedAt      time.Time     `json:"-"`
}

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (UserResponse) TableName() string {
	return "users"
}
