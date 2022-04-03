package author

import "github.com/perfectogo/catalog-service/genproto/catalog"

func (r *authorRepo) InsertAuthor(author catalog.Author) (catalog.Author, error) {
	var id string
	if err := r.db.QueryRow(`insert into authors (author_id, name) values ($1, $2) returning author_id`, author.AuthorId, author.AuthorName).Scan(&id); err != nil {
		return catalog.Author{}, err
	}
	r.db.Close()

	author, err := r.SelectAuthor(id)
	if err != nil {
		return catalog.Author{}, err
	}

	return author, nil
}

func (r *authorRepo) SelectAuthor(id string) (catalog.Author, error) {
	var auhor catalog.Author
	if err := r.db.QueryRow("select author_id, name from authors where author_id=$1", id).Scan(&auhor.AuthorId, &auhor.AuthorName); err != nil {
		return catalog.Author{}, err
	}
	return auhor, nil
}

func (r *authorRepo) SelectAuthors(page, limit int64) ([]*catalog.Author, int64, error) {
	return nil, 0, nil
}

func (r *authorRepo) UpdateAuthor(id string) (catalog.Author, error) {
	return catalog.Author{}, nil
}

func (r *authorRepo) DeleteAuthor(id string) error {
	return nil
}
