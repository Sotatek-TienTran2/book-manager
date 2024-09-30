package routes

import (
	"book-sto/config"
	"book-sto/handlers"
	"book-sto/repository"
	"book-sto/service"

	"github.com/gin-gonic/gin"
)

func CategoryRoute(router *gin.Engine) {

	handler := handlers.NewCategoryHandler(service.NewCategoryServices(repository.NewCategoryRepository(config.DB)))
	route := router.Group("/api/categories")
	{

		route.GET("/", handler.GetListCategories())
		route.POST("/create", handler.CreateCategory())
		route.POST("/search", handler.SearchCategory())
	}
}
