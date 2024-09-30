package routes

import (
	"book-sto/config"
	"book-sto/handlers"
	"book-sto/middlewares"
	"book-sto/proto"
	"book-sto/redis"
	"book-sto/repository"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Service(router *gin.Engine, conn *grpc.ClientConn) {
	client := proto.NewAddAuthorServiceClient(conn)
	route := router.Group("/api/service/author")
	jwtMiddleware := middlewares.NewJWTMiddleware(repository.NewAuthorRepository(config.DB), redis.RDB)
	//authMiddleware := middlewares.NewAuthMiddlewares(redis.RDB)
	handlers := handlers.NewService1Handler(client)
	{
		route.GET("/:id", handlers.FindBookByIdAuthor())
		//route.POST("/login", authMiddleware.CheckRegistration(), handlers.LoginGRPC())
		route.GET("/logout", jwtMiddleware.Verify(), handlers.LogoutAuthor())
	}
}
