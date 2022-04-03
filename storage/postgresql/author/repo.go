package author

import "github.com/jmoiron/sqlx"

type authorRepo struct {
	db *sqlx.DB
}

func NewAuthorRepo(db *sqlx.DB) *authorRepo {
	return &authorRepo{db: db}
}
