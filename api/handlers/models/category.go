package models

type Category struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ParentUUID string `json:"parrentUUID"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
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
