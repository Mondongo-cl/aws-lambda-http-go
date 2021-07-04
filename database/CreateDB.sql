CREATE DATABASE testdb

use testdb

CREATE TABLE Messages (
  ID int NOT NULL AUTO_INCREMENT,
  Message varchar(4000) NOT NULL,
  PRIMARY KEY (ID)
);

INSERT INTO Messages (Message) VALUES ('Database Created');
