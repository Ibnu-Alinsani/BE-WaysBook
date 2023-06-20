package models

import "time"

type Book struct {
	Id                int           `json:"id"`
	Title             string        `json:"title"`
	PublicationDate   string        `json:"publication_date"`
	Pages             string        `json:"pages"`
	ISBN              string        `json:"isbn"`
	Author            string        `json:"author"`
	Discount          int           `json:"discount"`
	Price             int           `json:"price"`
	DiscountAmount    int           `json:"discount_amount"`
	Description       string        `json:"description"`
	BookAttachment    string        `json:"book_attachment"`
	Thumbnail         string        `json:"thumbnail"`
	Transaction       []Transaction `json:"transaction" gorm:"many2many:transaction_books"`
	PublicIdBook      string        `json:"-"`
	PublicIdThumbnail string        `json:"-"`
	CreatedAt         time.Time     `json:"-"`
	UpdatedAt         time.Time     `json:"-"`
}

type BookResponse struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publication_date"`
	Pages           string `json:"pages"`
	ISBN            string `json:"isbn"`
	Author          string `json:"author"`
	DiscountAmount  string `json:"discount_amount"`
	Description     string `json:"description"`
	BookAttachment  string `json:"book_attachment"`
	Thumbnail       string `json:"thumbnail"`
}

func (BookResponse) TableName() string {
	return "books"
}
