CREATE DATABASE IF NOT EXISTS sybo;

USE sybo;

CREATE TABLE users (
    uuid VARCHAR(40) DEFAULT (uuid()),
    name VARCHAR(64) NOT NULL,
    PRIMARY KEY (uuid)
);

CREATE TABLE state (
    uuid VARCHAR(40),
    games_played INT NOT NULL DEFAULT (0),
    score INT NOT NULL DEFAULT (0),
    PRIMARY KEY (uuid)
);

CREATE TABLE friends (
    uuid VARCHAR(40),
    friends JSON NOT NULL DEFAULT ('{"friends": []}'),
    PRIMARY KEY (uuid)
);

DELIMITER $$
CREATE TRIGGER after_new_user
AFTER INSERT
ON users FOR EACH ROW
BEGIN
    INSERT INTO state (uuid) VALUES (NEW.uuid);
    INSERT INTO friends (uuid) VALUES (NEW.uuid);
END$$

DELIMITER ;
