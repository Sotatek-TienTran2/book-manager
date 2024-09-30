package repository

import (
	"book-sto/errs"
	"book-sto/model"
	"database/sql"
	"strings"
)

type BookRepository interface {
	IndexBook() ([]model.Book, *errs.AppError)
	CreateBook(book model.Book) (model.Book, *errs.AppError)
	SearchBookByAuthor(req string) ([]model.Book, *errs.AppError)
	SearchBookByCategory(req string) ([]model.Book, *errs.AppError)
}

type DefaultBookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {

	return DefaultBookRepository{

		db: db,
	}
}

func (a DefaultBookRepository) IndexBook() ([]model.Book, *errs.AppError) {
	var books []model.Book
	var book model.Book
	res, err := a.db.Query("SELECT book.idBook,book.Name FROM tientran.book as book")
	if err != nil {

		return nil, errs.ErrorGetData()
	}
	for res.Next() {
		err = res.Scan(&book.IdBook, &book.Name)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		books = append(books, book)
	}
	return books, nil
}

func (a DefaultBookRepository) CreateBook(book model.Book) (model.Book, *errs.AppError) {
	var idBook int
	NewBook, err := a.db.Prepare("INSERT INTO tientran.book(Name) VALUES (?) ")
	if err != nil {
		return book, errs.ErrorInsertData()
	}
	NewBook.Exec(book.Name)

	err1 := a.db.QueryRow("SELECT book.idBook FROM tientran.book as book WHERE book.Name = ?", book.Name).Scan(&idBook)
	if err1 != nil || err1 == sql.ErrNoRows {
		return book, errs.ErrorGetData()
	}

	author := strings.Split(book.NameOfAuthor, "; ")
	for i := 0; i < len(author); i++ {
		var idAuthor int
		err := a.db.QueryRow("SELECT author.idAuthor FROM tientran.author as author WHERE author.Name = ?", author[i]).Scan(&idAuthor)
		if err != nil || err == sql.ErrNoRows {
			return book, errs.ErrorGetData()
		}

		NewBookAuthor, err := a.db.Prepare("INSERT INTO tientran.book_author (idAuthor, idBook) VALUES (?, ?) ")
		if err != nil {
			return book, errs.ErrorInsertData()
		}
		NewBookAuthor.Exec(idAuthor, idBook)
	}

	categories := strings.Split(book.Category, "; ")
	for i := 0; i < len(categories); i++ {
		var idCategories int
		err1 := a.db.QueryRow("SELECT categories.idCategories FROM tientran.categories as categories WHERE categories.Category = ?", categories[i]).Scan(&idCategories)
		if err1 != nil || err1 == sql.ErrNoRows {
			return book, errs.ErrorGetData()
		}

		selDB1, err := a.db.Prepare("INSERT INTO tientran.book_categories (idCategories, idBook) VALUES (?, ?) ")
		if err != nil {
			return book, errs.ErrorInsertData()
		}
		selDB1.Exec(idCategories, idBook)
	}
	defer a.db.Close()

	return book, nil
}

func (a DefaultBookRepository) SearchBookByAuthor(req string) ([]model.Book, *errs.AppError) {
	var books []model.Book
	if req == "" {
		return books, nil
	}
	bodyString := "%" + req + "%"
	res, err := a.db.Query("SELECT book.Name, author.Name FROM tientran.book as book,  tientran.author as author, tientran.book_author as ba WHERE book.idBook = ba.idBook AND ba.idAuthor = author.idAuthor AND author.Name LIKE ?", bodyString)
	if err != nil {
		return books, errs.ErrorGetData()
	}
	for res.Next() {
		var book model.Book
		err = res.Scan(&book.Name, &book.NameOfAuthor)
		if err != nil {
			return nil, errs.ErrorReadData()
		}
		books = append(books, book)
	}
	return books, nil
}

func (a DefaultBookRepository) SearchBookByCategory(req string) ([]model.Book, *errs.AppError) {
	var books []model.Book
	if req == "" {
		return books, nil
	}
	bodyString := "%" + req + "%"
	res, err := a.db.Query("SELECT book.Name, categories.Category FROM tientran.book as book,  tientran.categories as categories, tientran.book_categories as bc WHERE book.idBook = bc.idBook AND bc.idCategories = categories.idCategories AND categories.Category LIKE ?", bodyString)

	if err != nil {
		return books, errs.ErrorGetData()
	}
	for res.Next() {
		var book model.Book
		err = res.Scan(&book.Name, &book.Category)
		if err != nil {
			return nil, errs.ErrorReadData()
		}
		books = append(books, book)
	}
	return books, nil
}
