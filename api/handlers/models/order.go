package models

type Order struct {
	BookId      string `json:"book_id"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateOrder struct {
	ID          string `json:"id"`
	BookId      string `json:"book_id"`
	Description string `json:"description"`
}

type ListOrders struct {
	Orders []Order `json:"orders"`
	Count  int64   `json:"count"`
}
