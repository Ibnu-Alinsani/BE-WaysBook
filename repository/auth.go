package repository

import (
	"errors"
	"waysbook/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(id int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	var count int
	err := r.db.Raw("SELECT COUNT(*) from users WHERE email=?", user.Email).Scan(&count).Error

	if err != nil {
		return user, err
	}

	if count < 1 {
		err = r.db.Create(&user).Error
	} else {
		err = errors.New("email already exists")
	}

	return user, err
}

func (r *repository) Login(email string) (models.User, error) {
	var user models.User

	err := r.db.First(&user, "email = ?", email).Error

	return user, err
}

func (r *repository) CheckAuth(id int) (models.User, error) {
	var user models.User
	err := r.db.Preload("CartItem").Preload("Transaction.Book").First(&user, id).Error

	return user, err
}
