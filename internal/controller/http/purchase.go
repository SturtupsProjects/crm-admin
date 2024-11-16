package http

import (
	"crm-admin/internal/entity"
	"crm-admin/internal/usecase"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type purchaseRoutes struct {
	useCase *usecase.PurchaseUseCase
	log     *slog.Logger
}

func newPurchaseRoutes(router *gin.RouterGroup, us *usecase.PurchaseUseCase, log *slog.Logger) {
	purchase := &purchaseRoutes{useCase: us, log: log}

	// ------------ purchase router ------------------
	router.POST("", purchase.CreatePurchase)
	router.PUT("/:id", purchase.UpdatePurchase)
	router.GET("/:id", purchase.GetPurchase)
	router.GET("", purchase.GetListPurchase)
	router.DELETE("/:id", purchase.DeletePurchase)
}

// CreatePurchase godoc
// @Summary Create Purchase
// @Description Create a new purchase
// @Tags Purchase
// @Accept json
// @Produce json
// @Param Purchase body entity.Purchase true "Purchase data"
// @Success 201 {object} entity.PurchaseResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /purchases [post]
func (p *purchaseRoutes) CreatePurchase(c *gin.Context) {
	var req entity.Purchase

	if err := c.ShouldBindJSON(&req); err != nil {
		p.log.Error("Error binding JSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := p.useCase.CreatePurchase(&req)
	if err != nil {
		p.log.Error("Error creating purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdatePurchase godoc
// @Summary Update Purchase
// @Description Update purchase details by ID
// @Tags Purchase
// @Accept json
// @Produce json
// @Param id path string true "Purchase ID"
// @Param PurchaseUpdate body entity.PurchaseUpdate true "Updated purchase data"
// @Success 200 {object} entity.PurchaseResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /purchases/{id} [put]
func (p *purchaseRoutes) UpdatePurchase(c *gin.Context) {
	var req entity.PurchaseUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		p.log.Error("Error binding JSON", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	req.ID = id

	res, err := p.useCase.UpdatePurchase(&req)
	if err != nil {
		p.log.Error("Error updating purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetPurchase godoc
// @Summary Get Purchase
// @Description Retrieve a purchase by ID
// @Tags Purchase
// @Accept json
// @Produce json
// @Param id path string true "Purchase ID"
// @Success 200 {object} entity.PurchaseResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /purchases/{id} [get]
func (p *purchaseRoutes) GetPurchase(c *gin.Context) {
	var req entity.PurchaseID
	req.ID = c.Param("id")

	res, err := p.useCase.GetPurchase(&req)
	if err != nil {
		p.log.Error("Error fetching purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListPurchase godoc
// @Summary List Purchases
// @Description Retrieve a list of purchases
// @Tags Purchase
// @Accept json
// @Produce json
// @Param FilterPurchase query entity.FilterPurchase false "Purchase filter parameters"
// @Success 200 {array} entity.PurchaseList
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /purchases [get]
func (p *purchaseRoutes) GetListPurchase(c *gin.Context) {
	var req entity.FilterPurchase

	if err := c.ShouldBindQuery(&req); err != nil {
		p.log.Error("Error binding query", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := p.useCase.GetListPurchase(&req)
	if err != nil {
		p.log.Error("Error fetching purchase list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeletePurchase godoc
// @Summary Delete Purchase
// @Description Delete a purchase by ID
// @Tags Purchase
// @Accept json
// @Produce json
// @Param id path string true "Purchase ID"
// @Success 200 {object} entity.Message
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /purchases/{id} [delete]
func (p *purchaseRoutes) DeletePurchase(c *gin.Context) {
	var req entity.PurchaseID
	req.ID = c.Param("id")

	res, err := p.useCase.DeletePurchase(&req)
	if err != nil {
		p.log.Error("Error deleting purchase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
