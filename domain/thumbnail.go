package domain

import "fmt"

func ThumbnailQuery(u *URL) (*Thumbnail, error) {
	doc, err := fetchDocument(u.URL)
	if err != nil {
		return nil, fmt.Errorf("fetch document error: %w", err)
	}

	// TODO find favicon if not found
	url := doc.Find(`html > head > meta[property="og:image"]`).First().AttrOr("content", "")
	return &Thumbnail{url}, nil
}
