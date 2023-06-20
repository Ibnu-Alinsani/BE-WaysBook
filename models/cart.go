package models

type Cart struct {
	Id     int          `json:"id"`
	UserId int          `json:"user_id"`
	User   UserResponse `json:"user" gorm:"foreignKey:UserId"`
	BookId int          `json:"book_id"`
	Book   Book `json:"book" gorm:"foreignKey:BookId"`
}
