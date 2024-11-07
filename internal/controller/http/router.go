package http

import (
	_ "crm-admin/docs"
	"crm-admin/internal/usecase"

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
func NewRouter(engine *gin.Engine, log *slog.Logger, u *usecase.UserUseCase) {

	engine.Use(CORSMiddleware())

	engine.GET("/swagger/*eny", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := engine.Group("/auth")

	newUserRoutes(user, u, log)
}
