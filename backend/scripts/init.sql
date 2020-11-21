CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    nickname varchar(32) UNIQUE NOT NULL,
    email varchar(32) NOT NULL,
    password varchar(32) NOT NULL,
    avatar varchar(100) NOT NULL DEFAULT ''
);
CREATE TABLE IF NOT EXISTS permissions (
    id serial PRIMARY KEY,
    read boolean NOT NULL DEFAULT false,
    write boolean NOT NULL DEFAULT false,
    admin boolean NOT NULL DEFAULT false
);
CREATE TABLE IF NOT EXISTS datetimes (
    id serial PRIMARY KEY,
    created bigint NOT NULL,
    updated bigint NOT NULL,
    accessed bigint NOT NULL
);
CREATE TABLE IF NOT EXISTS projects (
    id serial PRIMARY KEY,
    owner_id int REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    default_permissions_id int REFERENCES permissions (id) ON DELETE CASCADE NOT NULL,
    datetimes_id int REFERENCES datetimes (id) ON DELETE CASCADE NOT NULL,
    title varchar(50) NOT NULL,
    description text
);
CREATE TABLE IF NOT EXISTS project_users (
    id serial PRIMARY KEY,
    user_id int REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    project_id int REFERENCES projects (id) ON DELETE CASCADE NOT NULL,
    permissions_id int REFERENCES permissions (id) ON DELETE CASCADE NOT NULL
);
CREATE TABLE IF NOT EXISTS boards (
    id serial PRIMARY KEY,
    project_id int REFERENCES projects (id) ON DELETE CASCADE NOT NULL,
    default_permissions_id int REFERENCES permissions (id) ON DELETE CASCADE NOT NULL,
    datetimes_id int references datetimes (id) ON DELETE CASCADE NOT NULL,
    title varchar(30) NOT NULL
);
CREATE TABLE IF NOT EXISTS board_users (
    id serial NOT NULL unique,
    user_id int REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    board_id int REFERENCES boards (id) ON DELETE CASCADE NOT NULL,
    permissions_id int REFERENCES permissions (id) ON DELETE CASCADE NOT NULL
);
-- CREATE TABLE labels
-- (
--     id serial      not null unique,
-- 	name varchar(30) not null
-- );
-- CREATE TABLE task_lists
-- (
--     id serial      not null unique,
--     title varchar(30) not null,
--     board_id int references boards (id) on delete cascade      not null,
--     datetimes int references datetimes (id) on delete cascade      not null,
-- 	position smallint not null
-- );
-- CREATE TABLE tasks
-- (
-- 	id serial      not null unique,
--     list_id int references task_lists (id) on delete cascade      not null,
--     title varchar(30) not null,
-- 	datetimes int references datetimes (id) on delete cascade      not null,
--     position smallint not null
-- );
-- drop table tasks;
-- drop table task_lists;
-- drop table labels;
-- drop table board_users;
-- drop table boards;
-- drop table project_users;
-- drop table projects;
-- drop table datetimes;
-- drop table permissions;
-- drop table users;