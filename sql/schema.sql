
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS user_session;
DROP TABLE IF EXISTS users;

CREATE TABLE users ()
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(100) NOT NULL


-- Create User Session Table (to store sensitive data separately)
CREATE TABLE IF NOT EXISTS user_session (
  id           INT AUTO_INCREMENT PRIMARY KEY,
  username     VARCHAR(128) NOT NULL UNIQUE,
  password     VARCHAR(256) NOT NULL,
  email        VARCHAR(128) NOT NULL UNIQUE,
  FOREIGN KEY (email) REFERENCES users(email) ON DELETE CASCADE ON UPDATE CASCADE
);1


-- Create Template Table
CREATE TABLE IF NOT EXISTS template (
  id            INT AUTO_INCREMENT PRIMARY KEY,
  path          VARCHAR(128) NOT NULL
);

-- Insert Roles
INSERT IGNORE INTO roles (role_name)
VALUES
    ('admin'),
    ('user');


-- Insert User Sessions with Hashed Passwords
INSERT IGNORE INTO user_session (username, password, email)
VALUES
    ('Cesar',   '$2y$12$V0yU7tkK.G/KzhRFkFA9uOnlB4EFg6ScMxX1tynnzzBlrUABtEhuC', 'cesar@gmail.it'),
    ('Damaris', '$2y$12$V0yU7tkK.G/KzhRFkFA9uOnlB4EFg6ScMxX1tynnzzBlrUABtEhuC', 'damaris@yahoo.it'),
    ('Tabita',  '$2y$12$R1lzLz2/8Z/jjvR1C6Aeh.hoqgJXyq5T.eW4Tpr.q1ACf9FzHmbOC', 'tabita@yahoo.it'),
    ('Cristy',  '$2a$10$HyG1Q1L5nyvl1H/yrpxlC.gB8ZLPGO/6sW2g9pelYUz2.IY.eBvUK', 'cristy@yahoo.it');

-- Assign Roles to Users
INSERT IGNORE INTO user_roles (user_id, role_id)
VALUES
    ((SELECT id FROM users WHERE email = 'cesar@gmail.it'), (SELECT id FROM roles WHERE role_name = 'user')),
    ((SELECT id FROM users WHERE email = 'damaris@yahoo.it'), (SELECT id FROM roles WHERE role_name = 'user')),
    ((SELECT id FROM users WHERE email = 'tabita@yahoo.it'), (SELECT id FROM roles WHERE role_name = 'user')),
    ((SELECT id FROM users WHERE email = 'cristy@yahoo.it'), (SELECT id FROM roles WHERE role_name = 'admin'));
