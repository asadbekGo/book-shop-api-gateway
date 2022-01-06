package handlers

type Order struct {
	ID          string `json:"id"`
	BookId      string `json:"book_id"`
	Quantity    string `json:"quantity"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateOrder struct {
	ID          string `json:"id"`
	BookId      string `json:"book_id"`
	Quantity    string `json:"quantity"`
	Description string `json:"description"`
}

type ListOrders struct {
	Orders []Order `json:"orders"`
}
