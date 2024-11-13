package usecase

import (
	"crm-admin/internal/entity"
	"log/slog"
)

type PurchaseUseCase struct {
	repo    PurchasesRepo
	product ProductQuantity
	log     *slog.Logger
}

func NewPurchaseUseCase(repo PurchasesRepo, pr ProductQuantity, log *slog.Logger) *PurchaseUseCase {
	return &PurchaseUseCase{
		repo:    repo,
		product: pr,
		log:     log,
	}
}

func (p *PurchaseUseCase) CalculateTotalPurchases(in *entity.Purchase) (*entity.PurchaseRequest, error) {
	var result *entity.PurchaseRequest
	var totalSum float64

	var purchaseList []entity.PurchaseItemReq
	var purchase entity.PurchaseItemReq

	for _, pr := range *in.PurchaseItem {
		if pr.Quantity == 0 {
			continue
		}

		sum := float64(pr.Quantity) * pr.PurchasePrice

		purchase.PurchasePrice = pr.PurchasePrice
		purchase.ProductID = pr.ProductID
		purchase.Quantity = pr.Quantity
		purchase.TotalPrice = sum

		purchaseList = append(purchaseList, purchase)
		totalSum += sum
	}

	result.PurchasedBy = in.PurchasedBy
	result.SupplierID = in.SupplierID
	result.PurchaseItem = &purchaseList
	result.TotalCost = totalSum
	result.PaymentMethod = in.PaymentMethod
	result.Description = in.Description

	return result, nil
}

func (p *PurchaseUseCase) CreatePurchase(in *entity.Purchase) (*entity.PurchaseResponse, error) {

	req, err := p.CalculateTotalPurchases(in)
	if err != nil {
		p.log.Error("Error in calculating", "error", err.Error())
		return nil, err
	}

	res, err := p.repo.CreatePurchase(req)
	if err != nil {
		p.log.Error("Error in creating purchase", "error", err.Error())
		return nil, err
	}

	// You can use THE goroutines ----------------------------
	for _, item := range *res.PurchaseItem {
		productQuantityReq := &entity.UpdateProductNumber{
			Id:    item.ProductID,
			Count: item.Quantity,
		}
		_, err := p.product.AddProduct(productQuantityReq)
		if err != nil {
			p.log.Error("Error in adding product", "error", err.Error())
			return nil, err
		}
	}

	return res, nil
}
