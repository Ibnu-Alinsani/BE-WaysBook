package userdto

import "waysbook/models"

type UserResponse struct {
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Gender       string        `json:"gender"`
	Phone        string        `json:"phone"`
	Address      string        `json:"address"`
	Avatar       string        `json:"avatar"`
	Role         string        `json:"role"`
	Cart         models.Cart   `json:"cart" gorm:"foreignKey:UserId"`
	Transactions []models.Transaction `json:"transaction" gorm:"foreignKey:UserId"`
}