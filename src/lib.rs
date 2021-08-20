pub mod handler;
pub mod repo;

use anyhow::Result;
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
