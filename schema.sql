CREATE TABLE items (
    id          UUID PRIMARY KEY,
    title       TEXT NOT NULL,
    url         TEXT NOT NULL,
    thumbnail   TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);

CREATE TABLE tags (
    id          UUID PRIMARY KEY,
    name        TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX tags_name ON tags(name);

CREATE TABLE item_tags (
    id          UUID PRIMARY KEY,
    item_id     UUID NOT NULL REFERENCES items,
    tag_id      UUID NOT NULL REFERENCES tags,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX item_tags_item_id_tag_id ON item_tags(item_id, tag_id);
