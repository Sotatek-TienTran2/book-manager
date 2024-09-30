package handlers

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	services service.AuthorServices
}

func NewAuthorHandler(services service.AuthorServices) AuthorHandler {

	return AuthorHandler{
		services: services,
	}
}

func (a AuthorHandler) GetListAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		res, err := a.services.ListAuthor()
		if err != nil {

			WriteError(ctx, err)
			return
		}
		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (a AuthorHandler) CreateAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var author dto.CreateAutherRequest
		err := ctx.BindJSON(&author)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		_, e := a.services.CreateAuthor(author)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, dto.MessageCreateSuccess("Author"))
	}
}

func (a AuthorHandler) SearchAuthor() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var author dto.SearchAuthorRequest
		err := ctx.BindJSON(&author)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		res, e := a.services.SearchAuthor(author)
		if e != nil {

			WriteError(ctx, e)
			return
		}

		WriteRespon(ctx, http.StatusOK, res)
	}
}

func (a AuthorHandler) ShowBookByAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(string)
		res, e := a.services.ShowBookByAuthor(user)
		if e != nil {

			WriteError(ctx, e)
			return
		}

		WriteRespon(ctx, http.StatusOK, res)
	}
}
