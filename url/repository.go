package url

import (
	"context"
	"encore.dev/storage/sqldb"
)

// insert inserts a URL into the database.
func insert(ctx context.Context, id, url string) error {
	_, err := sqldb.Exec(ctx, `
		INSERT INTO url (id, original_url)
		VALUES ($1, $2)
	`, id, url)
	return err
}

// getById gets a URL by its ID
func getById(ctx context.Context, id string) (*URL, error) {
	url := &URL{
		ID:  id,
		URL: "",
	}
	query := "select original_url from public.url where id = $1"
	err := sqldb.QueryRow(ctx, query, id).Scan(&url.URL)
	if err != nil {
		return nil, err
	}

	return url, nil
}
