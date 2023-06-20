package models

import (
	"time"
)

type Transaction struct {
	Id           int       `json:"id"`
	CounterQty   int       `json:"counter_qty"`
	TotalPayment int       `json:"total_payment"`
	Status       string    `json:"status"`
	UserId       int       `json:"user_id"`
	User         User      `json:"user" gorm:"foreignKey:UserId"`
	Book         []Book    `json:"book" gorm:"many2many:transaction_books"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type TransactionResponse struct {
	Id           int          `json:"id"`
	UserId       int          `json:"user_id"`
	BookId       int          `json:"book_id"`
	Book         BookResponse `json:"book" gorm:"foreignKey:BookId"`
	CounterQty   int          `json:"counter_qty"`
	TotalPayment string       `json:"total_payment"`
	Status       string       `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
