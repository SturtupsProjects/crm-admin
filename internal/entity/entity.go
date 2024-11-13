package entity

import "time"

// ----------------- ProductCategory structs for Repo -----------------------------

type CategoryName struct {
	Name      string `json:"name" db:"name"`
	CreatedBy string `json:"created_by" db:"created_by"`
}

type CategoryID struct {
	ID string `json:"id" db:"id"`
}

type Category struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedBy string `json:"created_by" db:"created_by"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type CategoryList struct {
	Categories []Category `json:"categories"`
}

// ----------------------- Product structs for Repo -----------------------------------------

type ProductID struct {
	ID string `json:"id" db:"id"`
}

type FilterProduct struct {
	CategoryId string `json:"category_id" db:"category_id"`
	Name       string `json:"name" db:"name"`
	TotalCount string `json:"total_count" db:"total_count"`
	CreatedBy  string `json:"created_by" db:"created_by"`
}

type ProductRequest struct {
	CategoryID    string  `json:"category_id" db:"category_id"`
	Name          string  `json:"name" db:"name"`
	BillFormat    string  `json:"bill_format" db:"bill_format"`
	IncomingPrice float32 `json:"incoming_price" db:"incoming_price"`
	StandardPrice float32 `json:"standard_price" db:"standard_price"`
	CreatedBy     string  `json:"created_by" db:"created_by"`
}

type ProductUpdate struct {
	ID            string  `json:"id" db:"id"`
	CategoryID    string  `json:"category_id" db:"category_id"`
	Name          string  `json:"name" db:"name"`
	BillFormat    string  `json:"bill_format" db:"bill_format"`
	IncomingPrice float32 `json:"incoming_price" db:"incoming_price"`
	StandardPrice float32 `json:"standard_price" db:"standard_price"`
}

type ProductNumber struct {
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
	CreatedBy     string  `json:"created_by" db:"created_by"`
	CreatedAt     string  `json:"created_at" db:"created_at"`
}

type ProductList struct {
	Products []Product `json:"products"`
}

// ---------------------------------- Message ---------------------------------------------

type Message struct {
	Message string `json:"message"`
}

// -------------- PurchaseRequest is used for creating a purchase ----------------------------

type PurchaseUpdate struct {
	ID            string `json:"id" db:"id"`
	SupplierID    string `json:"supplier_id" db:"supplier_id"`
	Description   string `json:"description" db:"description"`
	PaymentMethod string `json:"payment_method" db:"payment_method"`
}

type PurchaseResponse struct {
	ID            string            `json:"id" db:"id"`
	SupplierID    string            `json:"supplier_id" db:"supplier_id"`
	PurchasedBy   string            `json:"purchased_by" db:"purchased_by"`
	TotalCost     float64           `json:"total_cost" db:"total_cost"`
	Description   string            `json:"description" db:"description"`
	PaymentMethod string            `json:"payment_method" db:"payment_method"`
	CreatedAt     string            `json:"created_at" db:"created_at"`
	PurchaseItem  []PurchaseItemReq `json:"purchase_item" db:"purchase_item"`
}

type PurchaseItemResponse struct {
	ID            string  `json:"id" db:"id"`
	PurchaseID    string  `json:"purchase_id" db:"purchase_id"`
	ProductID     string  `json:"product_id" db:"product_id"`
	Quantity      int     `json:"quantity" db:"quantity"`
	PurchasePrice float64 `json:"purchase_price" db:"purchase_price"`
	TotalPrice    float64 `json:"total_price" db:"total_price"`
}

type PurchaseRequest struct {
	SupplierID    string            `json:"supplier_id" db:"supplier_id"`
	PurchasedBy   string            `json:"purchased_by" db:"purchased_by"`
	TotalCost     float64           `json:"total_cost" db:"total_cost"`
	Description   string            `json:"description" db:"description"`
	PaymentMethod string            `json:"payment_method" db:"payment_method"`
	PurchaseItem  []PurchaseItemReq `json:"purchase_item" db:"purchase_item"`
}

type PurchaseItemReq struct {
	PurchaseID    string  `json:"purchase_id" db:"purchase_id"`
	ProductID     string  `json:"product_id" db:"product_id"`
	Quantity      int     `json:"quantity" db:"quantity"`
	PurchasePrice float64 `json:"purchase_price" db:"purchase_price"`
	TotalPrice    float64 `json:"total_price" db:"total_price"`
}

type Purchase struct {
	SupplierID    string            `json:"supplier_id" db:"supplier_id"`
	PurchasedBy   string            `json:"purchased_by" db:"purchased_by"`
	Description   string            `json:"description" db:"description"`
	PaymentMethod string            `json:"payment_method" db:"payment_method"`
	PurchaseItem  []PurchaseItemReq `json:"purchase_item" db:"purchase_item"`
}

type PurchaseItem struct {
	ProductID     string  `json:"product_id" db:"product_id"`
	Quantity      int     `json:"quantity" db:"quantity"`
	PurchasePrice float64 `json:"purchase_price" db:"purchase_price"`
}

type PurchaseID struct {
	ID string `json:"id" db:"id"`
}

type FilterPurchase struct {
	ProductID   string `json:"product_id" db:"product_id"`
	SupplierID  string `json:"salesperson_id" db:"salesperson_id"`
	PurchasedBy string `json:"bought_by" db:"bought_by"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

