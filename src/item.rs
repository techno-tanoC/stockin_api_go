use anyhow::Result;
use chrono::prelude::*;
use serde::Serialize;
use sqlx::sqlite::SqliteConnection;

#[derive(Debug, Clone, Serialize)]
pub struct Item {
    pub id: i64,
    pub title: String,
    pub url: String,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

impl Item {
    pub async fn all(conn: &mut SqliteConnection) -> Result<Vec<Item>> {
        let items = sqlx::query_as!(
            Item,
            r#"
            SELECT *
            FROM items
            "#
        )
        .fetch_all(conn)
        .await?;

        Ok(items)
    }

    pub async fn insert(conn: &mut SqliteConnection, title: &str, url: &str) -> Result<i64> {
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
