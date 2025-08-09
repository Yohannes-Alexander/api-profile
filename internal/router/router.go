package router

import (
	"database/sql"
	"github.com/Yohannes-Alexander/api-profile/internal/handler"
	"github.com/Yohannes-Alexander/api-profile/internal/repository"	
	"github.com/Yohannes-Alexander/api-profile/internal/service"	

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	api := r.Group("/api")
	{
		api.POST("/login", authHandler.Login)
		api.POST("/refresh", authHandler.RefreshToken)
	}

	return r
}
