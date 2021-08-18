use axum::prelude::*;
use chrono::prelude::*;
use serde::Serialize;
use std::convert::From;

use crate::SharedState;
use crate::repo;

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

pub async fn index_item(state: extract::Extension<SharedState>) -> response::Json<Vec<Item>> {
    let mut conn = state.pool.acquire().await.unwrap();
    let items = repo::Item::all(&mut conn).await.unwrap();
    response::Json(items.into_iter().map(|i| i.into()).collect())
}

#[derive(Debug, Clone, Serialize)]
pub struct Id {
    id: i64,
}


pub async fn create_item(state: extract::Extension<SharedState>) -> response::Json<Id> {
    let mut conn = state.pool.acquire().await.unwrap();
    let id = repo::Item::insert(&mut conn, "1", "one").await.unwrap();
    response::Json(Id { id })
}
