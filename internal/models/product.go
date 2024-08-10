package models

type Product struct {
	ProductID uint64 `json:"product_id"`
	ProductAttributes
}

type ProductAttributes struct {
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	RefPrice           uint64 `json:"ref_price"`
	UnixTimeCreated    uint64 `json:"unix_time_created"`
	CreatorID          uint64 `json:"creator_id"`
}

type ProductSearchSortBy int

const (
	Time ProductSearchSortBy = iota
	Price
)

type ProductRepo interface {
	Create(product ProductAttributes) error
	GetByProductID(productID uint64) (*Product, error)
	GetPageByCategoryID(categoryID uint64, pageNumber int, sortBy ProductSearchSortBy) ([]Product, error)
	GetPageByProductNameSearch(productName string, pageNumber int, sortBy ProductSearchSortBy) ([]Product, error)
	DeleteByProductID(productID uint64) error
}
