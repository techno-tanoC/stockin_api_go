use anyhow::Result;
use std::env;
use std::net::SocketAddr;

use stockin_api::*;

#[tokio::main]
async fn main() -> Result<()> {
    env_logger::init();

    let token = env::var("AUTH_TOKEN")?;

    let database_url = env::var("DATABASE_URL")?;
    let port = env::var("PORT").unwrap_or("3000".to_string()).parse()?;

    check_port(&database_url).await?;

    let state = new_state(&database_url).await?;
    let app = build_app(state, token);

    let addr = SocketAddr::from(([0, 0, 0, 0], port));
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();

    Ok(())
}
