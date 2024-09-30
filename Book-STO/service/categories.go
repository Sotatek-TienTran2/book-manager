package service

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/model"
	"book-sto/repository"
)

type CategoryServices interface {
	ListCategories() (*dto.GetAllCategoriesResponse, *errs.AppError)
	CreateCategory(req dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, *errs.AppError)
	SearchCategory(req dto.SearchCategoryRequest) (*dto.SearchCategoryResponse, *errs.AppError)
}

type DefaultCategoryServices struct {
	repo repository.CategoryRepository
}

func NewCategoryServices(repo repository.CategoryRepository) CategoryServices {

	return DefaultCategoryServices{

		repo: repo,
	}
}

func (a DefaultCategoryServices) ListCategories() (*dto.GetAllCategoriesResponse, *errs.AppError) {
	categories, err := a.repo.List()
	if err != nil {

		return nil, err
	}
	return &dto.GetAllCategoriesResponse{Categories: categories}, nil
}

func (a DefaultCategoryServices) CreateCategory(req dto.CreateCategoryRequest) (*dto.CreateCategoryResponse, *errs.AppError) {
	var category = model.Categories{
		Category: req.Category,
	}
	newCategory, err := a.repo.Create(category)
	if err != nil {
		return nil, err
	}
	return &dto.CreateCategoryResponse{Category: newCategory}, nil
}

func (a DefaultCategoryServices) SearchCategory(req dto.SearchCategoryRequest) (*dto.SearchCategoryResponse, *errs.AppError) {
	Category := req.Category
	res, e := a.repo.Search(Category)
	if e != nil {
		return nil, e
	}
	return &dto.SearchCategoryResponse{Categories: res}, nil
}
