CREATE TYPE Role_Type AS ENUM('Novice', 'Apprentice','Adept','Expert','Master','Loli','Anahoret');

CREATE TABLE IF NOT EXISTS roles
(
    role_id SERIAL PRIMARY KEY,
    role_type Role_Type
);

CREATE TABLE IF NOT EXISTS users
(
    user_id VARCHAR(40) PRIMARY KEY,
    username VARCHAR(200) NOT NULL,
    user_level INT DEFAULT 0,
    role_id INT NOT NULL DEFAULT 0,
    FOREIGN KEY role_id
        REFERENCES roles (role_id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages
(
    user_id VARCHAR(40) NOT NULL,
    message_count INT DEFAULT 0,
    FOREIGN KEY user_id
        REFERENCES users (user_id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tags
(
    tag_id SERIAL PRIMARY KEY,
    tag_name VARCHAR(100) NOT NULL,
    tag_body VARCHAR(2000) NOT NULL
);