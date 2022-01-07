package handlers

type Order struct {
	ID          string `json:"id"`
	BookId      string `json:"bookId"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type UpdateOrder struct {
	ID          string `json:"id"`
	BookId      string `json:"bookId"`
	Description string `json:"description"`
}

type ListOrders struct {
	Orders []Order `json:"orders"`
	Count  int64   `json:"count"`
}
