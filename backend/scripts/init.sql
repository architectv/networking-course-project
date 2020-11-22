drop table tasks;
-- drop table labels;
DROP TABLE task_lists;
DROP TABLE board_users;
DROP TABLE boards;
DROP TABLE project_users;
DROP TABLE projects;
DROP TABLE datetimes;
DROP TABLE permissions;
DROP TABLE users;
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
    owner_id int REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    default_permissions_id int REFERENCES permissions (id) ON DELETE CASCADE NOT NULL,
    datetimes_id int references datetimes (id) ON DELETE CASCADE NOT NULL,
    title varchar(30) NOT NULL
);
CREATE TABLE IF NOT EXISTS board_users (
    id serial PRIMARY KEY,
    user_id int REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    board_id int REFERENCES boards (id) ON DELETE CASCADE NOT NULL,
    permissions_id int REFERENCES permissions (id) ON DELETE CASCADE NOT NULL
);
-- CREATE TABLE labels
-- (
--     id serial      not null unique,
-- 	name varchar(30) not null
-- );
CREATE TABLE IF NOT EXISTS task_lists (
    id serial PRIMARY KEY,
    board_id int REFERENCES boards (id) ON DELETE CASCADE NOT NULL,
    title varchar(30) NOT NULL,
    position int NOT NULL
);
CREATE TABLE tasks
(
	id serial      not null unique,
    list_id int references task_lists (id) on delete cascade      not null,
    title varchar(30) not null,
	datetimes_id int references datetimes (id) on delete cascade      not null,
    position smallint not null
);
-- USERS
-- 1
INSERT INTO users (nickname, email, avatar, password)
VALUES ('alex', 'alex@mail.ru', '', 'qwerty');
-- 2
INSERT INTO users (nickname, email, avatar, password)
VALUES ('test_user', 'test_user@mail.ru', '', 'qwerty');
-- 3
INSERT INTO users (nickname, email, avatar, password)
VALUES ('nick1', 'nick1@mail.ru', '', 'qwerty');
-- PROJECTS
-- 1
INSERT INTO permissions (read, write, admin)
VALUES (true, true, true);
-- 2
INSERT INTO permissions (read, write, admin)
VALUES (true, true, false);
-- 1
INSERT INTO datetimes (created, updated, accessed)
VALUES (1605925262, 1605925262, 1605925262);
-- 1
INSERT INTO projects (
        owner_id,
        default_permissions_id,
        datetimes_id,
        title,
        description
    )
VALUES (
        1,
        2,
        1,
        'First project',
        'This is the first project'
    );
-- 1
INSERT INTO project_users (user_id, project_id, permissions_id)
VALUES (1, 1, 1);
-- BOARDS
-- 3
INSERT INTO permissions (read, write, admin)
VALUES (true, true, true);
-- 4
INSERT INTO permissions (read, write, admin)
VALUES (true, true, false);
-- 2
INSERT INTO datetimes (created, updated, accessed)
VALUES (1605925262, 1605925262, 1605925262);
-- 1
INSERT INTO boards (
        project_id,
        owner_id,
        default_permissions_id,
        datetimes_id,
        title
    )
VALUES (1, 1, 4, 2, 'First board');
-- 1
INSERT INTO board_users (user_id, board_id, permissions_id)
VALUES (1, 1, 3);
--
-- 5
-- INSERT INTO permissions (read, write, admin)
-- VALUES (true, true, true);
-- -- 3
-- INSERT INTO datetimes (created, updated, accessed)
-- VALUES (1605925262, 1605925262, 1605925262);
-- -- 2
-- INSERT INTO project_users (user_id, project_id, permissions_id)
-- VALUES (2, 1, 5);
-- LISTS
-- 1
INSERT INTO task_lists (board_id, title, position)
VALUES (1, 'Not Stated', 0);
-- 2
INSERT INTO task_lists (board_id, title, position)
VALUES (1, 'SSSSSSSSS', 1);
-- 3
INSERT INTO task_lists (board_id, title, position)
VALUES (1, 'herbfneifj', 2);

INSERT INTO task_lists (board_id, title, position)
VALUES (1, 'Not Stated', 0);
-- TASKS
-- 3
INSERT INTO datetimes (created, updated, accessed)
VALUES (1605925262, 1605925262, 1605925262);
-- 4
INSERT INTO datetimes (created, updated, accessed)
VALUES (1605925262, 1605925262, 1605925262);
-- 5
INSERT INTO datetimes (created, updated, accessed)
VALUES (1605925262, 1605925262, 1605925262);
-- 1
INSERT INTO tasks (list_id, title, datetimes_id, position)
VALUES (1, 'First task', 3, 1);
-- 2
INSERT INTO tasks (list_id, title, datetimes_id, position)
VALUES (1, 'Second task', 4, 2);
-- 3
INSERT INTO tasks (list_id, title, datetimes_id, position)
VALUES (1, 'Third task', 5, 3);