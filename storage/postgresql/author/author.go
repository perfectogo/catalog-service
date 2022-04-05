package author

import (
	"database/sql"
	"log"

	"github.com/perfectogo/catalog-service/genproto/catalog"
)

func (r *authorRepo) InsertAuthor(author catalog.Author) (catalog.Author, error) {
	var id string
	if err := r.db.QueryRow(`insert into authors (author_id, name) values ($1, $2) returning author_id`, author.AuthorId, author.AuthorName).Scan(&id); err != nil {
		return catalog.Author{}, nil
	}

	author, err := r.SelectAuthor(id)
	if err != nil {
		return catalog.Author{}, err
	}

	return author, nil
}
func (r *authorRepo) SelectAuthors(page, limit int64) ([]*catalog.Author, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query("select author_id, name, created_at, updated_at from authors where deleted_at is null limit $1 offset $2", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var (
		auhors []*catalog.Author
		count  int64
	)
	for rows.Next() {
		var auhor catalog.Author
		var updated_at sql.NullString
		if err := rows.Scan(
			&auhor.AuthorId,
			&auhor.AuthorName,
			&auhor.CreatedAt,
			&updated_at,
		); err != nil {
			return nil, 0, err
		}
		auhor.UpdatedAt = updated_at.String
		auhors = append(auhors, &auhor)
	}
	if err := r.db.QueryRow("select count(*) from authors where deleted_at is null").Scan(&count); err != nil {
		return nil, 0, err
	}

	return auhors, count, nil
}

func (r *authorRepo) SelectAuthor(id string) (catalog.Author, error) {
	var auhor catalog.Author
	var updated_at sql.NullString
	if err := r.db.QueryRow("select author_id, name, created_at, updated_at from authors where author_id=$1 and deleted_at is null",
		id).Scan(
		&auhor.AuthorId,
		&auhor.AuthorName,
		&auhor.CreatedAt,
		&updated_at,
	); err != nil {
		return catalog.Author{}, err
	}
	auhor.UpdatedAt = updated_at.String

	return auhor, nil
}

func (r *authorRepo) UpdateAuthor(author catalog.Author) (catalog.Author, error) {
	var id string
	log.Println(author.AuthorId)
	if err := r.db.QueryRow("update authors set name=$2, updated_at=current_timestamp where author_id=$1 and deleted_at is null returning author_id", author.AuthorId, author.AuthorName).Scan(&id); err != nil {
		return catalog.Author{}, err
	}
	author, err := r.SelectAuthor(id)
	if err != nil {
		return catalog.Author{}, err
	}

	return author, nil
}

func (r *authorRepo) DeleteAuthor(id string) error {
	result, err := r.db.Exec("update authors set deleted_at=current_timestamp where author_id=$1 and deleted_at is null", id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}
	return nil
}
