package repo

import "github.com/perfectogo/catalog-service/genproto/catalog"

type AuthorStorageInterface interface {
	InsertAuthor(author catalog.Author) (catalog.Author, error)
	SelectAuthors(page, limit int64) ([]*catalog.Author, int64, error)
	SelectAuthor(id string) (catalog.Author, error)
	UpdateAuthor(author catalog.Author) (catalog.Author, error)
	DeleteAuthor(id string) error
}
