pub mod handler;
pub mod repo;

use anyhow::Result;
use axum::{
    handler::*,
    routing::{BoxRoute, Router},
};
use sqlx::mysql::MySqlPool;
use std::sync::Arc;

pub struct State {
    pub pool: MySqlPool,
}

pub type SharedState = Arc<State>;

pub async fn new_state(url: &str) -> Result<SharedState> {
    let pool = MySqlPool::connect(url).await?;
    Ok(Arc::new(State { pool }))
}

#[derive(Debug, Clone)]
pub struct Bearer {
    token: String,
}

pub fn build_app(state: SharedState, token: String) -> Router<BoxRoute> {
    let item_actions = Router::new()
        .route("/", get(handler::find_by_range).post(handler::create))
        .route("/:item_id", get(handler::find).patch(handler::update).delete(handler::delete))
        .route("/:item_id/archive", patch(handler::archive))
        .route("/:item_id/unarchive", patch(handler::unarchive));

    Router::new()
        .nest("/items", item_actions)
        .layer(axum::AddExtensionLayer::new(state))
        .layer(axum::AddExtensionLayer::new(Bearer { token }))
        .layer(tower_http::trace::TraceLayer::new_for_http())
        .boxed()
}
