-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table friend (
    "from" Integer REFERENCES "user"(id),
    "to" Integer REFERENCES "user"(id),
    PRIMARY KEY("from", "to")
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table friend
