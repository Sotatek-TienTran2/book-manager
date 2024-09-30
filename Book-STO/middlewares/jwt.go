package middlewares

import (
	"book-sto/config"
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/handlers"
	"book-sto/repository"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type JWTMiddleware struct {
	repo  repository.AuthorRepository
	redis *redis.Client
}

func (m JWTMiddleware) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		authentiationHeader := c.Request.Header.Get("Authorization")
		if authentiationHeader == "" {
			handlers.WriteError(c, errs.NewUnauthenticatedError("Unauthorized"))
		}
		arr := strings.Split(authentiationHeader, " ")
		if len(arr) <= 1 {
			handlers.WriteError(c, errs.NewUnauthenticatedError("Invalid token"))
			return
		}
		token := arr[1]
		claims, err := config.VerifyJWTToken(token)
		if err != nil {
			handlers.WriteError(c, err)
			return
		}

		username := claims["data"].(string)
		val, err1 := m.redis.Get(username).Result()
		if err1 != nil {
			fmt.Println(err)
		}
		if val != "" {
			handlers.WriteRespon(c, http.StatusOK, dto.NotPermissions())
			return
		}

		user, err := m.repo.FindAuthorByUsername(username)
		if err != nil {
			handlers.WriteError(c, err)
			return
		}
		if user == "false" {
			handlers.WriteRespon(c, http.StatusOK, dto.NotPermissions())
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func NewJWTMiddleware(repo repository.AuthorRepository, redis *redis.Client) JWTMiddleware {
	return JWTMiddleware{
		repo:  repo,
		redis: redis,
	}
}
