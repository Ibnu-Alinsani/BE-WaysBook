package repository

import (
	// "errors"
	"errors"
	"waysbook/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllCart() ([]models.Cart, error)
	GetCartById(id int) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
	CartIdForDelete(userId int, bookId int) (models.Cart, error)
	AddCart(cart models.Cart) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllCart() ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Book").Preload("User").Find(&cart).Error

	return cart, err
}

func (r *repository) GetCartById(id int) (models.Cart, error) {
	var cart models.Cart

	err := r.db.Preload("Book").Preload("User").First(&cart, id).Error

	return cart, err
}

func (r *repository) CartIdForDelete(userId int, bookId int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("user_id = ? AND book_id = ?", userId, bookId).First(&cart).Error

	return cart, err
}

func (r *repository) AddCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Where("user_id = ? AND book_id = ?", cart.UserId, cart.BookId).First(&cart).Error
	if err == nil {
		return cart, errors.New("Book already exists in your cart")
	}

	err = r.db.Create(&cart).Error
	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}
