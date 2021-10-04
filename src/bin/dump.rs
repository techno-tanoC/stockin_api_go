use anyhow::Result;
use chrono::prelude::*;
use serde::Serialize;
use sqlx::mysql::MySqlPool;
use std::convert::From;
use std::env;

use stockin_api::repo;

#[tokio::main]
async fn main() -> Result<()> {
    let database_url = env::var("DATABASE_URL")?;
    let pool = MySqlPool::connect(&database_url).await?;

    let items: Vec<Item> = repo::Item::all(&pool)
        .await?
        .into_iter()
        .map(|i| i.into())
        .collect();

    println!("{}", serde_json::to_string_pretty(&items)?);

    Ok(())
}

#[derive(Debug, Clone, Serialize)]
pub struct Item {
    id: i64,
    title: String,
    url: String,
    created_at: NaiveDateTime,
    updated_at: NaiveDateTime,
}

impl From<repo::Item> for Item {
    fn from(item: repo::Item) -> Self {
        Item {
            id: item.id,
            title: item.title,
            url: item.url,
            created_at: item.created_at,
            updated_at: item.updated_at,
        }
    }
}
