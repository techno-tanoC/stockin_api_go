mod item;

pub use item::*;

pub trait Exe<'c>: sqlx::Executor<'c, Database = sqlx::Sqlite> {}
impl<'c, T: sqlx::Executor<'c, Database = sqlx::Sqlite>> Exe<'c> for T {}
