use axum::{
    http::StatusCode,
    response,
};
use serde::Serialize;

pub type Result<T> = std::result::Result<(StatusCode, response::Json<Success<T>>), (StatusCode, response::Json<Message>)>;

#[derive(Debug, Clone, Serialize)]
pub struct Success<T> {
    data: T,
}

#[derive(Debug, Clone, Serialize)]
pub struct Message {
    message: String,
}

pub fn ok<T: Serialize>(data: T) -> Result<T> {
    Ok((StatusCode::OK, response::Json(Success{ data })))
}

pub fn client_error<E>(_: E) -> (StatusCode, response::Json<Message>) {
    (StatusCode::BAD_REQUEST, response::Json(Message{ message: "bad request error".to_string()}))
}

pub fn server_error<E>(_: E) -> (StatusCode, response::Json<Message>) {
    (StatusCode::INTERNAL_SERVER_ERROR, response::Json(Message{ message: "internal server error".to_string()}))
}
