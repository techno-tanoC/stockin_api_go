package domain_test

import (
	"context"
	"database/sql"
	"fmt"
	"stockin/domain"
	"stockin/models"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var itemOpt = cmpopts.IgnoreFields(models.Item{}, "ID", "CreatedAt", "UpdatedAt")

// inserted:      |----------|
// items   : |----------|
func TestItemIndex1(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	defer release()

	inserted, err := insertItemMany(ctx, db)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}

	items, err := domain.ItemIndex(ctx, db, "ffffffff-ffff-ffff-ffff-ffffffffffff", 10)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	diff := cmp.Diff(len(items), 10)
	if diff != "" {
		t.Fatalf("TestItemIndex: %v", diff)
	}
	for i := 0; i < 10; i++ {
		diff = cmp.Diff(items[i], inserted[i])
		if diff != "" {
			t.Fatalf("TestItemIndex: %v", diff)
		}
	}
}

// inserted:      |----------|
// items   : |--------------------|
func TestItemIndex2(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	defer release()

	inserted, err := insertItemMany(ctx, db)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}

	items, err := domain.ItemIndex(ctx, db, "ffffffff-ffff-ffff-ffff-ffffffffffff", 100)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	diff := cmp.Diff(len(items), 20)
	if diff != "" {
		t.Fatalf("TestItemIndex: %v", diff)
	}
	for i := 0; i < 20; i++ {
		diff = cmp.Diff(items[i], inserted[i])
		if diff != "" {
			t.Fatalf("TestItemIndex: %v", diff)
		}
	}
}

// inserted: |----------|
// items   :      |--|
func TestItemIndex3(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	defer release()

	inserted, err := insertItemMany(ctx, db)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}

	items, err := domain.ItemIndex(ctx, db, inserted[9].ID, 5)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	diff := cmp.Diff(len(items), 5)
	if diff != "" {
		t.Fatalf("TestItemIndex: %v", diff)
	}
	for i := 0; i < 5; i++ {
		diff = cmp.Diff(items[i], inserted[i+10])
		if diff != "" {
			t.Fatalf("TestItemIndex: %v", diff)
		}
	}
}

// inserted: |----------|
// items   :      |----------|
func TestItemIndex4(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	defer release()

	inserted, err := insertItemMany(ctx, db)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}

	items, err := domain.ItemIndex(ctx, db, inserted[9].ID, 100)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	diff := cmp.Diff(len(items), 10)
	if diff != "" {
		t.Fatalf("TestItemIndex: %v", diff)
	}
	for i := 0; i < 10; i++ {
		diff = cmp.Diff(items[i], inserted[i+10])
		if diff != "" {
			t.Fatalf("TestItemIndex: %v", diff)
		}
	}
}

// inserted: |----------|
// items   :                |----------|
func TestItemIndex5(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	defer release()

	_, err = insertItemMany(ctx, db)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	items, err := domain.ItemIndex(ctx, db, "00000000-0000-0000-0000-000000000000", 100)
	if err != nil {
		t.Fatalf("TestItemIndex: %v", err)
	}
	diff := cmp.Diff(len(items), 0)
	if diff != "" {
		t.Fatalf("TestItemIndex: %v", diff)
	}
}

func TestItemCreate(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemCreate: %v", err)
	}
	defer release()

	item, err := domain.ItemCreate(ctx, db, "test", "https://example.com/", "https://example.com/thumbnail.jpg")
	if err != nil {
		t.Fatalf("TestItemCreate: %v", err)
	}

	diff := cmp.Diff(item, &models.Item{
		Title:     "test",
		URL:       "https://example.com/",
		Thumbnail: "https://example.com/thumbnail.jpg",
	}, itemOpt)
	if diff != "" {
		t.Fatalf("TestItemCreate: %v", diff)
	}

	i, err := models.FindItem(ctx, db, item.ID)
	if err != nil {
		t.Fatalf("TestItemCreate: %v", err)
	}

	diff = cmp.Diff(item, i)
	if diff != "" {
		t.Fatalf("TestItemCreate: %v", diff)
	}
}

func TestItemUpdate(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemCreate: %v", err)
	}
	defer release()

	item, err := insertItem(ctx, db, "test", "https://example.com/", "https://example.com/thumbnail.jpg")
	if err != nil {
		t.Fatalf("TestItemUpdate: %v", err)
	}

	updated, err := domain.ItemUpdate(ctx, db, item.ID, "test2", "https://example2.com/", "https://example2.com/thumbnail.jpg")
	if err != nil {
		t.Fatalf("TestItemUpdate: %v", err)
	}

	diff := cmp.Diff(updated, &models.Item{
		Title:     "test2",
		URL:       "https://example2.com/",
		Thumbnail: "https://example2.com/thumbnail.jpg",
	}, itemOpt)
	if diff != "" {
		t.Fatalf("TestItemUpdate: %v", diff)
	}

	i, err := models.FindItem(ctx, db, updated.ID)
	if err != nil {
		t.Fatalf("TestItemUpdate: %v", err)
	}

	diff = cmp.Diff(updated, i)
	if diff != "" {
		t.Fatalf("TestItemUpdate: %v", diff)
	}
}

func TestItemDelete(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemCreate: %v", err)
	}
	defer release()

	item, err := insertItem(ctx, db, "test", "https://example.com/", "https://example.com/thumbnail.jpg")
	if err != nil {
		t.Fatalf("TestItemDelete: %v", err)
	}

	err = domain.ItemDelete(ctx, db, item.ID)
	if err != nil {
		t.Fatalf("TestItemDelete: %v", err)
	}

	_, err = models.FindItem(ctx, db, item.ID)
	if err != sql.ErrNoRows {
		t.Fatalf("TestItemUpdate: %v", err)
	}
}

func TestItemExport(t *testing.T) {
	ctx := context.Background()

	db, release, err := buildMockDB(ctx)
	if err != nil {
		t.Fatalf("TestItemExport: %v", err)
	}
	defer release()

	inserted, err := insertItemMany(ctx, db)
	if err != nil {
		t.Fatalf("TestItemExport: %v", err)
	}

	items, err := domain.ItemExport(ctx, db)
	if err != nil {
		t.Fatalf("TestItemExport: %v", err)
	}

	for i, item := range items {
		diff := cmp.Diff(item, inserted[19-i])
		if diff != "" {
			t.Fatalf("TestItemExport: %v", diff)
		}
	}
}

func insertItem(ctx context.Context, db domain.DB, title, url, thumbnail string) (*models.Item, error) {
	item := &models.Item{
		Title:     title,
		URL:       url,
		Thumbnail: thumbnail,
	}

	err := item.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return item, nil
}

func insertItemMany(ctx context.Context, db domain.DB) ([]*models.Item, error) {
	items := make([]*models.Item, 20)
	for i := 0; i < 20; i++ {
		title := fmt.Sprintf("example%d", i)
		url := fmt.Sprintf("https://example%d.com/", i)
		thumbnail := fmt.Sprintf("https://example%d.com/thumbnail.jpg", i)
		item, err := insertItem(ctx, db, title, url, thumbnail)
		if err != nil {
			return nil, err
		}
		items[19-i] = item
	}
	return items, nil
}
