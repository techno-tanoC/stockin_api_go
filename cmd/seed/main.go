package main

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"os"
	"stockin-api/domain"
	"stockin-api/queries"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	database := os.Getenv("DATABASE")

	db, err := sql.Open("postgres", database)
	if err != nil {
		log.Fatal(err)
	}
	q := queries.New(db)

	for _, params := range items() {
		ps := params.BuildForInsert()
		err = q.InsertItem(ctx, *ps)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, params := range tags() {
		ps := params.BuildForInsert()
		err = q.InsertTag(ctx, *ps)
		if err != nil {
			log.Fatal(err)
		}
	}

	items, err := q.IndexItems(ctx)
	if err != nil {
		log.Fatal(err)
	}

	tags, err := q.IndexTags(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		size := rand.Intn(6)
		for i, tag := range tags {
			if i < size {
				params := domain.ItemTagParams{
					ItemID: item.ID,
					TagID:  tag.ID,
				}
				ps := params.BuildForInsert()
				err = q.InsertItemTag(ctx, *ps)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func tags() []*domain.TagParams {
	return []*domain.TagParams{
		{
			Name: "A",
		},
		{
			Name: "B",
		},
		{
			Name: "C",
		},
		{
			Name: "D",
		},
		{
			Name: "E",
		},
	}
}

func items() []*domain.ItemParams {
	return []*domain.ItemParams{
		{
			Title:     "Android",
			URL:       "https://www.android.com/",
			Thumbnail: "https://lh3.googleusercontent.com/GTmuiIZrppouc6hhdWiocybtRx1Tpbl52eYw4l-nAqHtHd4BpSMEqe-vGv7ZFiaHhG_l4v2m5Fdhapxw9aFLf28ErztHEv5WYIz5fA",
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
			Thumbnail: "https://www.docker.com/wp-content/uploads/2022/05/Docker_Temporary_Image_Social_Thumbnail_1200x630_v5.png",
		},
		{
			Title:     "Haskell Language",
			URL:       "https://www.haskell.org/",
			Thumbnail: "",
		},
		{
			Title:     "Rustプログラミング言語",
			URL:       "https://www.rust-lang.org/ja",
			Thumbnail: "https://www.rust-lang.org/static/images/rust-social-wide.jpg",
		},
		{
			Title:     "GitLab",
			URL:       "https://gitlab.com/",
			Thumbnail: "",
		},
		{
			Title:     "GitHub",
			URL:       "https://github.com/",
			Thumbnail: "https://github.githubassets.com/images/modules/site/social-cards/github-social.png",
		},
		{
			Title:     "Twitter",
			URL:       "https://twitter.com/home",
			Thumbnail: "",
		},
		{
			Title:     "ニコニコ動画",
			URL:       "https://www.nicovideo.jp/",
			Thumbnail: "https://nicovideo.cdn.nimg.jp/uni/images/ogp.png",
		},
		{
			Title:     "YouTube",
			URL:       "https://www.youtube.com/",
			Thumbnail: "https://www.youtube.com/img/desktop/yt_1200.png",
		},
		{
			Title:     "Qiita",
			URL:       "https://qiita.com/",
			Thumbnail: "https://cdn.qiita.com/assets/qiita-ogp-3b6fcfdd74755a85107071ffc3155898.png",
		},
		{
			Title:     "CircleCIでイメージをビルドしてGCRにプッシュする",
			URL:       "https://qiita.com/techno-tanoC/items/845bb906156e66a24b7f",
			Thumbnail: "https://qiita-user-contents.imgix.net/https%3A%2F%2Fcdn.qiita.com%2Fassets%2Fpublic%2Farticle-ogp-background-9f5428127621718a910c8b63951390ad.png?ixlib=rb-4.0.0&w=1200&mark64=aHR0cHM6Ly9xaWl0YS11c2VyLWNvbnRlbnRzLmltZ2l4Lm5ldC9-dGV4dD9peGxpYj1yYi00LjAuMCZ3PTkxNiZ0eHQ9Q2lyY2xlQ0klRTMlODElQTclRTMlODIlQTQlRTMlODMlQTElRTMlODMlQkMlRTMlODIlQjglRTMlODIlOTIlRTMlODMlOTMlRTMlODMlQUIlRTMlODMlODklRTMlODElOTclRTMlODElQTZHQ1IlRTMlODElQUIlRTMlODMlOTclRTMlODMlODMlRTMlODIlQjclRTMlODMlQTUlRTMlODElOTklRTMlODIlOEImdHh0LWNvbG9yPSUyMzIxMjEyMSZ0eHQtZm9udD1IaXJhZ2lubyUyMFNhbnMlMjBXNiZ0eHQtc2l6ZT01NiZ0eHQtY2xpcD1lbGxpcHNpcyZ0eHQtYWxpZ249bGVmdCUyQ3RvcCZzPWNkMWQ0NDlkNzRhZWUzMDUyOTc0NDBlYzY2MzlmNGI3&mark-x=142&mark-y=112&blend64=aHR0cHM6Ly9xaWl0YS11c2VyLWNvbnRlbnRzLmltZ2l4Lm5ldC9-dGV4dD9peGxpYj1yYi00LjAuMCZ3PTYxNiZ0eHQ9JTQwdGVjaG5vLXRhbm9DJnR4dC1jb2xvcj0lMjMyMTIxMjEmdHh0LWZvbnQ9SGlyYWdpbm8lMjBTYW5zJTIwVzYmdHh0LXNpemU9MzYmdHh0LWFsaWduPWxlZnQlMkN0b3Amcz0yMjcyZWM5Y2RjNTBhMjFiOTk5OWI2MWVmODEwMmRmNg&blend-x=142&blend-y=491&blend-mode=normal&s=5ee730838a45a3671bbae00f205833d6",
		},
		{
			Title:     "Zenn",
			URL:       "https://zenn.dev/",
			Thumbnail: "https://zenn.dev/images/logo-only-dark.png",
		},
		{
			Title:     "Rustの新しいWEBフレームワークaxumを触ってみた",
			URL:       "https://zenn.dev/techno_tanoc/articles/99e54c82cb049f",
			Thumbnail: "https://res.cloudinary.com/zenn/image/upload/s--iptHc_8i--/co_rgb:222%2Cg_south_west%2Cl_text:notosansjp-medium.otf_37_bold:techno-tanoC%2Cx_203%2Cy_98/c_fit%2Cco_rgb:222%2Cg_north_west%2Cl_text:notosansjp-medium.otf_70_bold:Rust%25E3%2581%25AE%25E6%2596%25B0%25E3%2581%2597%25E3%2581%2584WEB%25E3%2583%2595%25E3%2583%25AC%25E3%2583%25BC%25E3%2583%25A0%25E3%2583%25AF%25E3%2583%25BC%25E3%2582%25AFaxum%25E3%2582%2592%25E8%25A7%25A6%25E3%2581%25A3%25E3%2581%25A6%25E3%2581%25BF%25E3%2581%259F%2Cw_1010%2Cx_90%2Cy_100/g_south_west%2Ch_90%2Cl_fetch:aHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUFUWEFKd2p0VEdpUklNSmFrSFFJV0JIUmFyUEwzUmFSak1oSUhXeWtCRUc9czk2LWM=%2Cr_max%2Cw_90%2Cx_87%2Cy_72/v1627274783/default/og-base_z4sxah.png",
		},
	}
}
