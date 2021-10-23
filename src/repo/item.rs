use anyhow::Result;
use chrono::prelude::*;

use super::Exe;

#[derive(Debug, Clone)]
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
