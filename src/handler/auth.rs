use axum::{
    async_trait,
    extract::{Extension, FromRequest, RequestParts},
    http::{header, StatusCode},
};

use crate::Bearer;

#[derive(Debug, Clone)]
pub struct UserId;

const UNAUTHORIZED: (StatusCode, &'static str) = (StatusCode::UNAUTHORIZED, "Unauthorized");

#[async_trait]
impl<B> FromRequest<B> for UserId
where B: Send,
{
    type Rejection = (StatusCode, &'static str);

    async fn from_request(req: &mut RequestParts<B>) -> Result<Self, Self::Rejection> {
        let bearer = Extension::<Bearer>::from_request(req).await.expect("`Bearer` extension missing");

        let headers = req.headers().expect("other extractor taken headers");

        let value = headers.get(header::AUTHORIZATION)
            .ok_or(UNAUTHORIZED)?
            .to_str()
            .map_err(|_| UNAUTHORIZED)?
            .to_string();
        let token = value.trim_start_matches("Bearer ");

        if token == &bearer.token {
            Ok(UserId)
        } else {
            Err(UNAUTHORIZED)
        }
    }
}
