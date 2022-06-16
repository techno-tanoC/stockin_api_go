package main

import (
	"context"
	"database/sql"
	"log"
	"stockin/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:pass@(db)/dev?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items() {
		err = item.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func items() []*models.Item {
	return []*models.Item{
		{
			Title:     "Android",
			URL:       "https://www.android.com/",
			Thumbnail: "",
		},
		{
			Title:     "TypeScript",
			URL:       "https://www.typescriptlang.org/",
			Thumbnail: "",
		},
		{
			Title:     "The Go Programming Language",
			URL:       "https://golang.org/",
			Thumbnail: "",
		},
		{
			Title:     "Docker",
			URL:       "https://www.docker.com/",
			Thumbnail: "",
		},
		{
			Title:     "Haskell Language",
			URL:       "https://www.haskell.org/",
			Thumbnail: "",
		},
		{
			Title:     "Rustプログラミング言語",
			URL:       "https://www.rust-lang.org/ja",
			Thumbnail: "",
		},
		{
			Title:     "Qiita",
			URL:       "https://qiita.com/",
			Thumbnail: "",
		},
		{
			Title:     "Zenn",
			URL:       "https://zenn.dev/",
			Thumbnail: "",
		},
		{
			Title:     "GitLab",
			URL:       "https://gitlab.com/",
			Thumbnail: "",
		},
		{
			Title:     "GitHub",
			URL:       "https://github.com/",
			Thumbnail: "",
		},
		{
			Title:     "Twitter",
			URL:       "https://twitter.com/home",
			Thumbnail: "",
		},
		{
			Title:     "GMail",
			URL:       "https://mail.google.com/",
			Thumbnail: "",
		},
		{
			Title:     "Google Drive",
			URL:       "https://drive.google.com/",
			Thumbnail: "",
		},
		{
			Title:     "ニコニコ動画",
			URL:       "https://www.nicovideo.jp/",
			Thumbnail: "",
		},
		{
			Title:     "YouTube",
			URL:       "https://www.youtube.com/",
			Thumbnail: "",
		},
	}
}
