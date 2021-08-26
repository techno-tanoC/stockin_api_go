use anyhow::Result;
use chrono::prelude::*;

use super::Exe;

#[derive(Debug, Clone)]
pub struct Item {
    pub id: i64,
    pub title: String,
    pub url: String,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

impl Item {
    pub async fn find(exe: impl Exe<'_>, id: i64) -> Result<Item> {
        let item = sqlx::query_as!(
            Item,
            r#"
            SELECT *
            FROM items
            WHERE id = ?1
            "#,
            id
        )
        .fetch_one(exe)
        .await?;

        Ok(item)
    }

    pub async fn all(exe: impl Exe<'_>) -> Result<Vec<Item>> {
        let items = sqlx::query_as!(
            Item,
            r#"
            SELECT *
            FROM items
            "#
        )
        .fetch_all(exe)
        .await?;

        Ok(items)
    }

    pub async fn insert(exe: impl Exe<'_>, title: &str, url: &str) -> Result<i64> {
        let id = sqlx::query!(
            r#"
            INSERT INTO items (title, url)
            VALUES (?1, ?2)
            "#,
            title,
            url
        )
        .execute(exe)
        .await?
        .last_insert_rowid();

        Ok(id)
    }

    pub async fn update(exe: impl Exe<'_>, id: i64, title: &str, url: &str) -> Result<()> {
        let _ = sqlx::query!(
            r#"
            UPDATE items
            SET title = ?1, url = ?2, updated_at = CURRENT_TIMESTAMP
            WHERE id = ?3
            "#,
            title,
            url,
            id
        )
        .execute(exe)
        .await?;

        Ok(())
    }

    pub async fn delete(exe: impl Exe<'_>, id: i64) -> Result<()> {
        let _ = sqlx::query!(
            r#"
            DELETE FROM items
            WHERE id = ?1
            "#,
            id
        )
        .execute(exe)
        .await?;

        Ok(())
    }
}
