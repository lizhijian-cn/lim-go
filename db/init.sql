CREATE DATABASE LIM;
USE LIM;

CREATE TABLE User
(
    id       BIGINT      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(25) NOT NULL,
    password VARCHAR(10) NOT NULL
) AUTO_INCREMENT = 100;

CREATE TABLE Relation
(
    id  BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    aid BIGINT NOT NULL,
    bid BIGINT NOT NULL
) AUTO_INCREMENT = 100;