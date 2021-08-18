use anyhow::Result;
use chrono::prelude::*;

use super::Conn;

#[derive(Debug, Clone)]
pub struct Item {
    pub id: i64,
    pub title: String,
    pub url: String,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

impl Item {
    pub async fn all(pool: impl Conn<'_>) -> Result<Vec<Item>>
    {
        let items = sqlx::query_as!(
            Item,
            r#"
            SELECT *
            FROM items
            "#
        )
        .fetch_all(pool)
        .await?;

        Ok(items)
    }

    pub async fn insert(conn: impl Conn<'_>, title: &str, url: &str) -> Result<i64> {
        let id = sqlx::query!(
            r#"
            INSERT INTO items (title, url)
            VALUES (?1, ?2)
            "#,
            title,
            url
        )
        .execute(conn)
        .await?
        .last_insert_rowid();

        Ok(id)
    }
}
