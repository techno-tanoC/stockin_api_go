use axum::extract;
use chrono::prelude::*;
use serde::{Deserialize, Serialize};
use std::convert::From;

use crate::SharedState;
use crate::repo;
use super::response::*;

#[derive(Debug, Clone, Serialize)]
pub struct Item {
    id: i64,
    title: String,
    url: String,
    created_at: NaiveDateTime,
    updated_at: NaiveDateTime,
}

impl From<repo::Item> for Item {
    fn from(item: repo::Item) -> Self {
        Item {
            id: item.id,
            title: item.title,
            url: item.url,
            created_at: item.created_at,
            updated_at: item.updated_at,
        }
    }
}

#[derive(Debug, Clone, Deserialize)]
pub struct Id {
    item_id: i64,
}

#[derive(Debug, Clone, Deserialize)]
pub struct Params {
    title: String,
    url: String,
}

pub async fn index_item(state: extract::Extension<SharedState>) -> Result<Vec<Item>> {
    let items = repo::Item::all(&state.pool).await.map_err(internal_error)?;
    ok(items.into_iter().map(|i| i.into()).rev().collect())
}

pub async fn create_item(params: extract::Json<Params>, state: extract::Extension<SharedState>) -> Result<Item> {
    let mut conn = state.pool.acquire().await.map_err(internal_error)?;
    let id = repo::Item::insert(&mut conn, &params.title, &params.url).await.map_err(internal_error)?;
    let item = repo::Item::find(&mut conn, id).await.map_err(internal_error)?;
    ok(item.into())
}

pub async fn update_item(id: extract::Path<Id>, params: extract::Json<Params>, state: extract::Extension<SharedState>) -> Result<Item> {
    let mut conn = state.pool.acquire().await.map_err(internal_error)?;
    let _ = repo::Item::update(&mut conn, id.item_id, &params.title, &params.url).await.map_err(internal_error)?;
    let item = repo::Item::find(&mut conn, id.item_id).await.map_err(internal_error)?;
    ok(item.into())
}

pub async fn delete_item(id: extract::Path<Id>, state: extract::Extension<SharedState>) -> Result<()> {
    let _ = repo::Item::delete(&state.pool, id.item_id).await.map_err(internal_error)?;
    ok(())
}
