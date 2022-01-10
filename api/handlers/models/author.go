package models

type Author struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateAuthor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListAuthors struct {
	Authors []Author `json:"authors"`
	Count   int64    `json:"count"`
}
