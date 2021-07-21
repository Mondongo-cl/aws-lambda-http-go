CREATE DATABASE testdb

use testdb

CREATE TABLE Messages (
  ID int NOT NULL AUTO_INCREMENT,
  Message varchar(4000) NOT NULL,
  PRIMARY KEY (ID)
);

CREATE TABLE DelayedHost(
  ID int NOT NULL AUTO_INCREMENT,
  HostName varchar(4000) NOT NULL,
  CreationDate datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (ID)
)

INSERT INTO Messages (Message) VALUES ('Database Created');
