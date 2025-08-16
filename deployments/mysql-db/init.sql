CREATE
DATABASE IF NOT EXISTS main;
USE
main;

CREATE TABLE users
(
    id        binary(16) NOT NULL,
    firstName varchar(100),
    lastName  varchar(100),
    PRIMARY KEY (id)
);

CREATE TABLE chat_messages
(
    id      binary(16) NOT NULL,
    message varchar(1020),
    userId  binary(16) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (userId) REFERENCES users (id)
);

INSERT INTO users
VALUES (UUID_TO_BIN('aa48082a-5d5a-4147-9de3-2d994b6f790d'), 'John', 'Doe');