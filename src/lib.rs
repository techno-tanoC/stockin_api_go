pub mod handler;
pub mod repo;

use anyhow::Result;
use axum::routing::*;
use sqlx::mysql::MySqlPool;
use std::sync::Arc;

pub struct State {
    pub pool: MySqlPool,
}

pub type SharedState = Arc<State>;

pub async fn new_state(url: &str) -> Result<SharedState> {
    let pool = MySqlPool::connect_lazy(url)?;
    Ok(Arc::new(State { pool }))
}

#[derive(Debug, Clone)]
pub struct Bearer {
    token: String,
}

pub fn build_app(state: SharedState, token: String) -> Router {
    let item_actions = Router::new()
        .route("/", get(handler::find_by_range).post(handler::create))
        .route("/:item_id", get(handler::find).put(handler::update).delete(handler::delete))
        .route("/:item_id/archive", patch(handler::archive))
        .route("/:item_id/unarchive", patch(handler::unarchive));

    let title_actions = Router::new()
        .route("/query", post(handler::query));

    Router::new()
        .nest("/items", item_actions)
        .nest("/title", title_actions)
        .layer(axum::AddExtensionLayer::new(state))
        .layer(axum::AddExtensionLayer::new(Bearer { token }))
        .layer(tower_http::trace::TraceLayer::new_for_http())
}

#[cfg(test)]
mod tests {
    use super::*;
    use once_cell::sync::Lazy;

    pub static POOL: Lazy<MySqlPool> = Lazy::new(|| {
        let url = std::env::var("TEST_DATABASE_URL").unwrap();
        MySqlPool::connect_lazy(&url).unwrap()
    });
}
