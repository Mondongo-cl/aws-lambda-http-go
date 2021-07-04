package dataaccess

import (
	"database/sql"
	"errors"
	"log"
)

type MySQLConnection struct {
	CnnStr string
}

func (c *MySQLConnection) open() (*sql.DB, error) {
	log.Println("init connection check")
	if c == nil {
		log.Fatalln("connection object is nil")
		return nil, errors.New("invalid connection object")
	}
	log.Println("connection open to ", c.CnnStr, " using mysql as driver")
	cnn, err := sql.Open("mysql", c.CnnStr)
	if err != nil {
		log.Fatal(err)
		log.Fatal("error while open the connection ", err.Error())
		return nil, errors.New(err.Error())
	}
	log.Println("connection open successfully")
	return cnn, nil
}
