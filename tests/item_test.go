package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"stockin-api/handlers"
	"stockin-api/internal"
	"strings"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	schemaPath = "../schema.sql"
	base       = os.Getenv("TESTBASE")
)

func TestItemFind(t *testing.T) {
	ctx := context.Background()
	db, release := internal.WithTestDatabase(ctx, base, schemaPath)
	defer release()

	item, err := internal.CreateItem(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	app := handlers.BuildApp(db)
	req := httptest.NewRequest("GET", fmt.Sprintf("/items/%v", item.ID), nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	asserts := assert.New(t)
	ja := jsonassert.New(t)
	asserts.Equal(http.StatusOK, rec.Code)
	ja.Assertf(
		rec.Body.String(),
		`{
			"data": {
				"id": "%s",
				"title": "%s",
				"url": "%s",
				"thumbnail": "%s",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			"message": ""
		}`,
		item.ID,
		item.Title,
		item.URL,
		item.Thumbnail,
	)
}

func TestItemCreate(t *testing.T) {
	ctx := context.Background()
	db, release := internal.WithTestDatabase(ctx, base, schemaPath)
	defer release()

	title := "example"
	url := "https://example.com/"
	thumbnail := "https://example.com/image.jpg"

	app := handlers.BuildApp(db)
	req := httptest.NewRequest(
		"POST",
		"/items/",
		strings.NewReader(fmt.Sprintf(
			`{
				"title": "%s",
				"url": "%s",
				"thumbnail": "%s"
			}`,
			title,
			url,
			thumbnail,
		)),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	asserts := assert.New(t)
	ja := jsonassert.New(t)
	asserts.Equal(http.StatusOK, rec.Code)
	ja.Assertf(
		rec.Body.String(),
		`{
			"data": {
				"id": "<<PRESENCE>>",
				"title": "%s",
				"url": "%s",
				"thumbnail": "%s",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			"message": ""
		}`,
		title,
		url,
		thumbnail,
	)
}

func TestUpdateTest(t *testing.T) {
	ctx := context.Background()
	db, release := internal.WithTestDatabase(ctx, base, schemaPath)
	defer release()

	item, err := internal.CreateItem(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	title := "example"
	url := "https://example.com/"
	thumbnail := "https://example.com/image.jpg"

	app := handlers.BuildApp(db)
	req := httptest.NewRequest(
		"PUT",
		fmt.Sprintf("/items/%s", item.ID),
		strings.NewReader(fmt.Sprintf(
			`{
				"title": "%s",
				"url": "%s",
				"thumbnail": "%s"
			}`,
			title,
			url,
			thumbnail,
		)),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	asserts := assert.New(t)
	ja := jsonassert.New(t)
	asserts.Equal(http.StatusOK, rec.Code)
	ja.Assertf(
		rec.Body.String(),
		`{
			"data": {
				"id": "%s",
				"title": "%s",
				"url": "%s",
				"thumbnail": "%s",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			},
			"message": ""
		}`,
		item.ID,
		title,
		url,
		thumbnail,
	)

}
