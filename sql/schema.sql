CREATE DATABASE sybo;

USE sybo;

CREATE TABLE users (
    uuid VARCHAR(36),
    name VARCHAR(64),
    PRIMARY KEY (uuid)
);

CREATE TABLE state (
    uuid VARCHAR(36),
    games_played INT,
    score INT,
    PRIMARY KEY (uuid)
);

CREATE TABLE friends (
    uuid VARCHAR(36),
    friends JSON 
);
