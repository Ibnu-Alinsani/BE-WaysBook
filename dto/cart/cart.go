package dtocart 

type AddCart struct {
	BookId int `json:"book_id"`
}

type DeleteCart struct {
	BookId int `json:"book_id"`
}