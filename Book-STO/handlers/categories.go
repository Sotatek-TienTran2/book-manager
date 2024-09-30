package handlers

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	services service.CategoryServices
}

func NewCategoryHandler(services service.CategoryServices) CategoryHandler {

	return CategoryHandler{

		services: services,
	}
}

func (a CategoryHandler) GetListCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := a.services.ListCategories()
		if err != nil {

			WriteError(c, err)
			return
		}
		WriteRespon(c, http.StatusOK, res)
	}
}

func (a CategoryHandler) CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category dto.CreateCategoryRequest
		err := c.BindJSON(&category)
		if err != nil {

			WriteError(c, errs.ErrorReadRequestBody())
			return
		}
		_, e := a.services.CreateCategory(category)
		if e != nil {

			WriteError(c, e)
			return
		}
		WriteRespon(c, http.StatusOK, dto.MessageCreateSuccess("Category"))
	}
}

func (a CategoryHandler) SearchCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var author dto.SearchCategoryRequest
		err := c.BindJSON(&author)
		if err != nil {

			WriteError(c, errs.ErrorReadRequestBody())
			return
		}
		res, e := a.services.SearchCategory(author)
		if e != nil {

			WriteError(c, e)
			return
		}

		WriteRespon(c, http.StatusOK, res)
	}
}
