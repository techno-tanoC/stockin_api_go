mod item;

pub use item::*;

pub trait Exe<'c>: sqlx::Executor<'c, Database = sqlx::MySql> {}
impl<'c, T: sqlx::Executor<'c, Database = sqlx::MySql>> Exe<'c> for T {}
