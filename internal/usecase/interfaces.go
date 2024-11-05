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
	DeleteProduct(in entity.Product) (entity.Message, error)
	GetProduct(in entity.ProductID) (entity.Product, error)
	GetProductList(in entity.FilterProduct) (entity.ProductList, error)
}

type UsersRepo interface {
	CreateUser(in entity.UserRequest) (entity.User, error)
	GetUser(in entity.UserID) (entity.User, error)
	GetListUser(in entity.FilterUser) (entity.UserList, error)
	DeleteUser(in entity.UserID) (entity.Message, error)
	UpdateUser(in entity.UserRequest) (entity.User, error)
}
