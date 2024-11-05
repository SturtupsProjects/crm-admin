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
	query := `INSERT INTO product_categories (name, created_by) VALUES ($1, $2) RETURNING id, name, created_at`
	err := p.db.QueryRowx(query, in.Name, in.CreatedBy).Scan(&category.ID, &category.Name, &category.CreatedAt)
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
		INSERT INTO products (category_id, name, bill_format, incoming_price, standard_price, total_count,created_by , created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, category_id, name, bill_format, incoming_price, standard_price, total_count, created_at
	`
	err := p.db.QueryRowx(query, in.CategoryID, in.Name, in.BillFormat, in.IncomingPrice, in.StandardPrice, in.TotalCount, in.CreatedAt).
		Scan(&product.ID, &product.CategoryID, &product.Name, &product.BillFormat, &product.IncomingPrice, &product.StandardPrice, &product.TotalCount, &product.CreatedAt)
	if err != nil {
		return entity.Product{}, fmt.Errorf("failed to create product: %w", err)
	}
	return product, nil
}

func (p *ProductRepo) AddProduct(in entity.AddProductRequest) (entity.Product, error) {
	var product entity.Product

	query := `
		UPDATE products
		SET total_count = total_count + $1
		WHERE id = $2
		RETURNING id, category_id, name, bill_format, incoming_price, standard_price, total_count, created_at
	`
	err := p.db.QueryRowx(query, in.Count, in.Id).
		Scan(&product.ID, &product.CategoryID, &product.Name, &product.BillFormat, &product.IncomingPrice, &product.StandardPrice, &product.TotalCount, &product.CreatedAt)
	if err != nil {
		return entity.Product{}, fmt.Errorf("failed to add product stock: %w", err)
	}

	return product, nil
}

func (p *ProductRepo) UpdateProduct(in entity.ProductRequest) (entity.Product, error) {
	var product entity.Product
	query := `UPDATE products SET `
	var args []interface{}
	argCounter := 1

	// Dynamically build the query based on non-empty fields
	if in.CategoryID != "" {
		query += fmt.Sprintf("category_id = $%d, ", argCounter)
		args = append(args, in.CategoryID)
		argCounter++
	}
	if in.Name != "" {
		query += fmt.Sprintf("name = $%d, ", argCounter)
		args = append(args, in.Name)
		argCounter++
	}
	if in.BillFormat != "" {
		query += fmt.Sprintf("bill_format = $%d, ", argCounter)
		args = append(args, in.BillFormat)
		argCounter++
	}
	if in.IncomingPrice != 0 {
		query += fmt.Sprintf("incoming_price = $%d, ", argCounter)
		args = append(args, in.IncomingPrice)
		argCounter++
	}
	if in.StandardPrice != 0 {
		query += fmt.Sprintf("standard_price = $%d, ", argCounter)
		args = append(args, in.StandardPrice)
		argCounter++
	}
	if in.TotalCount != 0 {
		query += fmt.Sprintf("total_count = $%d, ", argCounter)
		args = append(args, in.TotalCount)
		argCounter++
	}

	// Remove trailing comma and space, add WHERE clause
	query = query[:len(query)-2] + fmt.Sprintf(" WHERE id = $%d RETURNING id, category_id, name, bill_format, incoming_price, standard_price, total_count, created_at", argCounter)
	args = append(args, in.ID)

	// Execute the query
	err := p.db.QueryRowx(query, args...).Scan(&product.ID, &product.CategoryID, &product.Name, &product.BillFormat, &product.IncomingPrice, &product.StandardPrice, &product.TotalCount, &product.CreatedAt)
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
	query := `SELECT id, category_id, name, bill_format, incoming_price, standard_price, total_count,created_by, created_at FROM products WHERE id = $1`
	err := p.db.Get(&product, query, in.ID)
	if err != nil {
		return entity.Product{}, fmt.Errorf("failed to get product: %w", err)
	}
	return product, nil
}

func (p *ProductRepo) GetProductList(in entity.FilterProduct) (entity.ProductList, error) {
	var products []entity.Product
	query := `
		SELECT id, category_id, name, bill_format, incoming_price, standard_price, total_count,created_by, created_at
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
