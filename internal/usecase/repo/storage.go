package repo

import (
	"crm-admin/internal/entity"
	"crm-admin/internal/usecase"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) usecase.ProductsRepo {
	return &ProductRepo{db: db}
}

// ---------------- Product Category CRUD -----------------------------------------------------------------------------

func (p *ProductRepo) CreateProductCategory(in entity.CategoryName) (entity.Category, error) {
	return entity.Category{}, nil
}

func (p *ProductRepo) DeleteProductCategory(in entity.CategoryID) (entity.Message, error) {
	return entity.Message{}, nil
}

func (p *ProductRepo) GetProductCategory(in entity.CategoryID) (entity.Category, error) {
	return entity.Category{}, nil
}

func (p *ProductRepo) GetListProductCategory(in entity.CategoryName) (entity.CategoryList, error) {
	return entity.CategoryList{}, nil
}

// ------------------- Product CRUD ------------------------------------------------------------------------

func (p *ProductRepo) CreateProduct(in entity.ProductRequest) (entity.Product, error) {
	return entity.Product{}, nil
}

func (p *ProductRepo) UpdateProduct(in entity.ProductRequest) (entity.Product, error) {
	return entity.Product{}, nil
}

func (p *ProductRepo) DeleteProduct(in entity.Product) (entity.Message, error) {
	return entity.Message{}, nil
}

func (p *ProductRepo) GetProduct(in entity.ProductID) (entity.Product, error) {
	return entity.Product{}, nil
}

func (p *ProductRepo) GetProductList(in entity.FilterProduct) (entity.ProductList, error) {
	return entity.ProductList{}, nil
}
