package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"stockin-api/handlers"
	"stockin-api/internal"
	"testing"

	"github.com/kinbiko/jsonassert"
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
