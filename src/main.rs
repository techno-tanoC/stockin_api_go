mod handler;
mod repo;

use anyhow::Result;
use axum::prelude::*;
use sqlx::sqlite::SqlitePool;
use std::env;
use std::net::SocketAddr;
use std::sync::Arc;

#[tokio::main]
async fn main() -> Result<()> {
    let database_url = env::var("DATABASE_URL")?;
    let pool = SqlitePool::connect(&database_url).await?;
    let state = Arc::new(State { pool });

    let item_actions = get(handler::index_item).post(handler::create_item);
    let app = route("/items", item_actions)
        .layer(axum::AddExtensionLayer::new(state));

    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();

    Ok(())
}

pub struct State {
    pool: SqlitePool,
}

pub type SharedState = Arc<State>;
