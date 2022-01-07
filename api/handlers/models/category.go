package handlers

type Category struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ParentUUID string `json:"parrentUUID"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type UpdateCategory struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ParentUUID string `json:"parrentUUID"`
}

type ListCategories struct {
	Categories []Category `json:"categories"`
	Count      int64      `json:"count"`
}
