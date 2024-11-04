package entity

import "time"

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
	ID            string  `json:"id" db:"id"`
	CategoryID    string  `json:"category_id" db:"category_id"`
	Name          string  `json:"name" db:"name"`
	BillFormat    string  `json:"bill_format" db:"bill_format"`
	IncomingPrice float32 `json:"incoming_price" db:"incoming_price"`
	StandardPrice float32 `json:"standard_price" db:"standard_price"`
	TotalCount    int     `json:"total_count" db:"total_count"`
	CreatedAt     string  `json:"created_at" db:"created_at"`
}
type AddProductRequest struct {
	Id    string `json:"id" db:"id"`
	Count int    `json:"count" db:"count"`
}

type Product struct {
	ID            string  `json:"id" db:"id"`
	CategoryID    string  `json:"category_id" db:"category_id"`
	Name          string  `json:"name" db:"name"`
	BillFormat    string  `json:"bill_format" db:"bill_format"`
	IncomingPrice float32 `json:"incoming_price" db:"incoming_price"`
	StandardPrice float32 `json:"standard_price" db:"standard_price"`
	TotalCount    int     `json:"total_count" db:"total_count"`
	CreatedAt     string  `json:"created_at" db:"created_at"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

// Message ---------------------------------------------

type Message struct {
	Message string `json:"message"`
}

type User struct {
	UserID      string    `json:"user_id" db:"user_id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// UserRequest is used for creating or updating a user.
type UserRequest struct {
	UserID      string `json:"user_id,omitempty"` // Omitted for Create
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

// UserID is used for identifying a user in retrieval or deletion operations.
type UserID struct {
	ID string `json:"id"`
}

// FilterUser is used for filtering users in list operations.
type FilterUser struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Role      string `json:"role,omitempty"`
}

// UserList represents a list of users, used in GetListUser responses.
type UserList struct {
	Users []User `json:"users"`
}
