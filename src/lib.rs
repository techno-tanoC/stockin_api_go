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

pub fn build_app(state: SharedState) -> Router<BoxRoute> {
    let item_actions = Router::new()
        .route("/", get(handler::get).post(handler::create))
        .route("/:item_id", get(handler::find).put(handler::update).delete(handler::delete));

    Router::new()
        .nest("/items", item_actions)
        .layer(axum::AddExtensionLayer::new(state))
        .layer(tower_http::trace::TraceLayer::new_for_http())
        .boxed()
}
