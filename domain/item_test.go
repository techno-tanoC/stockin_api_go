package domain_test

import (
	"context"
	"database/sql"
	"stockin/domain"
	"stockin/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var itemOpt = cmpopts.IgnoreFields(models.Item{}, "ID", "CreatedAt", "UpdatedAt")

func TestItemCreate(t *testing.T) {
	rawDB, err := sql.Open("mysql", "root:pass@(db)/test?parseTime=true")
	if err != nil {
		t.Fatalf("TestItemCreate: %v", err)
	}
	defer rawDB.Close()

	ctx := context.Background()

	tx, err := rawDB.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("TestItemCreate: %v", err)
	}
	defer tx.Rollback()
	db := &domain.MockDB{&domain.MockTx{tx}}

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
		t.Fatalf("TestItemCreate: %v", err)
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
	rawDB, err := sql.Open("mysql", "root:pass@(db)/test?parseTime=true")
	if err != nil {
		t.Fatalf("TestItemUpdate: %v", err)
	}
	defer rawDB.Close()

	ctx := context.Background()

	tx, err := rawDB.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("TestItemUpdate: %v", err)
	}
	defer tx.Rollback()
	db := &domain.MockDB{&domain.MockTx{tx}}

	item := &models.Item{
		Title:     "test",
		URL:       "https://example.com/",
		Thumbnail: "https://example.com/thumbnail.jpg",
	}
	err = item.Insert(ctx, db, boil.Infer())
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
	rawDB, err := sql.Open("mysql", "root:pass@(db)/test?parseTime=true")
	if err != nil {
		t.Fatalf("TestItemDelete: %v", err)
	}
	defer rawDB.Close()

	ctx := context.Background()

	tx, err := rawDB.BeginTx(ctx, nil)
	if err != nil {
		t.Fatalf("TestItemDelete: %v", err)
	}
	defer tx.Rollback()
	db := &domain.MockDB{&domain.MockTx{tx}}

	item := &models.Item{
		Title:     "test",
		URL:       "https://example.com/",
		Thumbnail: "https://example.com/thumbnail.jpg",
	}
	err = item.Insert(ctx, db, boil.Infer())
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
