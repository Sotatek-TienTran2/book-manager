package handlers

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service1Handler struct {
	client proto.AddAuthorServiceClient
}

func NewService1Handler(client proto.AddAuthorServiceClient) *Service1Handler {
	return &Service1Handler{
		client: client,
	}
}

func (h Service1Handler) LoginGRPC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var author dto.LoginAuthorRequest
		author.Username = ctx.MustGet("user").(string)
		author.Password = ctx.MustGet("password").(string)

		req := &proto.LoginRequest{Username: author.Username, Password: author.Password}
		if response, err := h.client.LoginGPRC(ctx, req); err == nil {
			if response.Status == "False" {
				WriteRespon(ctx, http.StatusOK, dto.LoginFalse())
			} else {
				WriteRespon(ctx, http.StatusOK, dto.LoginSuccess("author", response.Token))
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func (h Service1Handler) FindBookByIdAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter ID"})
			return
		}

		req := &proto.FindBookByIdAuthorRequest{IdAuthor: int64(a)}
		if response, err := h.client.FindBookByIdAuthor(ctx, req); err == nil {
			WriteRespon(ctx, http.StatusOK, response)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func (h Service1Handler) LogoutAuthor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(string)

		req := &proto.LogoutRequest{Token: user}
		if response, err := h.client.Logout(ctx, req); err == nil {
			if response.Status == "true" {
				WriteRespon(ctx, http.StatusOK, response)
			} else {
				WriteError(ctx, errs.ErrorInsertData())
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
