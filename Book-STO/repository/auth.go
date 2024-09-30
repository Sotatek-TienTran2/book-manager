package repository

import (
	"book-sto/config"
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/model"
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
)

type AuthRepository interface {
	Login(dto.LoginAuthorRequest) (dto.LoginAuthorResponse, *errs.AppError)
}

type DefaultAuthRepository struct {
	redis *redis.Client
	db    *sql.DB
}

func NewAuthRepository(db *sql.DB, redis *redis.Client) AuthRepository {

	return DefaultAuthRepository{

		db:    db,
		redis: redis,
	}
}

func (a DefaultAuthRepository) Login(req dto.LoginAuthorRequest) (dto.LoginAuthorResponse, *errs.AppError) {
	author := model.Author{
		Username: req.Username,
		Password: req.Password,
	}
	var username string
	var response dto.LoginAuthorResponse
	res, err := a.db.Query("select author.username from tientran.author as author where author.username = ? and author.password = ?", author.Username, author.Password)
	for res.Next() {
		err = res.Scan(&username)
		if err != nil {
			response.Status = "False"
			return response, errs.ErrorReadData()
		}
	}
	response.Status = "False"
	if username != "" {
		a.redis.Del(username)
		jwtToken, token, e := config.NewJWTToken(username)
		if e != nil {
			return response, e
		}
		claims := jwtToken.Claims.(jwt.MapClaims)
		expiredAt := claims["exp"].(int64)
		res := dto.LoginAuthorResponse{
			Status:   "Success",
			Username: username,
			Token:    *token,
			ExpireAt: expiredAt,
		}
		return res, nil
	}
	return response, nil
}
