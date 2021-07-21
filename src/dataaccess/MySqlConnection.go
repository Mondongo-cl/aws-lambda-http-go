package dataaccess

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type MySQLConnection struct {
	CnnStr string
}

func (c *MySQLConnection) open() (*sql.DB, error) {
	sw := time.Now()
	log.Printf("[%s]::init connection check", getHostName())
	if c == nil {
		log.Fatalf("[%s]::connection object is nil", getHostName())
		return nil, errors.New("invalid connection object")
	}
	defer log.Printf("[%s]::connection open successfully in %d (ms)", getHostName(), time.Since(sw).Milliseconds())
	log.Printf("[%s]::connection open using mysql as driver", getHostName())
	cnn, err := sql.Open("mysql", c.CnnStr)
	if err != nil {
		log.Fatalf("[%s]::error while open the connection  --- %s", getHostName(), err.Error())
		return nil, errors.New(err.Error())
	}

	return cnn, nil
}
