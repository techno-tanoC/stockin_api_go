mod common;

use assert_json::assert_json;
use axum::body::Body;
use axum::http::{Method, StatusCode};
use serde_json::json;
use tower::ServiceExt;

use common::*;
use stockin_api::*;

#[tokio::test]
async fn test_item_find() {
    run_test(async {
        let title = "example";
        let url = "https://example.com/";
        let id = repo::Item::insert(&*POOL, title, url).await.unwrap();

        {
            let req = authed_request()
                .uri(format!("/items/{}", id))
                .body(Body::empty())
                .unwrap();

            let res = build_test_app()
                .await
                .oneshot(req)
                .await
                .unwrap();

            assert_eq!(res.status(), StatusCode::OK);

            let body = hyper::body::to_bytes(res.into_body()).await.unwrap();
            let json = std::str::from_utf8(&body).unwrap();
            assert_json!(
                json,
                {
                    "data": {
                        "id": id,
                        "title": title,
                        "url": url,
                    },
                }
            );
        }

        {
            let req = authed_request()
                .uri(format!("/items/{}", id + 1))
                .body(Body::empty())
                .unwrap();

            let res = build_test_app()
                .await
                .oneshot(req)
                .await
                .unwrap();

            assert_eq!(res.status(), StatusCode::OK);

            let body = hyper::body::to_bytes(res.into_body()).await.unwrap();
            let json = std::str::from_utf8(&body).unwrap();
            assert_json!(
                json,
                {
                    "data": assert_json::validators::null(),
                }
            );
        }
    }).await;
}

#[tokio::test]
async fn test_item_find_by_range() {
    run_test(async {
        {
            let req = authed_request()
                .uri("/items")
                .body(Body::empty())
                .unwrap();

            let res = build_test_app()
                .await
                .oneshot(req)
                .await
                .unwrap();

            assert_eq!(res.status(), StatusCode::OK);

            let body = hyper::body::to_bytes(res.into_body()).await.unwrap();
            let json = std::str::from_utf8(&body).unwrap();
            assert_json!(
                json,
                {
                    "data": [],
                }
            );
        }

        {
            for _ in 0..10 {
                repo::Item::insert(&*POOL, "example", "https://example.com/").await.unwrap();
            }

            let req = authed_request()
                .uri("/items")
                .body(Body::empty())
                .unwrap();

            let res = build_test_app()
                .await
                .oneshot(req)
                .await
                .unwrap();

            assert_eq!(res.status(), StatusCode::OK);

            let body = hyper::body::to_bytes(res.into_body()).await.unwrap();
            let json: serde_json::Value = serde_json::from_slice(&body).unwrap();
            let data = json.get("data").unwrap();
            assert_eq!(data.as_array().unwrap().len(), 10);
        }
    }).await;
}

#[tokio::test]
async fn test_item_create() {
    run_test(async {
        let title = "example";
        let url = "https://example.com/";

        let body = serde_json::to_vec(
            &json!({
                "title": title,
                "url": url
            })
        ).unwrap();

        let req = authed_request()
            .method(Method::POST)
            .uri("/items")
            .body(Body::from(body))
            .unwrap();

        let res = build_test_app()
            .await
            .oneshot(req)
            .await
            .unwrap();

        assert_eq!(res.status(), StatusCode::OK);

        let body = hyper::body::to_bytes(res.into_body()).await.unwrap();
        let json = std::str::from_utf8(&body).unwrap();
        assert_json!(
            json,
            {
                "data": {
                    "title": title,
                    "url": url,
                    "is_archived": false,
                }
            }
        );
    }).await;
}

#[tokio::test]
async fn test_item_update() {
    run_test(async {
        let id = repo::Item::insert(&*POOL, "example", "https://example.com/").await.unwrap();

        let new_title = "sample";
        let new_url = "https://sample.com/";

        let body = serde_json::to_vec(
            &json!({
                "title": new_title,
                "url": new_url
            })
        ).unwrap();

        let req = authed_request()
            .method(Method::PUT)
            .uri(format!("/items/{}", id))
            .body(Body::from(body))
            .unwrap();

        let res = build_test_app()
            .await
            .oneshot(req)
            .await
            .unwrap();

        assert_eq!(res.status(), StatusCode::OK);

        let body = hyper::body::to_bytes(res.into_body()).await.unwrap();
        let json = std::str::from_utf8(&body).unwrap();
        assert_json!(
            json,
            {
                "data": {
                    "id": id,
                    "title": new_title,
                    "url": new_url,
                    "is_archived": false,
                }
            }
        );
    }).await;
}

#[tokio::test]
async fn test_item_archive() {
    run_test(async {
        let title = "example";
        let url = "https://example.com/";
        let id = repo::Item::insert(&*POOL, title, url).await.unwrap();

        let req = authed_request()
            .method(Method::PATCH)
            .uri(format!("/items/{}/archive", id))
            .body(Body::empty())
            .unwrap();

        let res = build_test_app()
            .await
            .oneshot(req)
            .await
            .unwrap();

        assert_eq!(res.status(), StatusCode::OK);

        let body = hyper::body::to_bytes(res.into_body()).await.unwrap();
        let json = std::str::from_utf8(&body).unwrap();
        assert_json!(
            json,
            {
                "data": {
                    "title": title,
                    "url": url,
                    "is_archived": true
                }
            }
        );
    }).await;
}

#[tokio::test]
async fn test_item_unarchive() {
    run_test(async {
        let title = "example";
        let url = "https://example.com/";
        let id = repo::Item::insert(&*POOL, title, url).await.unwrap();
        repo::Item::update_is_archived(&*POOL, id, true).await.unwrap();

        let req = authed_request()
            .method(Method::PATCH)
            .uri(format!("/items/{}/unarchive", id))
            .body(Body::empty())
            .unwrap();

        let res = build_test_app()
            .await
            .oneshot(req)
            .await
            .unwrap();

        assert_eq!(res.status(), StatusCode::OK);

        let body = hyper::body::to_bytes(res.into_body()).await.unwrap();
        let json = std::str::from_utf8(&body).unwrap();
        assert_json!(
            json,
            {
                "data": {
                    "title": title,
                    "url": url,
                    "is_archived": false,
                }
            }
        );
    }).await;
}

#[tokio::test]
async fn test_item_delete() {
    run_test(async {
        let id = repo::Item::insert(&*POOL, "example", "https://example.com/").await.unwrap();

        let req = authed_request()
            .method(Method::DELETE)
            .uri(format!("/items/{}", id))
            .body(Body::empty())
            .unwrap();

        let res = build_test_app()
            .await
            .oneshot(req)
            .await
            .unwrap();

        assert_eq!(res.status(), StatusCode::OK);

        let option = repo::Item::find(&*POOL, id).await.unwrap();
        assert_eq!(option, None);
    }).await;
}
