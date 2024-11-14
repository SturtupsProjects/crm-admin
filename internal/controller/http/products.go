package http

import (
	"crm-admin/internal/entity"
	"crm-admin/internal/usecase"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type productRoutes struct {
	useCase *usecase.ProductsUseCase
	log     *slog.Logger
}

func newProductRoutes(router *gin.RouterGroup, us *usecase.ProductsUseCase, log *slog.Logger) {
	product := productRoutes{useCase: us, log: log}

	// ------------ product category router ------------------
	router.POST("/category", product.CreateCategory)
	router.GET("/category/:id", product.GetCategory)
	router.GET("/category", product.GetListCategory)
	router.DELETE("/category/:id", product.DeleteCategory)

	// -------------- product router --------------------------
	router.POST("", product.CreateProduct)
	router.GET("/:id", product.GetProduct)
	router.GET("", product.GetProductList)
	router.PUT("/:id", product.UpdateProduct)
	router.DELETE("/:id", product.DeleteProduct)
}

// CreateCategory godoc
// @Summary Create Product Category
// @Description Create a new product category
// @Tags Category
// @Accept json
// @Produce json
// @Param Category body entity.CategoryName true "Category data"
// @Success 201 {object} entity.Category
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/category [post]
func (p *productRoutes) CreateCategory(c *gin.Context) {
	var req entity.CategoryName

	if err := c.ShouldBindJSON(&req); err != nil {
		p.log.Error("Error in getting from body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := p.useCase.CreateCategory(&req)
	if err != nil {
		p.log.Error("Error in creating category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetCategory godoc
// @Summary Get Product Category
// @Description Retrieve a product category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} entity.Category
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/category/{id} [get]
func (p *productRoutes) GetCategory(c *gin.Context) {
	var req *entity.CategoryID

	id := c.Param("id")

	req.ID = id

	res, err := p.useCase.GetCategory(req)
	if err != nil {
		p.log.Error("Error in getting category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetListCategory godoc
// @Summary List Product Categories
// @Description Retrieve a list of product categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {array} entity.CategoryList
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/category [get]
func (p *productRoutes) GetListCategory(c *gin.Context) {
	var req *entity.CategoryName

	if err := c.ShouldBindQuery(req); err != nil {
		p.log.Error("Error in getting from body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := p.useCase.GetListCategory(req)
	if err != nil {
		p.log.Error("Error in getting category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteCategory godoc
// @Summary Delete Product Category
// @Description Delete a product category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} entity.Message
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/category/{id} [delete]
func (p *productRoutes) DeleteCategory(c *gin.Context) {
	var req *entity.CategoryID

	id := c.Param("id")

	req.ID = id

	res, err := p.useCase.DeleteCategory(req)
	if err != nil {
		p.log.Error("Error in deleting category", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateProduct godoc
// @Summary Create Product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body entity.ProductRequest true "Product data"
// @Success 201 {object} entity.Product
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products [post]
func (p *productRoutes) CreateProduct(c *gin.Context) {
	var req *entity.ProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		p.log.Error("Error in getting from body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := p.useCase.CreateProduct(req)
	if err != nil {
		p.log.Error("Error in creating product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetProduct godoc
// @Summary Get Product
// @Description Retrieve a product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/{id} [get]
func (p *productRoutes) GetProduct(c *gin.Context) {
	var req *entity.ProductID

	id := c.Param("id")

	req.ID = id

	res, err := p.useCase.GetProduct(req)
	if err != nil {
		p.log.Error("Error in getting product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetProductList godoc
// @Summary List Products
// @Description Retrieve a list of products with optional filters
// @Tags Product
// @Accept json
// @Produce json
// @Param FilterProduct query entity.FilterProduct false "Product filter parameters"
// @Success 200 {array} entity.ProductList
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products [get]
func (p *productRoutes) GetProductList(c *gin.Context) {
	var req *entity.FilterProduct

	if err := c.ShouldBindQuery(req); err != nil {
		p.log.Error("Error in getting from body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := p.useCase.GetProductList(req)
	if err != nil {
		p.log.Error("Error in getting product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateProduct godoc
// @Summary Update Product
// @Description Update product details
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param UpdateProduct body entity.ProductUpdate true "Updated product data"
// @Success 200 {object} entity.Product
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/{id} [put]
func (p *productRoutes) UpdateProduct(c *gin.Context) {
	var req *entity.ProductUpdate

	if err := c.ShouldBindJSON(req); err != nil {
		p.log.Error("Error in getting from body", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	req.ID = id

	res, err := p.useCase.UpdateProduct(req)
	if err != nil {
		p.log.Error("Error in updating product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteProduct godoc
// @Summary Delete Product
// @Description Delete a product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Message
// @Failure 400 {object} entity.Error
// @Failure 500 {object} entity.Error
// @Router /products/{id} [delete]
func (p *productRoutes) DeleteProduct(c *gin.Context) {
	var req *entity.ProductID

	id := c.Param("id")
	req.ID = id

	res, err := p.useCase.DeleteProduct(req)
	if err != nil {
		p.log.Error("Error in deleting product", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
