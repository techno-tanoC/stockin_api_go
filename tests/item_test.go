package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"stockin-api/domain"
	"stockin-api/handlers"
	"stockin-api/internal"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/kinbiko/jsonassert"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	schemaPath = "../schema.sql"
	base       = os.Getenv("TESTBASE")
)

func TestItemCRUD(t *testing.T) {
	ctx := context.Background()
	db, release := internal.WithTestDatabase(ctx, base, schemaPath)
	defer release()

	app := handlers.BuildApp(db)
	asserts := assert.New(t)
	ja := jsonassert.New(t)

	item := internal.NewItem()
	newItem := internal.NewItem()

	{
		req := httptest.NewRequest(
			"POST",
			"/items",
			strings.NewReader(fmt.Sprintf(
				`{
					"title": "%s",
					"url": "%s",
					"thumbnail": "%s"
				}`,
				item.Title,
				item.URL,
				item.Thumbnail,
			)),
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

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
			item.Title,
			item.URL,
			item.Thumbnail,
		)

		i := &struct {
			Data struct {
				ID uuid.UUID `json:"id"`
			} `json:"data"`
		}{}
		json.Unmarshal(rec.Body.Bytes(), i)
		item.ID = i.Data.ID
	}

	{
		req := httptest.NewRequest("GET", fmt.Sprintf("/items/%v", item.ID), nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

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

	{
		req := httptest.NewRequest(
			"PUT",
			fmt.Sprintf("/items/%s", item.ID),
			strings.NewReader(fmt.Sprintf(`{
					"title": "%s",
					"url": "%s",
					"thumbnail": "%s"
				}`,
				newItem.Title,
				newItem.URL,
				newItem.Thumbnail,
			)),
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

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
			newItem.Title,
			newItem.URL,
			newItem.Thumbnail,
		)
	}

	{
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/items/%v", item.ID), nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		asserts := assert.New(t)
		asserts.Equal(http.StatusNoContent, rec.Code)
	}
}

func TestItemFindByRange(t *testing.T) {
	ctx := context.Background()
	db, release := internal.WithTestDatabase(ctx, base, schemaPath)
	defer release()

	items, err := internal.CreateItems(ctx, db, 3)
	if err != nil {
		t.Fatal(err)
	}

	app := handlers.BuildApp(db)
	req := httptest.NewRequest("GET", fmt.Sprintf("/items?before=%s&limit=1", items[2].ID), nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	asserts := assert.New(t)
	ja := jsonassert.New(t)
	asserts.Equal(http.StatusOK, rec.Code)
	ja.Assertf(
		rec.Body.String(),
		`{
			"data": [{
				"id": "%s",
				"title": "%s",
				"url": "%s",
				"thumbnail": "%s",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			}],
			"message": ""
		}`,
		items[1].ID,
		items[1].Title,
		items[1].URL,
		items[1].Thumbnail,
	)
}

func TestItemImportAndExport(t *testing.T) {
	ctx := context.Background()
	db, release := internal.WithTestDatabase(ctx, base, schemaPath)
	defer release()

	app := handlers.BuildApp(db)
	asserts := assert.New(t)
	ja := jsonassert.New(t)

	items := []*domain.Item{}
	for i := 0; i < 3; i++ {
		item := internal.NewItem().Domain()
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ID.String() < items[j].ID.String() })
	bs, err := json.Marshal(items)
	if err != nil {
		t.Fatal(err)
	}
	j := string(bs)

	{
		req := httptest.NewRequest(
			"POST",
			"/items/import",
			strings.NewReader(j),
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		asserts.Equal(http.StatusNoContent, rec.Code)
	}

	{
		req := httptest.NewRequest("GET", "/items/export", nil)
		rec := httptest.NewRecorder()
		app.ServeHTTP(rec, req)

		asserts.Equal(http.StatusOK, rec.Code)
		ja.Assertf(rec.Body.String(), j)
	}
}
