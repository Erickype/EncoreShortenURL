package url

import (
	"context"
)

// Shorten shortens a URL.
//
//encore:api public method=POST path=/url
func Shorten(ctx context.Context, p *ShortenParams) (*URL, error) {
	id, err := generateID()
	if err != nil {
		return nil, err
	} else if err := insert(ctx, id, p.URL); err != nil {
		return nil, err
	}
	return &URL{ID: id, URL: p.URL}, nil
}

// Get retrieves the original URL for the id.
//
//encore:api public method=GET path=/url/:id
func Get(ctx context.Context, id string) (*URL, error) {
	return getById(ctx, id)
}

// GetAll return all the URL entries in datasource as a []*URL
//
//encore:api public method=GET path=/url
func GetAll(ctx context.Context) (*URLs, error) {
	return getAll(ctx)
}
