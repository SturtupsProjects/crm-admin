package http

import (
	_ "crm-admin/docs"
	"crm-admin/internal/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log/slog"
)

// title Api For CRM
// version 1.0
// description Admin Panel
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token here
func NewRouter(engine *gin.Engine, log *slog.Logger, ctr *controller.Controller) {

	engine.Use(CORSMiddleware())

	engine.GET("/swagger/*eny", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := engine.Group("/auth")
	product := engine.Group("/products")
	purchase := engine.Group("/purchase")
	sales := engine.Group("/sales")

	newUserRoutes(user, ctr.Auth, log)
	newProductRoutes(product, ctr.Product, log)
	newPurchaseRoutes(purchase, ctr.Purchase, log)
	newSalesRoutes(sales, ctr.Sales, log)
}
