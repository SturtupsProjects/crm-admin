package repo

import (
	"crm-admin/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

// ---------------- Product Category CRUD -----------------------------------------------------------------------------

func (p *ProductRepo) CreateProductCategory(in entity.CategoryName) (entity.Category, error) {
	var category entity.Category
	query := `INSERT INTO product_categories (name) VALUES ($1) RETURNING id, name, created_at`
	err := p.db.QueryRowx(query, in.Name).Scan(&category.ID, &category.Name, &category.CreatedAt)
	if err != nil {
		return entity.Category{}, fmt.Errorf("failed to create product category: %w", err)
	}
	return category, nil
}

func (p *ProductRepo) DeleteProductCategory(in entity.CategoryID) (entity.Message, error) {
	query := `DELETE FROM product_categories WHERE id = $1`
	res, err := p.db.Exec(query, in.ID)
	if err != nil {
		return entity.Message{}, fmt.Errorf("failed to delete product category: %w", err)
	}
	rows, _ := res.RowsAffected()
	return entity.Message{Message: fmt.Sprintf("Deleted %d category(ies)", rows)}, nil
}

func (p *ProductRepo) GetProductCategory(in entity.CategoryID) (entity.Category, error) {
	var category entity.Category
	query := `SELECT id, name, created_at FROM product_categories WHERE id = $1`
	err := p.db.Get(&category, query, in.ID)
	if err != nil {
		return entity.Category{}, fmt.Errorf("failed to get product category: %w", err)
	}
	return category, nil
}

func (p *ProductRepo) GetListProductCategory() (entity.CategoryList, error) {
	var categories []entity.Category
	query := `SELECT id, name, created_at FROM product_categories`
	err := p.db.Select(&categories, query)
	if err != nil {
		return entity.CategoryList{}, fmt.Errorf("failed to list product categories: %w", err)
	}
	return entity.CategoryList{Categories: categories}, nil
}

// ---------------- End Product Category CRUD ------------------------------------------------------------------------

// ------------------- Product CRUD ------------------------------------------------------------------------

func (p *ProductRepo) CreateProduct(in entity.ProductRequest) (entity.Product, error) {
	var product entity.Product
	query := `
		INSERT INTO products (category_id, name, bill_format, base_price, total_count)
		VALUES ($1, $2, $3, $4, 0)
		RETURNING id, category_id, name, bill_format, base_price, total_count, created_at
	`
	err := p.db.QueryRowx(query, in.CategoryID, in.Name, in.BillFormat, in.BasicPrice).
		Scan(&product.ID, &product.CategoryID, &product.Name, &product.BillFormat, &product.BasicPrice, &product.TotalCount, &product.CreatedAt)
	if err != nil {
		return entity.Product{}, fmt.Errorf("failed to create product: %w", err)
	}
	return product, nil
}

func (p *ProductRepo) UpdateProduct(in entity.ProductRequest) (entity.Product, error) {
	var product entity.Product
	query := `
		UPDATE products 
		SET name = $1, bill_format = $2, base_price = $3, category_id = $4
		WHERE id = $5
		RETURNING id, category_id, name, bill_format, base_price, total_count, created_at
	`
	err := p.db.QueryRowx(query, in.Name, in.BillFormat, in.BasicPrice, in.CategoryID, in.ID).
		Scan(&product.ID, &product.CategoryID, &product.Name, &product.BillFormat, &product.BasicPrice, &product.TotalCount, &product.CreatedAt)
	if err != nil {
		return entity.Product{}, fmt.Errorf("failed to update product: %w", err)
	}
	return product, nil
}

func (p *ProductRepo) DeleteProduct(in entity.ProductID) (entity.Message, error) {
	query := `DELETE FROM products WHERE id = $1`
	res, err := p.db.Exec(query, in.ID)
	if err != nil {
		return entity.Message{}, fmt.Errorf("failed to delete product: %w", err)
	}
	rows, _ := res.RowsAffected()
	return entity.Message{Message: fmt.Sprintf("Deleted %d product(s)", rows)}, nil
}

func (p *ProductRepo) GetProduct(in entity.ProductID) (entity.Product, error) {
	var product entity.Product
	query := `SELECT id, category_id, name, bill_format, base_price, total_count, created_at FROM products WHERE id = $1`
	err := p.db.Get(&product, query, in.ID)
	if err != nil {
		return entity.Product{}, fmt.Errorf("failed to get product: %w", err)
	}
	return product, nil
}

func (p *ProductRepo) GetProductList(in entity.FilterProduct) (entity.ProductList, error) {
	var products []entity.Product
	query := `
		SELECT id, category_id, name, bill_format, base_price, total_count, created_at
		FROM products 
		WHERE ($1::UUID IS NULL OR category_id = $1) 
		  AND ($2::VARCHAR IS NULL OR name ILIKE '%' || $2 || '%')
		ORDER BY created_at DESC
	`
	err := p.db.Select(&products, query, in.CategoryId, in.Name)
	if err != nil {
		return entity.ProductList{}, fmt.Errorf("failed to list products: %w", err)
	}
	return entity.ProductList{Products: products}, nil
}

// ------------------- End Product CRUD ------------------------------------------------------------------------
