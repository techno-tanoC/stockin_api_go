CREATE TABLE items (
    id         BIGINT PRIMARY KEY AUTO_INCREMENT,
    title      TEXT NOT NULL,
    url        TEXT NOT NULL,
    thumbnail  TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
