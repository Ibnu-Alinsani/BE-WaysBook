package dtobook

type BookRequest struct {
	Title           string `json:"string" form:"title"`
	PublicationDate string `json:"publication_date" form:"publication_date"`
	Pages           string `json:"pages" form:"pages"`
	ISBN            string `json:"isbn" form:"isbn"`
	Author          string `json:"author" form:"author"`
	Discount        string `json:"discount" form:"discount"`
	Price           string `json:"price" form:"price"`
	DiscountAmount  string `json:"discount_amount" form:"discount_amount" `
	Description     string `json:"description" form:"description"`
	BookAttachment  string `json:"book_attachment" form:"book_attachment"`
	Thumbnail       string `json:"thumbnail" form:"thumbnail"`
}