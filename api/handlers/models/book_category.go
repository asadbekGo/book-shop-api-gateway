package handlers

type BookCategory struct{
	BookId 		string `json:"bookId"`
	CategoryId 	string `json:"category"`
}

type BookCategoryList struct{
	BookCategories []BookCategory `json:"bookCategories"`
}