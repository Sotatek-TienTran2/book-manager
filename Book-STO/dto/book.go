package dto

import "book-sto/model"

type GetAllBookResponse struct {
	Books []model.Book
}

type CreateBookRequest struct {
	Name         string `json:"Name,omitempty"`
	NameOfAuthor string `json:"NameOfAuthor,omitempty"`
	Category     string `json:"Category,omitempty"`
}

type CreateBookResponse struct {
	Book model.Book
}

type SearchBookByCategoryRequest struct {
	Category string `json:"Category"`
}

type SearchBookByCategoryResponse struct {
	Books []model.Book
}

type SearchBookByAuthorRequest struct {
	NameOfAuthor string `json:"NameOfAuthor,omitempty"`
}

type SearchBookByAuthorResponse struct {
	Books []model.Book
}
