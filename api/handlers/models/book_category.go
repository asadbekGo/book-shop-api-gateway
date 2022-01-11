package models

type BookCategory struct {
	BookId     string `json:"bookId"`
	CategoryId string `json:"categoryId"`
}

type BookCategoryDelete struct {
	CategoryId string `json:"categoryId"`
}
