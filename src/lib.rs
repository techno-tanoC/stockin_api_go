pub mod handler;
pub mod repo;

use anyhow::Result;
use axum::routing::*;
use sqlx::mysql::MySqlPool;
use std::sync::Arc;

use handler::*;

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
    let item_routes = Router::new()
        .route("/", get(item::find_by_range).post(item::create))
        .route("/:item_id", get(item::find).put(item::update).delete(item::delete))
        .route("/:item_id/archive", patch(item::archive))
        .route("/:item_id/unarchive", patch(item::unarchive));

    let title_routes = Router::new()
        .route("/query", post(title::query));

    Router::new()
        .nest("/items", item_routes)
        .nest("/title", title_routes)
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
