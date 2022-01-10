package models

type Order struct {
	BookId      string `json:"bookId"`
	Description string `json:"description"`
}

type OrderResp struct {
	BookId      string `json:"bookId"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type UpdateOrder struct {
	BookId      string `json:"bookId"`
	Description string `json:"description"`
}

type ListOrders struct {
	Orders []OrderResp `json:"orders"`
	Count  int64       `json:"count"`
}
