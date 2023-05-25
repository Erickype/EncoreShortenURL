package url

import (
	"context"
	"testing"
)

// TestShortenAndRetrieve - test that the shortened URL is stored and retrieved from database.
func TestShortenAndRetrieve(t *testing.T) {
	testURL := "https://erickype.dev"
	sp := ShortenParams{URL: testURL}
	resp, err := Shorten(context.Background(), &sp)
	if err != nil {
		t.Fatal(err)
	}
	wantURL := testURL
	if resp.URL != wantURL {
		t.Errorf("got %q, want %q", resp.URL, wantURL)
	}

	firstURL := resp
	gotURL, err := Get(context.Background(), firstURL.ID)
	if err != nil {
		t.Fatal(err)
	}
	if *gotURL != *firstURL {
		t.Errorf("got %v, want %v", *gotURL, *firstURL)
	}
}
