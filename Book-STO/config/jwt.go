package config

import (
	"book-sto/errs"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var HmacSecret []byte

func NewJWTToken(data string) (*jwt.Token, *string, *errs.AppError) {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
	HmacSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(30 * 24 * time.Hour).Unix(),
		"data": data,
	})

	tokenString, err := token.SignedString(HmacSecret)
	if err != nil {
		fmt.Printf("Error when signed string token " + err.Error())
		return nil, nil, errs.NewUnexpectedError("Unexpected error when signed string token " + err.Error())
	}
	return token, &tokenString, nil
}

func VerifyJWTToken(tokenString string) (jwt.MapClaims, *errs.AppError) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf(fmt.Sprintf("Unexpected signing method: %v", t.Header["alg"]))
			return nil, errors.New("Unexpected signing method")
		}
		return HmacSecret, nil
	})

	if err != nil {
		v, _ := err.(*jwt.ValidationError)

		if v.Errors == jwt.ValidationErrorExpired {
			return nil, errs.NewUnauthenticatedError(v.Error())
		}

		return nil, errs.NewUnexpectedError("Unexpected error when parse token: " + err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		fmt.Printf("Error when verify token")
		return nil, errs.NewUnauthenticatedError("Invalid token")
	}
	return claims, nil
}
