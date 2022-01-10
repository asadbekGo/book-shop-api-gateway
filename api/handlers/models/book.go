package models

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	AuthorId string `json:"author_id"`
}

type BookResp struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateBook struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	AuthorId string `json:"author_id"`
}

type ListBooks struct {
	Books []Book `json:"books"`
	Count int64  `json:"count"`
}
