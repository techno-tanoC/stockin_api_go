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
}
