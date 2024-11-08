package app

import (
	"crm-admin/config"
	"crm-admin/internal/controller/http"
	"crm-admin/internal/usecase"
	"crm-admin/internal/usecase/repo"
	"crm-admin/internal/usecase/token"
	"crm-admin/pkg/logger"
	"crm-admin/pkg/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

func Run(cfg config.Config) {
	logger1 := logger.NewLogger()

	db, err := postgres.Connection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = token.ConfigToken(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repo.NewUserRepo(db)
	userUseCase := usecase.NewUserUseCase(userRepo, logger1)

	engine := gin.Default()

	http.NewRouter(engine, logger1, userUseCase)

	log.Fatal(engine.Run(cfg.RUN_PORT))
}
