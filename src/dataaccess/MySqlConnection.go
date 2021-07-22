package dataaccess

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const cnnStr string = "%s:%s@tcp(%s:%d)/%s"

type MySQLConnection struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (c *MySQLConnection) Configure(Hostname *string, Port *int, Username *string, Password *string, Database *string) error {
	log.Printf("Starting MySql Connection Configuration with %s:****@%s:%d/%s", *Username, *Hostname, *Port, *Database)
	c.Host = *Hostname
	c.Port = *Port
	c.Username = *Username
	c.Password = *Password
	c.Database = *Database
	return nil
}

func (c *MySQLConnection) Open() (*sql.DB, error) {
	log.Println("init connection check")
	if c == nil {
		log.Fatalln("connection object is nil")
		return nil, errors.New("invalid connection object")
	}
	log.Printf("connection open to %s:%d using mysql as driver", c.Host, c.Port)
	cnn, err := sql.Open("mysql", fmt.Sprintf(cnnStr, c.Username, c.Password, c.Host, c.Port, c.Database))
	if err != nil {
		log.Fatal(err)
		log.Fatal("error while open the connection ", err.Error())
		return nil, errors.New(err.Error())
	}
	cnn.SetConnMaxIdleTime(time.Second * 1)
	cnn.SetConnMaxLifetime(time.Second * 3)
	cnn.SetMaxIdleConns(10)
	cnn.SetMaxOpenConns(10)
	return cnn, cnn.Ping()
}
