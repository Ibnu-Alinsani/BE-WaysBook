package repository

import (
	// "os/user"
	"fmt"
	"waysbook/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransaction() ([]models.Transaction, error)
	GetTransactionById(id int) (models.Transaction, error)
	AddTransaction(cart models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
	DeleteTransaction(userId int) (models.Cart, error)

	FindBooks(bookId int) ([]models.Book, error)
	GetUser(userId int) (models.User, error)
	UpdateUserCart(user models.User) (models.User, error)
	Delete(id int) error
	GetBookId(id int) (models.Book, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllTransaction() ([]models.Transaction, error) {
	var transac []models.Transaction

	err := r.db.Preload("User").Preload("Book").Find(&transac).Error

	return transac, err
}

func (r *repository) GetTransactionById(id int) (models.Transaction, error) {
	var transac models.Transaction

	err := r.db.Preload("User").Preload("Book").First(&transac, id).Error

	return transac, err
}

func (r *repository) AddTransaction(transac models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transac).Error

	return transac, err
}

func (r *repository) UpdateTransaction(status string, id int) (models.Transaction, error) {
	var transac models.Transaction

	r.db.Preload("Book").First(&transac, id)

	transac.Status = status
	fmt.Println(transac, "ini transaction")
	err := r.db.Save(&transac).Error

	return transac, err
}

func (r *repository) DeleteTransaction(userId int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Exec("DELETE FROM carts where user_id = ?", userId).Scan(&cart).Error
	return cart, err
}

func (r *repository) GetUser(userId int) (models.User, error) {
	var user models.User

	err := r.db.Preload("CartItem.User").Preload("CartItem.Book").Find(&user, userId).Error

	return user, err
}

func (r *repository) FindBooks(bookId int) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books, bookId).Error

	return books, err
}

func (r *repository) GetBookId(id int) (models.Book, error) {
	var book models.Book
	err := r.db.Preload("Transaction").First(&book, id).Error

	return book, err
}

func (r *repository) UpdateUserCart(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) Delete(id int) error {
	err := r.db.Exec("DELETE FROM carts WHERE user_id = ?", id).Error

	return err
}
