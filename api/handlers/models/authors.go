package handlers

type Author struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAd"`
}

type UpdateAuthor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListAuthors struct {
	Authors []Author `json:"authors"`
	Count   int64    `json:"count"`
}
