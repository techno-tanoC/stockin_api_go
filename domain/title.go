package domain

import (
	"fmt"
)

func TitleQuery(u *URL) (*Title, error) {
	doc, err := fetchDocument(u.URL)
	if err != nil {
		return nil, fmt.Errorf("fetch document error: %w", err)
	}

	title := doc.Find("html > head > title").First().Text()
	return &Title{title}, nil
}
