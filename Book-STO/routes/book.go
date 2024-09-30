package routes

import (
	"book-sto/config"
	"book-sto/handlers"
	"book-sto/repository"
	"book-sto/service"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {

	handler := handlers.NewBookHandler(service.NewBookServices(repository.NewBookRepository(config.DB)))
	route := router.Group("/api/book")
	{
		route.GET("/", handler.IndexBook())
		route.POST("/create", handler.CreateBook())
		route.POST("/search-by-author", handler.SearchBookByAuthor())
		route.POST("/search-by-category", handler.SearchBookByCategory())
	}
}
