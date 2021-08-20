use anyhow::Result;
use sqlx::sqlite::SqlitePool;
use std::env;

use stockin_api::repo;

#[tokio::main]
async fn main() -> Result<()> {
    let database_url = env::var("DATABASE_URL")?;
    let pool = SqlitePool::connect(&database_url).await?;

    for item in items() {
        repo::Item::insert(&pool, item.0, item.1).await?;
    }

    Ok(())
}

fn items() -> Vec<(&'static str, &'static str)> {
    vec![
        (
            "Android",
            "https://www.android.com/"
        ),
        (
            "TypeScript",
            "https://www.typescriptlang.org/"
        ),
        (
            "The Go Programming Language",
            "https://golang.org/"
        ),
        (
            "Docker",
            "https://www.docker.com/"
        ),
        (
            "Haskell Language",
            "https://www.haskell.org/"
        ),
        (
            "Rustプログラミング言語",
            "https://www.rust-lang.org/ja"
        ),
        (
            "Qiita",
            "https://qiita.com/"
        ),
        (
            "Zenn",
            "https://zenn.dev/"
        ),
        (
            "GitLab",
            "https://gitlab.com/"
        ),
        (
            "GitHub",
            "https://github.com/"
        ),
        (
            "Twitter",
            "https://twitter.com/home"
        ),
        (
            "GMail",
            "https://mail.google.com/"
        ),
        (
            "Google Drive",
            "https://drive.google.com/"
        ),
        (
            "ニコニコ動画",
            "https://www.nicovideo.jp/"
        ),
        (
            "YouTube",
            "https://www.youtube.com/"
        ),
    ]
}
