package dtotransaction

type TransactionRequest struct {
	Id           int    `json:"id" validation:"required"`
	UserId       int    `json:"user_id" validation:"required"`
	BookId       int    `json:"book_id" validation:"required"`
	CounterQty   int    `json:"counter_qty" validation:"required"`
	TotalPayment int `json:"total_payment" validation:"required"`
	Status       string `json:"status" validation:"required"`
}
