package repo

import (
	"crm-admin/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SalesRepo struct {
	db *sqlx.DB
}

func NewSalesRepo(db *sqlx.DB) *SalesRepo {
	return &SalesRepo{db: db}
}

func (s *SalesRepo) CreateSale(in entity.SaleRequest) (entity.Sale, error) {
	var sale entity.Sale
	query := `
		INSERT INTO sales (product_id, client_id, sale_price, quantity, total_sale_price, payment_method, sold_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, product_id, client_id, sale_price, quantity, total_sale_price, payment_method, sold_by, created_at
	`
	err := s.db.QueryRowx(query, in.ProductID, in.ClientID, in.SalePrice, in.Quantity, in.TotalSalePrice, in.PaymentMethod, in.SoldBy).
		Scan(&sale.ID, &sale.ProductID, &sale.ClientID, &sale.SalePrice, &sale.Quantity, &sale.TotalSalePrice, &sale.PaymentMethod, &sale.SoldBy, &sale.CreatedAt)
	if err != nil {
		return entity.Sale{}, fmt.Errorf("failed to create sale: %w", err)
	}
	return sale, nil
}

func (s *SalesRepo) UpdateSale(in entity.SaleRequest) (entity.Sale, error) {
	var sale entity.Sale
	query := `
		UPDATE sales
		SET sale_price = $1, quantity = $2, total_sale_price = $3, payment_method = $4
		WHERE id = $5
		RETURNING id, product_id, client_id, sale_price, quantity, total_sale_price, payment_method, sold_by, created_at
	`
	err := s.db.QueryRowx(query, in.SalePrice, in.Quantity, in.TotalSalePrice, in.PaymentMethod, in.ProductID).
		Scan(&sale.ID, &sale.ProductID, &sale.ClientID, &sale.SalePrice, &sale.Quantity, &sale.TotalSalePrice, &sale.PaymentMethod, &sale.SoldBy, &sale.CreatedAt)
	if err != nil {
		return entity.Sale{}, fmt.Errorf("failed to update sale: %w", err)
	}
	return sale, nil
}

func (s *SalesRepo) GetSale(in entity.Sale) (entity.Sale, error) {
	var sale entity.Sale
	query := `SELECT * FROM sales WHERE id = $1`
	err := s.db.Get(&sale, query, in.ID)
	if err != nil {
		return entity.Sale{}, fmt.Errorf("failed to get sale: %w", err)
	}
	return sale, nil
}

func (s *SalesRepo) GetSaleList() (entity.SaleList, error) {
	var sales []entity.Sale
	query := `SELECT * FROM sales`
	err := s.db.Select(&sales, query)
	if err != nil {
		return entity.SaleList{}, fmt.Errorf("failed to list sales: %w", err)
	}
	return entity.SaleList{Sales: sales}, nil
}

func (s *SalesRepo) DeleteSale(in entity.Sale) (entity.Message, error) {
	query := `DELETE FROM sales WHERE id = $1`
	res, err := s.db.Exec(query, in.ID)
	if err != nil {
		return entity.Message{}, fmt.Errorf("failed to delete sale: %w", err)
	}
	rows, _ := res.RowsAffected()
	return entity.Message{Message: fmt.Sprintf("Deleted %d sale(s)", rows)}, nil
}
