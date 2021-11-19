use anyhow::Result;
use chrono::prelude::*;

use super::Exe;

#[derive(Debug, PartialEq, Clone)]
pub struct Item {
    pub id: i64,
    pub title: String,
    pub url: String,
    pub is_archived: bool,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

impl Item {
    pub async fn find(exe: impl Exe<'_>, id: u64) -> Result<Option<Item>> {
        let item = sqlx::query_as!(
            Item,
            r#"
            SELECT id, title, url, is_archived as "is_archived: _", created_at, updated_at
            FROM items
            WHERE id = ?
            "#,
            id
        )
        .fetch_optional(exe)
        .await?;

        Ok(item)
    }

    pub async fn all(exe: impl Exe<'_>) -> Result<Vec<Item>> {
        let items = sqlx::query_as!(
            Item,
            r#"
            SELECT id, title, url, is_archived as "is_archived: _", created_at, updated_at
            FROM items
            "#
        )
        .fetch_all(exe)
        .await?;

        Ok(items)
    }

    pub async fn find_by_range(exe: impl Exe<'_>, before: u64, count: u64) -> Result<Vec<Item>> {
        let items = sqlx::query_as!(
            Item,
            r#"
            SELECT id, title, url, is_archived as "is_archived: _", created_at, updated_at
            FROM items
            WHERE id < ?
            ORDER BY id DESC
            LIMIT ?
            "#,
            before,
            count
        )
        .fetch_all(exe)
        .await?;

        Ok(items)
    }

    pub async fn insert(exe: impl Exe<'_>, title: &str, url: &str) -> Result<u64> {
        let id = sqlx::query!(
            r#"
            INSERT INTO items (title, url)
            VALUES (?, ?)
            "#,
            title,
            url
        )
        .execute(exe)
        .await?
        .last_insert_id();

        Ok(id)
    }

    pub async fn update(exe: impl Exe<'_>, id: u64, title: &str, url: &str) -> Result<()> {
        let _ = sqlx::query!(
            r#"
            UPDATE items
            SET title = ?, url = ?
            WHERE id = ?
            "#,
            title,
            url,
            id
        )
        .execute(exe)
        .await?;

        Ok(())
    }

    pub async fn update_is_archived(exe: impl Exe<'_>, id: u64, is_archived: bool) -> Result<()> {
        let _ = sqlx::query!(
            r#"
            UPDATE items
            SET is_archived = ?
            WHERE id = ?
            "#,
            is_archived,
            id
        )
        .execute(exe)
        .await?;

        Ok(())
    }

    pub async fn delete(exe: impl Exe<'_>, id: u64) -> Result<()> {
        let _ = sqlx::query!(
            r#"
            DELETE FROM items
            WHERE id = ?
            "#,
            id
        )
        .execute(exe)
        .await?;

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::tests::*;

    #[tokio::test]
    async fn test_find() {
        let mut tx = POOL.begin().await.unwrap();

        let title = "example";
        let url = "https://example.com/";
        let id = Item::insert(&mut tx, title, url).await.unwrap();

        {
            let item = Item::find(&mut tx, id).await.unwrap().unwrap();
            assert_eq!(item.title, title.to_string());
            assert_eq!(item.url, url.to_string());
        }

        {
            let item = Item::find(&mut tx, id + 1).await.unwrap();
            assert_eq!(item, None)
        }
    }

    #[tokio::test]
    async fn test_all() {
        let mut tx = POOL.begin().await.unwrap();

        let mut ids = vec![];
        for i in 0..10 {
            let title = format!("example {}", i);
            let url = format!("https://example.com/{}", i);
            let id = Item::insert(&mut tx, &title, &url).await.unwrap();
            ids.push(id);
        }

        ids.reverse();

        {
            let items = Item::all(&mut tx).await.unwrap();
            assert_eq!(items.len(), 10);
            assert_eq!(items.first().unwrap().title, "example 0");
            assert_eq!(items.last().unwrap().title, "example 9");
        }
    }

    #[tokio::test]
    async fn test_find_by_range() {
        let mut tx = POOL.begin().await.unwrap();

        let mut ids = vec![];
        for i in 0..10 {
            let title = format!("example {}", i);
            let url = format!("https://example.com/{}", i);
            let id = Item::insert(&mut tx, &title, &url).await.unwrap();
            ids.push(id);
        }

        ids.reverse();

        {
            let items = Item::find_by_range(&mut tx, u64::MAX, 5).await.unwrap();
            assert_eq!(items.len(), 5);
            assert_eq!(items.first().unwrap().title, "example 9");
            assert_eq!(items.last().unwrap().title, "example 5");
        }

        {
            let items = Item::find_by_range(&mut tx, u64::MAX, 20).await.unwrap();
            assert_eq!(items.len(), 10);
            assert_eq!(items.first().unwrap().title, "example 9");
            assert_eq!(items.last().unwrap().title, "example 0");
        }

        {
            let items = Item::find_by_range(&mut tx, ids[3], 3).await.unwrap();
            assert_eq!(items.len(), 3);
            assert_eq!(items.first().unwrap().title, "example 5");
            assert_eq!(items.last().unwrap().title, "example 3");
        }

        {
            let items = Item::find_by_range(&mut tx, 0, 10).await.unwrap();
            assert_eq!(items.len(), 0);
        }
    }

    #[tokio::test]
    async fn test_insert() {
        let mut tx = POOL.begin().await.unwrap();

        let title = "example";
        let url = "https://example.com/";
        let id = Item::insert(&mut tx, title, url).await;

        assert!(id.is_ok());
    }

    #[tokio::test]
    async fn test_update() {
        let mut tx = POOL.begin().await.unwrap();

        let title = "example";
        let url = "https://example.com/";
        let id = Item::insert(&mut tx, title, url).await.unwrap();

        let new_title = "new";
        let new_url = "https://new.com/";
        Item::update(&mut tx, id, new_title, new_url).await.unwrap();

        let item = Item::find(&mut tx, id).await.unwrap().unwrap();
        assert_eq!(item.title, new_title.to_string());
        assert_eq!(item.url, new_url.to_string());
    }

    #[tokio::test]
    async fn test_update_is_archived() {
        let mut tx = POOL.begin().await.unwrap();

        let title = "example";
        let url = "https://example.com/";
        let id = Item::insert(&mut tx, title, url).await.unwrap();

        {
            Item::update_is_archived(&mut tx, id, true).await.unwrap();

            let item = Item::find(&mut tx, id).await.unwrap().unwrap();
            assert!(item.is_archived);
        }

        {
            Item::update_is_archived(&mut tx, id, false).await.unwrap();

            let item = Item::find(&mut tx, id).await.unwrap().unwrap();
            assert!(!item.is_archived);
        }
    }

    #[tokio::test]
    async fn test_delete() {
        let mut tx = POOL.begin().await.unwrap();

        let title = "example";
        let url = "https://example.com/";
        let id = Item::insert(&mut tx, title, url).await.unwrap();

        Item::delete(&mut tx, id).await.unwrap();

        let item = Item::find(&mut tx, id).await.unwrap();
        assert_eq!(item, None);
    }
}
