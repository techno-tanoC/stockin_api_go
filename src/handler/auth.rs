use axum::{
    async_trait,
    extract::{Extension, FromRequest, RequestParts},
    http::{header, StatusCode},
};

use crate::Bearer;

#[derive(Debug, Clone)]
pub struct UserId;

#[async_trait]
impl<B> FromRequest<B> for UserId
where B: Send,
{
    type Rejection = StatusCode;

    async fn from_request(req: &mut RequestParts<B>) -> Result<Self, Self::Rejection> {
        let bearer = Extension::<Bearer>::from_request(req).await.expect("`Bearer` extension missing");

        let token_option = req
            .headers()
            .get(header::AUTHORIZATION)
            .and_then(|value| value.to_str().ok())
            .map(|value| value.trim_start_matches("Bearer "));

        if let Some(token) = token_option {
            if token == bearer.token {
                return Ok(Self)
            }
        }

        Err(StatusCode::UNAUTHORIZED)
    }
}
