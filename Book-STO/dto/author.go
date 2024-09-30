package dto

import "book-sto/model"

type GetAllAuthorResponse struct {
	Authors []model.Author
}

type CreateAuthorResponse struct {
	Author model.Author
}

type FindBookByIdAuthor struct {
	Book []model.Book
}

type CreateAutherRequest struct {
	Name       string `json:"Name"`
	NativeLand string `json:"NativeLand"`
}

type SearchAuthorRequest struct {
	Name string `json:"Name"`
}

type SearchAuthorResponse struct {
	Authors []model.Author
}

type LoginAuthorRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type LoginAuthorResponse struct {
	Status   string `json:"status"`
	Username string `json:"username"`
	Token    string `json:"token" `
	ExpireAt int64  `json:"expire_at"`
}

type Author struct {
	IdAuthor   int    `json:"idAuthor,omitempty"`
	Name       string `json:"name,omitempty"`
	NativeLand string `json:"NativeLand,omitempty"`
	Username   string `json:"username,omitempty"`
}

type ShowBookByAuthorRequest struct {
	Username string `json:"username,omitempty"`
}

type ShowBookByAuthorResponse struct {
	Books []model.Book
}
