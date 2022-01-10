package models

type Author struct {
	Name string `json:"name"`
}

type AuthorResp struct {
	AuthorId  string `json:"authorId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UpdateAuthor struct {
	Name string `json:"name"`
}

type ListAuthors struct {
	Authors []AuthorResp `json:"authors"`
	Count   int64        `json:"count"`
}

type Empty struct{}
