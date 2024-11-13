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

func (s *SalesRepo) CreateSale(in entity.SaleRequest) (entity.SaleResponse, error) {
	var sale entity.SaleResponse
	query := `
		INSERT INTO sales (client_id, sold_by, total_sale_price, payment_method)
		VALUES ($1, $2, $3, $4)
		RETURNING id, client_id, sold_by, total_sale_price, payment_method, created_at
	`
	err := s.db.QueryRowx(query, in.ClientID, in.SoldBy, in.TotalSalePrice, in.PaymentMethod).
		Scan(&sale.ID, &sale.ClientID, &sale.SoldBy, &sale.TotalSalePrice, &sale.PaymentMethod, &sale.CreatedAt)
	if err != nil {
		return entity.SaleResponse{}, fmt.Errorf("failed to create sale: %w", err)
	}
	return sale, nil
}

func (s *SalesRepo) UpdateSale(in entity.SaleRequest) (entity.SaleResponse, error) {
	var sale entity.SaleResponse
	query := `
		UPDATE sales
		SET total_sale_price = $1, payment_method = $2
		WHERE id = $3
		RETURNING id, client_id, sold_by, total_sale_price, payment_method, created_at
	`
	err := s.db.QueryRowx(query, in.TotalSalePrice, in.PaymentMethod, in.Id).
		Scan(&sale.ID, &sale.ClientID, &sale.SoldBy, &sale.TotalSalePrice, &sale.PaymentMethod, &sale.CreatedAt)
	if err != nil {
		return entity.SaleResponse{}, fmt.Errorf("failed to update sale: %w", err)
	}
	return sale, nil
}

func (s *SalesRepo) GetSale(in entity.SaleID) (entity.SaleResponse, error) {
	var sale entity.SaleResponse
	query := `SELECT * FROM sales WHERE id = $1`
	err := s.db.Get(&sale, query, in.ID)
	if err != nil {
		return entity.SaleResponse{}, fmt.Errorf("failed to get sale: %w", err)
	}
	return sale, nil
}

func (s *SalesRepo) GetSaleList(filter entity.SaleFilter) (entity.SaleList, error) {
	var sales []entity.SaleResponse
	query := `SELECT * FROM sales`
	if filter.StartDate != "" || filter.EndDate != "" || filter.ClientID != "" {
		query += ` WHERE `
		if filter.StartDate != "" {
			query += `created_at >= $1`
		}
		if filter.EndDate != "" {
			if filter.StartDate != "" {
				query += ` AND `
			}
			query += `created_at <= $2`
		}
		if filter.ClientID != "" {
			if filter.StartDate != "" || filter.EndDate != "" {
				query += ` AND `
			}
			query += `client_id = $3`
		}
	}
	err := s.db.Select(&sales, query, filter.StartDate, filter.EndDate, filter.ClientID)
	if err != nil {
		return entity.SaleList{}, fmt.Errorf("failed to list sales: %w", err)
	}
	return entity.SaleList{Sales: sales}, nil
}

func (s *SalesRepo) DeleteSale(in entity.SaleID) (entity.Message, error) {
	query := `DELETE FROM sales WHERE id = $1`
	res, err := s.db.Exec(query, in.ID)
	if err != nil {
		return entity.Message{}, fmt.Errorf("failed to delete sale: %w", err)
	}
	rows, _ := res.RowsAffected()
	return entity.Message{Message: fmt.Sprintf("Deleted %d sale(s)", rows)}, nil
}
