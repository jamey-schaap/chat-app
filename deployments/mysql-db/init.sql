CREATE DATABASE IF NOT EXISTS main;
USE main;

CREATE TABLE users (
    id varchar(36),
    firstName varchar(100),
    lastName varchar(100),
    PRIMARY KEY (id)
);
    
CREATE TABLE chat_messages (
    id varchar(36) NOT NULL,
    message varchar(1020),
    userId varchar(36) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (userId) REFERENCES users(id)
);
