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
    let item_actions = Router::new()
        .route("/", get(handler::index_item).post(handler::create_item))
        .route("/:item_id", put(handler::update_item).delete(handler::delete_item));

    Router::new()
        .nest("/items", item_actions)
        .layer(axum::AddExtensionLayer::new(state))
        .boxed()
}
