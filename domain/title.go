package domain

import (
	"fmt"
)

func TitleQuery(u string) (string, error) {
	doc, err := fetchDocument(u)
	if err != nil {
		return "", fmt.Errorf("fetch document error: %w", err)
	}

	title := doc.Find("html > head > title").First().Text()
	return title, nil
}
