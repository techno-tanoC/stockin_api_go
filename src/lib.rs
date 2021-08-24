pub mod handler;
pub mod repo;

use anyhow::Result;
use axum::{
    handler::*,
    routing::{BoxRoute, Router},
};
use sqlx::sqlite::SqlitePool;
use std::sync::Arc;

pub struct State {
    pub pool: SqlitePool,
}

pub type SharedState = Arc<State>;

pub async fn new_state(url: &str) -> Result<SharedState> {
    let pool = SqlitePool::connect(url).await?;
    Ok(Arc::new(State { pool }))
}

pub fn build_app(state: SharedState) -> Router<BoxRoute> {
    let item_actions = get(handler::index_item).post(handler::create_item);

    Router::new()
        .route("/items", item_actions)
        .layer(axum::AddExtensionLayer::new(state))
        .boxed()
}
