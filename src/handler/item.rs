use axum::prelude::*;
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

pub async fn index_item(state: extract::Extension<SharedState>) -> Result<Vec<Item>> {
    let items = repo::Item::all(&state.pool).await.map_err(internal_error)?;
    ok(items.into_iter().map(|i| i.into()).collect())
}

#[derive(Debug, Clone, Deserialize)]
pub struct NewItem {
    title: String,
    url: String,
}

#[derive(Debug, Clone, Serialize)]
pub struct Id {
    id: i64,
}

pub async fn create_item(new_item: extract::Json<NewItem>, state: extract::Extension<SharedState>) -> Result<Id> {
    let id = repo::Item::insert(&state.pool, &new_item.title, &new_item.url).await.map_err(internal_error)?;
    ok(Id { id })
}
