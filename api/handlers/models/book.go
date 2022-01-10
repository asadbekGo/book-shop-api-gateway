package models

type Book struct {
	Name     string `json:"name"`
	AuthorId string `json:"authorId"`
}

type BookResp struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Author    AuthorResp `json:"author"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type UpdateBook struct {
	Name     string `json:"name"`
	AuthorId string `json:"author_id"`
}

type ListBooks struct {
	Books []BookResp `json:"books"`
	Count int64      `json:"count"`
}
