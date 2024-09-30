package routes

import (
	"book-sto/config"
	"book-sto/handlers"
	"book-sto/middlewares"
	"book-sto/redis"
	"book-sto/repository"
	"book-sto/service"

	"github.com/gin-gonic/gin"
)

func AuthorRoute(router *gin.Engine) {

	handler := handlers.NewAuthorHandler(service.NewAuthorServices(repository.NewAuthorRepository(config.DB)))
	jwtMiddleware := middlewares.NewJWTMiddleware(repository.NewAuthorRepository(config.DB), redis.RDB)
	route := router.Group("/api/author")
	{

		route.GET("/", handler.GetListAuthor())
		route.POST("/create", handler.CreateAuthor())
		route.POST("/search", handler.SearchAuthor())
		route.GET("/show", jwtMiddleware.Verify(), handler.ShowBookByAuthor())
	}
}
