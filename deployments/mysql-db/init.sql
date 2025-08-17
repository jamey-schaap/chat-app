CREATE
DATABASE IF NOT EXISTS main;
USE
main;

CREATE TABLE users
(
    id         binary(16) NOT NULL,
    first_name varchar(100),
    last_name  varchar(100),
    created_at dateTime NOT NULL,
    updated_at dateTime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE chat_messages
(
    id         binary(16) NOT NULL,
    message    varchar(1020),
    user_id    binary(16) NOT NULL,
    created_at dateTime NOT NULL,
    updated_at dateTime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO users (`id`, first_name, last_name, `created_at`)
VALUES (UUID_TO_BIN('aa48082a-5d5a-4147-9de3-2d994b6f790d'), 'John', 'Doe', (SELECT NOW()));

INSERT INTO chat_messages (`id`, `message`, `user_id`, `created_at`)
values (UUID_TO_BIN((SELECT UUID())), "Hi", UUID_TO_BIN('aa48082a-5d5a-4147-9de3-2d994b6f790d'), (SELECT NOW())),
       (UUID_TO_BIN((SELECT UUID())), "How are you?", UUID_TO_BIN('aa48082a-5d5a-4147-9de3-2d994b6f790d'), (SELECT NOW())),
       (UUID_TO_BIN((SELECT UUID())), "Great to hear!", UUID_TO_BIN('aa48082a-5d5a-4147-9de3-2d994b6f790d'), (SELECT NOW()));