package handlers

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	services service.AuthServices
}

func NewAuthHandler(services service.AuthServices) AuthHandler {

	return AuthHandler{
		services: services,
	}
}

func (a AuthHandler) LoginAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var author dto.LoginAuthorRequest
		err := ctx.BindJSON(&author)

		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		response, e := a.services.LoginAuthor(author)
		if err != nil {

			WriteError(ctx, e)
			return
		}
		if response.Status == "Success" {
			WriteRespon(ctx, http.StatusOK, dto.LoginSuccess("author", response.Token))
		} else {
			WriteRespon(ctx, http.StatusOK, dto.LoginFalse())
		}
	}
}

func (a AuthHandler) LogoutAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
