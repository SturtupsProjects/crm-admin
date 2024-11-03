package entity

// ProductCategory structs for Repo -----------------------------

type CategoryName struct {
	Name string `json:"name" db:"name"`
}

type CategoryID struct {
	ID string `json:"id" db:"id"`
}

type Category struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type CategoryList struct {
	Categories []Category `json:"categories"`
}

// Product structs for Repo -----------------------------------------

type ProductID struct {
	ID string `json:"id" db:"id"`
}

type FilterProduct struct {
	CategoryId string `json:"category_id" db:"category_id"`
	Name       string `json:"name" db:"name"`
	TotalCount string `json:"total_count" db:"total_count"`
	CreatedAt  string `json:"created_at" db:"created_at"`
}

type ProductRequest struct {
	CategoryID string  `json:"category_id" db:"category_id"`
	Name       string  `json:"name" db:"name"`
	BillFormat string  `json:"bill_format" db:"bill_format"`
	BasicPrice float32 `json:"basic_price" db:"basic_price"`
}

type Product struct {
	ID         string  `json:"id" db:"id"`
	CategoryID string  `json:"category_id" db:"category_id"`
	Name       string  `json:"name" db:"name"`
	BillFormat string  `json:"bill_format" db:"bill_format"`
	BasicPrice float32 `json:"basic_price" db:"basic_price"`
	TotalCount int     `json:"total_count" db:"total_count"`
	CreatedAt  string  `json:"created_at" db:"created_at"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

// Message ---------------------------------------------

type Message struct {
	Message string `json:"message"`
}
