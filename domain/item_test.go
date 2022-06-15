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
