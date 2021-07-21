package dataaccess

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/common"
)

type MySQLConnection struct {
	CnnStr string
}

func (c *MySQLConnection) open() (*sql.DB, error) {
	sw := time.Now()
	log.Printf("[%s]::init connection check", common.GetHostName())
	if c == nil {
		log.Fatalf("[%s]::connection object is nil", common.GetHostName())
		return nil, errors.New("invalid connection object")
	}
	defer log.Printf("[%s]::connection open successfully in %d (ms)", common.GetHostName(), time.Since(sw).Milliseconds())
	log.Printf("[%s]::connection open using mysql as driver", common.GetHostName())
	cnn, err := sql.Open("mysql", c.CnnStr)
	if err != nil {
		log.Fatalf("[%s]::error while open the connection  --- %s", common.GetHostName(), err.Error())
		return nil, errors.New(err.Error())
	}

	return cnn, nil
}
