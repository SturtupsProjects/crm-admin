package http

import (
	"crm-admin/internal/entity"
	"crm-admin/internal/usecase"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type salesRoutes struct {
	useCase *usecase.SalesUseCase
	log     *slog.Logger
}

func newSalesRoutes(router *gin.RouterGroup, us *usecase.SalesUseCase, log *slog.Logger) {
	sales := &salesRoutes{useCase: us, log: log}

	// Sales routes
	router.POST("", sales.CreateSale)
	router.GET("/:id", sales.GetSale)
	router.GET("", sales.GetListSales)
	router.PUT("/:id", sales.UpdateSale)
	router.DELETE("/:id", sales.DeleteSale)
}

// CreateSale godoc
// @Summary Create Sale
// @Description Record a new sale transaction
// @Tags Sales
// @Accept json
// @Produce json
// @Param SaleRequest body entity.SaleRequest true "Sale data"
// @Success 201 {object} entity.SaleResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /sales [post]
func (s *salesRoutes) CreateSale(c *gin.Context) {
	var req entity.SaleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error("Error binding JSON in CreateSale", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.useCase.CreateSales(&req)
	if err != nil {
		s.log.Error("Error creating sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetSale godoc
// @Summary Get Sale
// @Description Retrieve a sale by ID
// @Tags Sales
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {object} entity.SaleResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /sales/{id} [get]
func (s *salesRoutes) GetSale(c *gin.Context) {
	var req entity.SaleID
	req.ID = c.Param("id")

	res, err := s.useCase.GetSales(&req)
	if err != nil {
		s.log.Error("Error retrieving sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListSales godoc
// @Summary List Sales
// @Description Retrieve a list of sales with optional filters
// @Tags Sales
// @Accept json
// @Produce json
// @Param SaleFilter query entity.SaleFilter false "Sales filter parameters"
// @Success 200 {array} entity.SaleList
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /sales [get]
func (s *salesRoutes) GetListSales(c *gin.Context) {
	var req entity.SaleFilter

	if err := c.ShouldBindQuery(&req); err != nil {
		s.log.Error("Error binding query parameters in GetListSales", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.useCase.GetListSales(&req)
	if err != nil {
		s.log.Error("Error retrieving sales list", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateSale godoc
// @Summary Update Sale
// @Description Update details of an existing sale
// @Tags Sales
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Param SaleUpdate body entity.SaleUpdate true "Updated sale data"
// @Success 200 {object} entity.SaleResponse
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /sales/{id} [put]
func (s *salesRoutes) UpdateSale(c *gin.Context) {
	var req entity.SaleUpdate
	req.ID = c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		s.log.Error("Error binding JSON in UpdateSale", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.useCase.UpdateSales(&req)
	if err != nil {
		s.log.Error("Error updating sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteSale godoc
// @Summary Delete Sale
// @Description Delete a sale by ID
// @Tags Sales
// @Accept json
// @Produce json
// @Param id path string true "Sale ID"
// @Success 200 {object} entity.Message
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /sales/{id} [delete]
func (s *salesRoutes) DeleteSale(c *gin.Context) {
	var req entity.SaleID
	req.ID = c.Param("id")

	res, err := s.useCase.DeleteSales(&req)
	if err != nil {
		s.log.Error("Error deleting sale", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
