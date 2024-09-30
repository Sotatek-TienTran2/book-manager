package handlers

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHander struct {
	services service.BookService
}

func NewBookHandler(services service.BookService) BookHander {
	return BookHander{

		services: services,
	}
}

func (a BookHander) IndexBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := a.services.IndexBook()
		if err != nil {

			WriteError(c, err)
			return
		}
		WriteRespon(c, http.StatusOK, res)
	}
}

func (a BookHander) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book dto.CreateBookRequest
		err := c.BindJSON(&book)
		if err != nil {

			WriteError(c, errs.ErrorReadRequestBody())
			return
		}
		_, e := a.services.CreateBook(book)
		if e != nil {

			WriteError(c, e)
			return
		}
		WriteRespon(c, http.StatusOK, dto.MessageCreateSuccess("Book"))
	}
}

func (a BookHander) SearchBookByAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var author dto.SearchBookByAuthorRequest
		err := c.BindJSON(&author)
		if err != nil {

			WriteError(c, errs.ErrorReadRequestBody())
			return
		}
		res, e := a.services.SearchBookByAuthor(author)
		if e != nil {

			WriteError(c, e)
			return
		}

		WriteRespon(c, http.StatusOK, res)
	}
}

func (a BookHander) SearchBookByCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category dto.SearchBookByCategoryRequest
		err := c.BindJSON(&category)
		if err != nil {

			WriteError(c, errs.ErrorReadRequestBody())
			return
		}
		res, e := a.services.SearchBookByCategory(category)
		if e != nil {

			WriteError(c, e)
			return
		}

		WriteRespon(c, http.StatusOK, res)
	}
}
