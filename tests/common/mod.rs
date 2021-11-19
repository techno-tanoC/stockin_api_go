use axum::http;
use axum::routing::Router;
use futures::future::FutureExt as _;
use once_cell::sync::Lazy;
use sqlx::MySqlPool;

use stockin_api::*;

pub async fn build_test_app() -> Router {
    let url = std::env::var("TEST_DATABASE_URL").unwrap();
    let state = new_state(&url).await.unwrap();
    build_app(state, "test".to_string())
}

pub fn authed_request() -> http::request::Builder {
    http::Request::builder()
        .header(http::header::CONTENT_TYPE, "application/json")
        .header(http::header::AUTHORIZATION, "Bearer test")
}

pub async fn reset_database() {
    let url = std::env::var("TEST_DATABASE_URL").unwrap();
    let pool = MySqlPool::connect(&url).await.unwrap();
    sqlx::query!(r#"DELETE FROM items"#).execute(&pool).await.unwrap();
}

pub async fn run_test<T>(test: T)
where
    T: std::future::Future + Send + 'static,
    T::Output: Send + 'static,
{
    let result = std::panic::AssertUnwindSafe(test).catch_unwind().await;

    reset_database().await;

    if let Err(err) = result {
        std::panic::resume_unwind(err);
    }
}

pub static POOL: Lazy<MySqlPool> = Lazy::new(|| {
    let url = std::env::var("TEST_DATABASE_URL").unwrap();
    MySqlPool::connect_lazy(&url).unwrap()
});
