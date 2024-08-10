package models

type Category struct {
	CatID uint64 `json:"cat_id"`
	CategoryAttributes
}

type CategoryAttributes struct {
	CatName string `json:"cat_name"`
}

type CategoryRepo interface {
	Create(category CategoryAttributes) error
	GetByCategoryID(categoryID uint64) (*Category, error)
	DeleteByCategoryID(categoryID uint64) error
}
