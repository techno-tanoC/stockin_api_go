mod item;

use anyhow::Result;
use axum::prelude::*;
use sqlx::sqlite::SqlitePool;
use std::env;
use std::net::SocketAddr;

use crate::item::Item;

#[tokio::main]
async fn main() -> Result<()> {
    let database_url = env::var("DATABASE_URL")?;
    let pool = SqlitePool::connect(&database_url).await?;
    let mut conn = pool.acquire().await?;

    Item::insert(&mut conn, "title", "url").await?;

    let items = Item::all(&mut conn).await?;
    for item in items {
        dbg!(item);
    }

    Ok(())
}
