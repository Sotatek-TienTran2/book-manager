package service

import (
	"book-sto/dto"
	"book-sto/errs"
	"book-sto/model"
	"book-sto/repository"
)

type BookService interface {
	IndexBook() (*dto.GetAllBookResponse, *errs.AppError)
	CreateBook(req dto.CreateBookRequest) (*dto.CreateBookResponse, *errs.AppError)
	SearchBookByAuthor(req dto.SearchBookByAuthorRequest) (*dto.SearchBookByAuthorResponse, *errs.AppError)
	SearchBookByCategory(req dto.SearchBookByCategoryRequest) (*dto.SearchBookByCategoryResponse, *errs.AppError)
}

type DefaultBookService struct {
	repo repository.BookRepository
}

func NewBookServices(repo repository.BookRepository) BookService {
	return DefaultBookService{

		repo: repo,
	}
}

func (a DefaultBookService) IndexBook() (*dto.GetAllBookResponse, *errs.AppError) {
	books, err := a.repo.IndexBook()
	if err != nil {

		return nil, err
	}
	return &dto.GetAllBookResponse{Books: books}, nil
}

func (a DefaultBookService) CreateBook(req dto.CreateBookRequest) (*dto.CreateBookResponse, *errs.AppError) {
	book := model.Book{
		Name:         req.Name,
		NameOfAuthor: req.NameOfAuthor,
		Category:     req.Category,
	}
	newBook, err := a.repo.CreateBook(book)
	if err != nil {
		return nil, err
	}
	return &dto.CreateBookResponse{Book: newBook}, nil
}

func (a DefaultBookService) SearchBookByAuthor(req dto.SearchBookByAuthorRequest) (*dto.SearchBookByAuthorResponse, *errs.AppError) {
	Name := req.NameOfAuthor
	res, err := a.repo.SearchBookByAuthor(Name)
	if err != nil {
		return nil, err
	}
	return &dto.SearchBookByAuthorResponse{Books: res}, nil
}
func (a DefaultBookService) SearchBookByCategory(req dto.SearchBookByCategoryRequest) (*dto.SearchBookByCategoryResponse, *errs.AppError) {
	Category := req.Category
	res, err := a.repo.SearchBookByCategory(Category)
	if err != nil {
		return nil, err
	}
	return &dto.SearchBookByCategoryResponse{Books: res}, nil
}
