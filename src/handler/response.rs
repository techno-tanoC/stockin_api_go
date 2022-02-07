use axum::{
    http::StatusCode,
    response,
};
use serde::Serialize;

#[derive(Debug, Clone, Serialize)]
pub struct Success<T> {
    data: T,
}

#[derive(Debug, Clone, Serialize)]
pub struct Message {
    message: String,
}

pub type Result<T> = std::result::Result<(StatusCode, response::Json<Success<T>>), (StatusCode, response::Json<Message>)>;

pub fn ok<T: Serialize>(data: T) -> Result<T> {
    Ok((StatusCode::OK, response::Json(Success{ data })))
}

pub type RawResult<T> = std::result::Result<(StatusCode, response::Json<T>), (StatusCode, response::Json<Message>)>;

pub fn ok_raw<T: Serialize>(data: T) -> RawResult<T> {
    Ok((StatusCode::OK, response::Json(data)))
}

pub fn client_error<E>(err: E) -> (StatusCode, response::Json<Message>)
where
    E: Into<anyhow::Error>,
{
    (StatusCode::BAD_REQUEST, response::Json(Message{ message: err.into().to_string() }))
}

pub fn server_error<E>(_: E) -> (StatusCode, response::Json<Message>) {
    (StatusCode::INTERNAL_SERVER_ERROR, response::Json(Message{ message: "internal server error".to_string()}))
}
