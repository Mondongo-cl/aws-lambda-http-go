package dataaccess

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/common"
)

func (c *MySQLConnection) Open() (*sql.DB, error) {
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
	cnn.SetConnMaxIdleTime(time.Second * 1)
	cnn.SetMaxOpenConns(10)
	cnn.SetMaxIdleConns(10)
	return cnn, nil
}

func (c *MySQLConnection) CreateDelayedHostTable(cnn *sql.DB) error {

	_, err := cnn.Exec(`CREATE TABLE DelayedHost(
		ID int NOT NULL AUTO_INCREMENT,
		HostName varchar(4000) NOT NULL,
		CreationDate datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (ID)
	  )`)
	if err != nil {
		cnn.Close()
		return err
	}
	_, err = cnn.Exec("insert into DelayedHost(hostname) values(?)", common.GetHostName())
	cnn.Close()
	if err != nil {
		log.Printf("[%s]:: can't set this host as delayed in while creating delayed host table error: %s", common.GetHostName(), err.Error())
	}

	return nil
}

func (c *MySQLConnection) CkeckIfTableExists(tablename string) bool {
	cnn, err := c.Open()
	if err != nil {
		return false
	}

	row, err := cnn.Query("SHOW TABLES LIKE ?", tablename)
	if err != nil {
		return false
	}
	var t *string
	if row.Next() {
		err = row.Scan(&t)
		if err != nil || t == nil {
			*t = ""
		}
	}
	cnn.Close()
	if *t == "" {
		return false
	} else {
		return true
	}
}

func (c *MySQLConnection) CheckVesion() (*string, error) {
	cnn, err := c.Open()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	versionRow, err := cnn.Query("select version() as Version")
	if err != nil {
		cnn.Close()
		return nil, err
	}
	var currentVersion *string
	if versionRow.Next() {
		err = versionRow.Scan(&currentVersion)
		if err != nil {
			*currentVersion = "unknown"
		}
	}
	cnn.Close()
	return currentVersion, nil
}
