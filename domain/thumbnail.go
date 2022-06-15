package domain

import "fmt"

func ThumbnailQuery(u string) (string, error) {
	doc, err := fetchDocument(u)
	if err != nil {
		return "", fmt.Errorf("fetch document error: %w", err)
	}

	// TODO find favicon if not found
	url := doc.Find(`html > head > meta[property="og:image"]`).First().AttrOr("content", "")
	return url, nil
}
