CREATE TABLE users (
    id serial PRIMARY KEY,
    nickname varchar(32) UNIQUE NOT NULL,
    email varchar(32) NOT NULL,
    password varchar(32) NOT NULL,
    avatar varchar(100) NOT NULL DEFAULT ''
);
-- CREATE TABLE permissions
-- (
--     id serial      not null unique,
-- 	read boolean,
-- 	write boolean,
-- 	admin boolean
-- );
-- CREATE TABLE datetimes 
-- (
--     id serial      not null unique,
-- 	created timestamp, 
-- 	updated timestamp,
-- 	accessed timestamp
-- );
-- CREATE TABLE projects
-- (
--     id serial      not null unique,
--     owner_id int references users (id) on delete cascade      not null,
--     default_permissions int references permissions (id) on delete cascade      not null,
--     datetimes int references datetimes (id) on delete cascade      not null,
--     title varchar(30) not null,
--     description varchar(100)
-- );
-- CREATE TABLE project_users
-- (
--     id serial      not null unique,
--     user_id int references users (id) on delete cascade      not null,
--     project_id int references projects (id) on delete cascade      not null,
--     permissions int references permissions (id) on delete cascade      not null
-- );
-- CREATE TABLE boards
-- (
--     id serial      not null unique,
--     title varchar(30) not null,
--     project_id int references projects (id) on delete cascade      not null,
--     default_permissions int references permissions (id) on delete cascade      not null,
--     datetimes int references datetimes (id) on delete cascade      not null
-- );
-- CREATE TABLE board_users
-- (
--     id serial      not null unique,
--     user_id int references users (id) on delete cascade      not null,
--     board_id int references boards (id) on delete cascade      not null,
--     permissions int references permissions (id) on delete cascade      not null
-- );
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