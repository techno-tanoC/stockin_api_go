use axum::extract;
use serde::{Deserialize, Serialize};

use crate::handler::response::*;
use crate::handler::auth::UserId;

#[derive(Debug, Clone, Serialize)]
pub struct Title {
    title: String,
}

#[derive(Debug, Clone, Deserialize)]
pub struct Url {
    url: String,
}

impl Url {
    fn validate(&self) -> anyhow::Result<()> {
        url::Url::parse(&self.url).map_err(|_| anyhow::anyhow!("url is invalid"))?;
        Ok(())
    }
}

pub async fn query(url: extract::Json<Url>, _: UserId) -> Result<Option<Title>> {
    url.validate().map_err(client_error)?;

    async fn fetch_html(u: &str) -> anyhow::Result<String> {
        let html = reqwest::get(u)
            .await?
            .text()
            .await?;
        Ok(html)
    }

    // TODO: timeout
    let html = fetch_html(&url.url)
        .await
        .map_err(|_| client_error(anyhow::anyhow!("could not fetch html")))?;

    let doc = scraper::Html::parse_document(&html);
    let selector = scraper::Selector::parse("html > head > title").unwrap();
    let opt = doc
        .select(&selector)
        .next()
        .and_then(|e|
            e.text().next())
        .map(|title|
            Title { title: title.to_string() }
        );

    ok(opt)
}
