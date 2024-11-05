package repo

import (
	"crm-admin/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PurchasesRepo struct {
	db *sqlx.DB
}

func NewPurchasesRepo(db *sqlx.DB) *PurchasesRepo {
	return &PurchasesRepo{db: db}
}

func (p *PurchasesRepo) CreatePurchase(in entity.PurchaseRequest) (entity.Purchase, error) {
	var purchase entity.Purchase
	query := `
		INSERT INTO purchases (product_id, salesperson_id, quantity, price, total_price, description, bought_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, product_id, salesperson_id, quantity, price, total_price, description, bought_by, created_at
	`
	err := p.db.QueryRowx(query, in.ProductID, in.SalespersonID, in.Quantity, in.Price, in.TotalPrice, in.Description, in.BoughtBy).
		Scan(&purchase.ID, &purchase.ProductID, &purchase.SalespersonID, &purchase.Quantity, &purchase.Price, &purchase.TotalPrice, &purchase.Description, &purchase.BoughtBy, &purchase.CreatedAt)
	if err != nil {
		return entity.Purchase{}, fmt.Errorf("failed to create purchase: %w", err)
	}
	return purchase, nil
}

func (p *PurchasesRepo) UpdatePurchase(in entity.PurchaseRequest) (entity.Purchase, error) {
	var purchase entity.Purchase
	query := `
		UPDATE purchases
		SET quantity = $1, price = $2, total_price = $3, description = $4
		WHERE id = $5
		RETURNING id, product_id, salesperson_id, quantity, price, total_price, description, bought_by, created_at
	`
	err := p.db.QueryRowx(query, in.Quantity, in.Price, in.TotalPrice, in.Description, in.ProductID).
		Scan(&purchase.ID, &purchase.ProductID, &purchase.SalespersonID, &purchase.Quantity, &purchase.Price, &purchase.TotalPrice, &purchase.Description, &purchase.BoughtBy, &purchase.CreatedAt)
	if err != nil {
		return entity.Purchase{}, fmt.Errorf("failed to update purchase: %w", err)
	}
	return purchase, nil
}

func (p *PurchasesRepo) GetPurchase(in entity.PurchaseID) (entity.Purchase, error) {
	var purchase entity.Purchase
	query := `SELECT * FROM purchases WHERE id = $1`
	err := p.db.Get(&purchase, query, in.ID)
	if err != nil {
		return entity.Purchase{}, fmt.Errorf("failed to get purchase: %w", err)
	}
	return purchase, nil
}

func (p *PurchasesRepo) GetPurchaseList(in entity.FilterPurchase) (entity.PurchaseList, error) {
	var purchases []entity.Purchase
	query := `
		SELECT *
		FROM purchases
		WHERE ($1::UUID IS NULL OR product_id = $1)
		  AND ($2::UUID IS NULL OR salesperson_id = $2)
		  AND ($3::VARCHAR IS NULL OR bought_by = $3)
		  AND ($4::TIMESTAMP IS NULL OR created_at = $4)
	`
	err := p.db.Select(&purchases, query, in.ProductID, in.SalespersonID, in.BoughtBy, in.CreatedAt)
	if err != nil {
		return entity.PurchaseList{}, fmt.Errorf("failed to list purchases: %w", err)
	}
	return entity.PurchaseList{Purchases: purchases}, nil
}

func (p *PurchasesRepo) DeletePurchase(in entity.Purchase) (entity.Message, error) {
	query := `DELETE FROM purchases WHERE id = $1`
	res, err := p.db.Exec(query, in.ID)
	if err != nil {
		return entity.Message{}, fmt.Errorf("failed to delete purchase: %w", err)
	}
	rows, _ := res.RowsAffected()
	return entity.Message{Message: fmt.Sprintf("Deleted %d purchase(s)", rows)}, nil
}
