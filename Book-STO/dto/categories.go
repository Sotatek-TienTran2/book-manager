package dto

import "book-sto/model"

type GetAllCategoriesResponse struct {
	Categories []model.Categories
}

type CreateCategoryResponse struct {
	Category model.Categories
}

type CreateCategoryRequest struct {
	Category string `json:"Category"`
}

type SearchCategoryRequest struct {
	Category string `json:"Category"`
}

type SearchCategoryResponse struct {
	Categories []model.Categories
}
