package handlers

type BookCategory struct {
	BookId     string `json:"bookId"`
	CategoryId string `json:"categoryId"`
}

type BookCategoryList struct {
	BookCategories []BookCategory `json:"bookCategories"`
	Count          int64          `json:"count"`
}
