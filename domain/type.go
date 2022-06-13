package domain

type URL struct {
	URL string `json:"url"`
}

type Title struct {
	Title string `json:"title"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type ItemParams struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}
