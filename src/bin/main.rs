use anyhow::Result;
use axum::{
    handler::*,
    Router,
};
use std::env;
use std::net::SocketAddr;

use stockin_api::*;

#[tokio::main]
async fn main() -> Result<()> {
    let database_url = env::var("DATABASE_URL")?;
    let state = new_state(&database_url).await?;

    let item_actions = get(handler::index_item).post(handler::create_item);
    let app = Router::new().route("/items", item_actions)
        .layer(axum::AddExtensionLayer::new(state));

    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();

    Ok(())
}
