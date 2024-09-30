package model

type Book struct {
	IdBook       int    `json:"_idBook,omitempty"`
	Name         string `json:"Name,omitempty"`
	NameOfAuthor string `json:"NameOfAuthor,omitempty"`
	Category     string `json:"Category,omitempty"`
}
