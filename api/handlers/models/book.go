package models

type Book struct {
	Name     string   `json:"name"`
	AuthorId string   `json:"authorId"`
	Category []string `json:"categories"`
}

type BookResp struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Author     AuthorResp     `json:"author"`
	Categories []CategoryResp `json:"categories"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
}

type UpdateBook struct {
	Name     string   `json:"name"`
	AuthorId string   `json:"authorId"`
	Category []string `json:"categories"`
}

type ListBooks struct {
	Books []BookResp `json:"books"`
	Count int64      `json:"count"`
}
