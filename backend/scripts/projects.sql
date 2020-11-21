INSERT INTO permissions (id, read, write, admin)
VALUES (1, true, true, true);
INSERT INTO permissions (id, read, write, admin)
VALUES (2, true, true, false);

INSERT INTO datetimes (id, created, updated, accessed)
VALUES (1, 1605925262, 1605925262, 1605925262);

INSERT INTO projects (id, owner_id, default_permissions_id, datetimes_id, title, description)
VALUES (1, 1, 2, 1, 'First project', 'This is the first project');

INSERT INTO project_users (id, user_id, project_id, permissions_id)
VALUES (1, 1, 1, 1);