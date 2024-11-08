package usecase

import "crm-admin/internal/entity"

type ProductsRepo interface {
	CreateProductCategory(in entity.CategoryName) (entity.Category, error)
	DeleteProductCategory(in entity.CategoryID) (entity.Message, error)
	GetProductCategory(in entity.CategoryID) (entity.Category, error)
	GetListProductCategory(in entity.CategoryName) (entity.CategoryList, error)

	CreateProduct(in entity.ProductRequest) (entity.Product, error)
	AddProduct(in entity.AddProductRequest) (entity.Product, error)
	UpdateProduct(in entity.ProductRequest) (entity.Product, error)
	DeleteProduct(in entity.ProductID) (entity.Message, error)
	GetProduct(in entity.ProductID) (entity.Product, error)
	GetProductList(in entity.FilterProduct) (entity.ProductList, error)
}

type UsersRepo interface {
	AddAdmin(in entity.AdminPass) (entity.Message, error)
	CreateUser(in entity.User) (entity.UserRequest, error)
	GetUser(in entity.UserID) (entity.UserRequest, error)
	GetListUser(in entity.FilterUser) (entity.UserList, error)
	DeleteUser(in entity.UserID) (entity.Message, error)
	UpdateUser(in entity.UserRequest) (entity.UserRequest, error)
	LogIn(in entity.PhoneNumber) (entity.LogInReq, error)
}

type PurchasesRepo interface {
	CreatePurchase(in entity.PurchaseRequest) (entity.Purchase, error)
	UpdatePurchase(in entity.PurchaseRequest) (entity.Purchase, error)
	GetPurchase(in entity.PurchaseID) (entity.Purchase, error)
	GetPurchaseList(in entity.FilterPurchase) (entity.PurchaseList, error)
	DeletePurchase(in entity.Purchase) (entity.Message, error)
}

type SalesRepo interface {
	CreateSale(in entity.SaleRequest) (entity.SaleResponse, error)
	UpdateSale(in entity.SaleRequest) (entity.SaleResponse, error)
	GetSale(in entity.SaleID) (entity.SaleResponse, error)
	GetSaleList(filter entity.SaleFilter) (entity.SaleList, error)
	DeleteSale(in entity.SaleID) (entity.Message, error)
}

type ReturnedProductsRepo interface {
	CreateReturnedProducts() error
	UpdateReturnedProducts() error
	GetReturnedProducts() error
	GetReturnedProductsList() error
	DeleteReturnedProducts() error
}
