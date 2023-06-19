package repository

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUser() ([]models.User, error) {
	var user []models.User 

	err := r.db.Preload("Transaction").Preload("CartItem.User").Preload("CartItem.Book").Find(&user).Error

	return user, err
}

func (r *repository) GetUserById(id int) (models.User, error) {
	var user models.User

	err := r.db.Preload("Transaction").Preload("Transaction.Book").Preload("CartItem.User").Preload("CartItem.Book").First(&user, id).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}