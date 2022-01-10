package models

type BookCategory struct {
	BookId     string `json:"book_id"`
	CategoryId string `json:"category_id"`
}

type BookCategoryResp struct {
	Book     BookResp   `json:"book_resp"`
	Category []Category `json:"categories"`
}

type BookCategoryList struct {
	BookCategories []BookCategoryResp `json:"book_categories"`
	Count          int64              `json:"count"`
}
