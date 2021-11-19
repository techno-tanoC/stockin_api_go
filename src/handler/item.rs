use axum::extract;
use chrono::prelude::*;
use serde::{Deserialize, Serialize};
use std::convert::From;

use crate::SharedState;
use crate::repo;
use super::response::*;
use super::auth::UserId;

#[derive(Debug, Clone, Serialize)]
pub struct Item {
    id: i64,
    title: String,
    url: String,
    is_archived: bool,
    created_at: NaiveDateTime,
    updated_at: NaiveDateTime,
}

impl From<repo::Item> for Item {
    fn from(item: repo::Item) -> Self {
        Item {
            id: item.id,
            title: item.title,
            url: item.url,
            is_archived: item.is_archived,
            created_at: item.created_at,
            updated_at: item.updated_at,
        }
    }
}

#[derive(Debug, Clone, Deserialize)]
pub struct Range {
    #[serde(default = "default_before")]
    before: u64,
    #[serde(default = "default_size")]
    size: u64,
}

fn default_before() -> u64 {
    std::u64::MAX
}

fn default_size() -> u64 {
    50
}

#[derive(Debug, Clone, Deserialize)]
pub struct Id {
    item_id: u64,
}

#[derive(Debug, Clone, Deserialize)]
pub struct Params {
    title: String,
    url: String,
}

impl Params {
    fn validate(&self) -> anyhow::Result<()> {
        if self.title.len() > 1024 {
            Err(anyhow::anyhow!("title length is over 1024"))?;
        }

        if self.url.len() > 4096 {
            Err(anyhow::anyhow!("url length is over 4096"))?;
        }

        url::Url::parse(&self.url).map_err(|_| anyhow::anyhow!("url is invalid"))?;

        Ok(())
    }
}

pub async fn find(id: extract::Path<Id>, state: extract::Extension<SharedState>, _: UserId) -> Result<Option<Item>> {
    let option = repo::Item::find(&state.pool, id.item_id).await.map_err(server_error)?;
    ok(option.map(|i| i.into()))
}

pub async fn find_by_range(range: extract::Query<Range>, state: extract::Extension<SharedState>, _: UserId) -> Result<Vec<Item>> {
    if range.size > 100 {
        Err(client_error(anyhow::anyhow!("range size is over 100")))?;
    }

    let items = repo::Item::find_by_range(&state.pool, range.before, range.size).await.map_err(server_error)?;
    ok(items.into_iter().map(|i| i.into()).collect())
}

pub async fn create(params: extract::Json<Params>, state: extract::Extension<SharedState>, _: UserId) -> Result<Option<Item>> {
    params.validate().map_err(client_error)?;

    let mut conn = state.pool.acquire().await.map_err(server_error)?;
    let id = repo::Item::insert(&mut conn, &params.title, &params.url).await.map_err(server_error)?;
    let option = repo::Item::find(&mut conn, id).await.map_err(server_error)?;
    ok(option.map(|i| i.into()))
}

pub async fn update(id: extract::Path<Id>, params: extract::Json<Params>, state: extract::Extension<SharedState>, _: UserId) -> Result<Option<Item>> {
    params.validate().map_err(client_error)?;

    let mut conn = state.pool.acquire().await.map_err(server_error)?;
    let _ = repo::Item::update(&mut conn, id.item_id, &params.title, &params.url).await.map_err(server_error)?;
    let option = repo::Item::find(&mut conn, id.item_id).await.map_err(server_error)?;
    ok(option.map(|i| i.into()))
}

pub async fn archive(id: extract::Path<Id>, state: extract::Extension<SharedState>, _: UserId) -> Result<Option<Item>> {
    let mut conn = state.pool.acquire().await.map_err(server_error)?;
    let _ = repo::Item::update_is_archived(&state.pool, id.item_id, true).await.map_err(server_error)?;
    let option = repo::Item::find(&mut conn, id.item_id).await.map_err(server_error)?;
    ok(option.map(|i| i.into()))
}

pub async fn unarchive(id: extract::Path<Id>, state: extract::Extension<SharedState>, _: UserId) -> Result<Option<Item>> {
    let mut conn = state.pool.acquire().await.map_err(server_error)?;
    let _ = repo::Item::update_is_archived(&state.pool, id.item_id, false).await.map_err(server_error)?;
    let option = repo::Item::find(&mut conn, id.item_id).await.map_err(server_error)?;
    ok(option.map(|i| i.into()))
}

pub async fn delete(id: extract::Path<Id>, state: extract::Extension<SharedState>, _: UserId) -> Result<()> {
    let _ = repo::Item::delete(&state.pool, id.item_id).await.map_err(server_error)?;
    ok(())
}
