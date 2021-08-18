mod item;

pub use item::*;

pub trait Conn<'c>: sqlx::Executor<'c, Database = sqlx::Sqlite> {}
impl<'c, T: sqlx::Executor<'c, Database = sqlx::Sqlite>> Conn<'c> for T {}
