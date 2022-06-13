package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"stockin/domain"
	"stockin/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port  int    `env:"PORT,default=3000"`
	Token string `env:"TOKEN,required"`
}

func main() {
	ctx := context.Background()
	conf := new(Config)
	err := envconfig.Process(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	rawDB, err := sql.Open("mysql", "root:pass@(db)/dev?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer rawDB.Close()
	db := &domain.RealDB{DB: rawDB}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.KeyAuth(auth(conf.Token)))

	item := e.Group("/items")
	item.GET("", handler.ItemIndex(db))
	item.POST("", handler.ItemCreate(db))
	item.PUT("/:id", handler.ItemUpdate(db))
	item.DELETE("/:id", handler.ItemDelete(db))

	title := e.Group("/title")
	title.POST("/query", handler.TitleQuery)

	thumbnail := e.Group("/thumbnail")
	thumbnail.POST("/query", handler.ThumbnailQuery)

	err = e.Start(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func auth(token string) middleware.KeyAuthValidator {
	return func(key string, c echo.Context) (bool, error) {
		return key == token, nil
	}
}