type PurchaseList struct {
	Purchases []PurchaseResponse `json:"purchases"`
}

// --------------- Sales structs for repo -----------------------------------------------

type SaleRequest struct {
	ClientID       string      `json:"client_id" db:"client_id"`
	SoldBy         string      `json:"sold_by" db:"sold_by"`
	TotalSalePrice float64     `json:"total_sale_price" db:"total_sale_price"`
	PaymentMethod  string      `json:"payment_method" db:"payment_method"`
	SoldProducts   []SalesItem `json:"products" db:"products"`
}

type SalesItemReq struct {
	SaleID     string  `json:"sale_id" db:"sale_id"`
	ProductID  string  `json:"product_id" db:"product_id"`
	Quantity   int     `json:"quantity" db:"quantity"`
	SalePrice  float64 `json:"sale_price" db:"sale_price"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
}

type SaleResponse struct {
	ID             string      `json:"id" db:"id"`
	ClientID       string      `json:"client_id" db:"client_id"`
	SoldBy         string      `json:"sold_by" db:"sold_by"`
	TotalSalePrice float64     `json:"total_sale_price" db:"total_sale_price"`
	PaymentMethod  string      `json:"payment_method" db:"payment_method"`
	CreatedAt      string      `json:"created_at" db:"created_at"`
	SoldProducts   []SalesItem `json:"products" db:"products"`
}

type SalesItem struct {
	ID         string  `json:"id" db:"id"`
	SaleID     string  `json:"sale_id" db:"sale_id"`
	ProductID  string  `json:"product_id" db:"product_id"`
	Quantity   int     `json:"quantity" db:"quantity"`
	SalePrice  float64 `json:"sale_price" db:"sale_price"`
	TotalPrice float64 `json:"total_price" db:"total_price"`
}

type SaleUpdate struct {
	ID            string `json:"id" db:"id"`
	ClientID      string `json:"client_id" db:"client_id"`
	PaymentMethod string `json:"payment_method" db:"payment_method"`
}

type SaleList struct {
	Sales []SaleResponse `json:"sales"`
}

type SaleID struct {
	ID string `json:"id" db:"id"`
}

type SaleFilter struct {
	StartDate string `json:"start_date" db:"start_date"`
	EndDate   string `json:"end_date" db:"end_date"`
	ClientID  string `json:"client_id" db:"client_id"`
	SoldBy    string `json:"sold_by" db:"sold_by"`
}

// -------- User structs for Repo -----------------------------------------

type User struct {
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"password"`
	Role        string `json:"role" db:"role"`
}

type UserRequest struct {
	UserID      string    `json:"user_id" db:"user_id"` // Omitted for Create
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type UserUpdate struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

type UserID struct {
	ID string `json:"id"`
}

type FilterUser struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Role      string `json:"role,omitempty"`
}

type UserList struct {
	Users []UserRequest `json:"users"`
}

type AdminPass struct {
	Login    string `json:"phone_number" db:"phone_number"`
	Password string `json:"password" db:"password"`
}

type LogIn struct {
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Password    string `json:"password" db:"password"`
}

type Token struct {
	AccessToken  string `json:"access_token" db:"access_token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	ExpireAt     int    `json:"expire_at" db:"expire_at"`
}

type LogInReq struct {
	Id          string `json:"id" db:"user_id"`
	FirstName   string `json:"first_name" db:"first_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Role        string `json:"role" db:"role"`
}

type PhoneNumber struct {
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}

type Error struct {
	Error error
}
