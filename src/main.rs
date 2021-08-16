mod item;

use anyhow::Result;
use axum::prelude::*;
use sqlx::sqlite::SqlitePool;
use std::env;
use std::net::SocketAddr;
use std::sync::Arc;

use crate::item::Item;

#[tokio::main]
async fn main() -> Result<()> {
    let database_url = env::var("DATABASE_URL")?;
    let pool = SqlitePool::connect(&database_url).await?;
    let state = Arc::new(State { pool });

    let item_actions = get(index_item).post(create_item);
    let app = route("/items", item_actions)
        .layer(axum::AddExtensionLayer::new(state));

    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();

    Ok(())
}

struct State {
    pool: SqlitePool,
}

type SharedState = Arc<State>;

async fn index_item(state: extract::Extension<SharedState>) -> response::Json<Vec<Item>> {
    let mut conn = state.pool.acquire().await.unwrap();
    let items = Item::all(&mut conn).await.unwrap();
    response::Json(items)
}

async fn create_item(state: extract::Extension<SharedState>) -> String {
    let mut conn = state.pool.acquire().await.unwrap();
    let id = Item::insert(&mut conn, "1", "one").await.unwrap();
    id.to_string()
}
