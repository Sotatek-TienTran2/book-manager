package routes

import (
	"book-sto/config"
	"book-sto/handlers"
	"book-sto/redis"
	"book-sto/repository"
	"book-sto/service"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	handler := handlers.NewAuthHandler(service.NewAuthServices(repository.NewAuthRepository(config.DB, redis.RDB)))
	route := router.Group("/api/auth")
	{
		route.POST("/login", handler.LoginAuthor())
	}
}
