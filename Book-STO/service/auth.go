package service

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/repository"
)

type AuthServices interface {
	LoginAuthor(req dto.LoginAuthorRequest) (*dto.LoginAuthorResponse, *errs.AppError)
}

type DefaultAuthServices struct {
	repo repository.AuthRepository
}

func NewAuthServices(repo repository.AuthRepository) AuthServices {

	return DefaultAuthServices{

		repo: repo,
	}
}

func (a DefaultAuthServices) LoginAuthor(req dto.LoginAuthorRequest) (*dto.LoginAuthorResponse, *errs.AppError) {
	author := dto.LoginAuthorRequest{
		Username: req.Username,
		Password: req.Password,
	}
	response, err := a.repo.Login(author)
	if err != nil {
		return &dto.LoginAuthorResponse{Status: response.Status, Username: response.Username, Token: response.Token, ExpireAt: response.ExpireAt}, err
	}
	return &dto.LoginAuthorResponse{Status: response.Status, Username: response.Username, Token: response.Token, ExpireAt: response.ExpireAt}, nil
}
