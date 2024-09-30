package handlers

import (
	"book-sto/errs"

	"github.com/gin-gonic/gin"
)

func WriteRespon(ctx *gin.Context, code int, data interface{}) {

	ctx.JSON(code, data)
}

func WriteError(ctx *gin.Context, err *errs.AppError) {

	ctx.JSON(err.Code, err)
}
