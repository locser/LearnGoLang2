-- +migrate Up
CREATE table "users"
(
    "user_id"    text primary key,
    "fuul_name"  text,
    "email"      text unique,
    "password"   text,
    "role"       text,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" timestamp not NULL

);

CREATE table "repos"
(
    "name"        text primary key,
    "description" text,
    "url"         text,
    "color"       text,
    "lang"        text,
    "fork"        text,
    "stars"       text,
    "stars_today" text,
    "build_by"    text,
    "created_at"  TIMESTAMP not null,
    "updated_at"  Timestamp not null
);

create table "bookmarks"
(
    "bid"        text primary key,
    "user_id"    text,
    "repo_name"  text,
    "created_at" TIMESTAMP not null,
    "updated_at" timestamp not null
);

ALTER table "bookmarks"
    add foreign key ("user_id") references "users" ("user_id");
ALTER table "bookmarks"
    add foreign key ("repo_name") references "repos" ("name");

-- +migrate Down
drop table bookmarks;
drop table users;
drop table reops;
