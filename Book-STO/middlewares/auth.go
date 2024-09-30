package middlewares

import (
	"book-sto/dto"
	"book-sto/handlers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type AuthMiddlewares struct {
	redis *redis.Client
}

func NewAuthMiddlewares(redis *redis.Client) *AuthMiddlewares {
	return &AuthMiddlewares{
		redis: redis,
	}
}

func (a AuthMiddlewares) CheckRegistration() gin.HandlerFunc {
	return func(c *gin.Context) {
		var author dto.LoginAuthorRequest
		err := c.BindJSON(&author)
		val, err1 := a.redis.Get(author.Username).Result()
		if err1 != nil {
			fmt.Println(err)
		}
		fmt.Printf(val)
		if val != "" {
			handlers.WriteRespon(c, http.StatusOK, dto.NotPermissions())
			return
		}
		c.Set("user", author.Username)
		c.Set("password", author.Password)
		c.Next()
	}
}
