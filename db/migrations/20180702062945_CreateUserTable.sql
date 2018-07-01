-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table "user" (
    id Integer,
    name TEXT NOT NULL,
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table user;
