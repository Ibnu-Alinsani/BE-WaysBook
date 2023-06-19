package repository

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBook() ([]models.Book, error)
	GetBookById(id int) (models.Book, error)
	AddBook(book models.Book) (models.Book, error)
	DeleteBook(book models.Book) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
}

func RepositoryBook(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllBook() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Preload("Transaction").Find(&books).Error

	return books, err
}

func (r *repository) GetBookById(id int) (models.Book, error) {
	var book models.Book
	err := r.db.Preload("Transaction").First(&book, id).Error

	return book, err
}

func (r *repository) AddBook(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) DeleteBook(book models.Book) (models.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}

func (r *repository) UpdateBook(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}
