package handlers

type Book struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AuthorId  string `json:"authorId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAd"`
}

type UpdateBook struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AuthorId  string `json:"authorId"`
}

type ListBooks struct {
	Books []Book `json:"books"`
}
