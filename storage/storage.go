package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/perfectogo/catalog-service/storage/postgresql/author"
	"github.com/perfectogo/catalog-service/storage/repo"
)

type InterfaceStorage interface {
	Author() repo.AuthorStorageInterface
}

type storagePg struct {
	db         *sqlx.DB
	authorRepo repo.AuthorStorageInterface
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:         db,
		authorRepo: author.NewAuthorRepo(db),
	}

}
func (s *storagePg) Author() repo.AuthorStorageInterface {
	return s.authorRepo
}
