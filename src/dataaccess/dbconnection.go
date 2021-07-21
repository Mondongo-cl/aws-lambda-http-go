package dataaccess

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Mondongo-cl/http-rest-echo-go/common"
)

func (c *MySQLConnection) CreateDelayedHostTable() error {
	cnn, err := c.open()
	if err != nil {
		return err
	}
	defer cnn.Close()
	_, err = cnn.Exec(`CREATE TABLE DelayedHost(
		ID int NOT NULL AUTO_INCREMENT,
		HostName varchar(4000) NOT NULL,
		CreationDate datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (ID)
	  )`)
	if err != nil {
		return err
	}
	_, err = cnn.Exec("insert into DelayedHost(hostname) values(?)", common.GetHostName())
	if err != nil {
		log.Printf("[%s]:: can't set this host as delayed in while creating delayed host table error: %s", common.GetHostName(), err.Error())
	}
	return nil
}

func (c *MySQLConnection) CkeckIfTableExists(tablename string) bool {
	cnn, err := c.open()
	if err != nil {
		return false
	}
	defer cnn.Close()
	row, err := cnn.Query("SHOW TABLES LIKE ?", tablename)
	if err != nil {
		return false
	}
	var t *string
	if row.Next() {
		err = row.Scan(&t)
		if err != nil || t == nil {
			return false
		}
	}

	if *t == "" {
		return false
	} else {
		return true
	}
}

func (c *MySQLConnection) CheckVesion() (*string, error) {
	cnn, err := c.open()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer cnn.Close()
	versionRow, err := cnn.Query("select version() as Version")
	if err != nil {
		return nil, err
	}
	var currentVersion *string
	if versionRow.Next() {
		err = versionRow.Scan(&currentVersion)
		if err != nil {
			*currentVersion = "unknown"
		}
	}
	return currentVersion, nil
}

func (c *MySQLConnection) Select(query string, params ...interface{}) (*sql.Rows, error) {
	cnn, err := c.open()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer cnn.Close()
	if params != nil {
		return cnn.Query(query, params...)
	} else {
		return cnn.Query(query)
	}
}

func (c *MySQLConnection) SelectOne(query string, params ...interface{}) (*sql.Row, error) {
	cnn, err := c.open()

	if err != nil {
		return nil, err
	}

	defer cnn.Close()
	if params != nil {
		return cnn.QueryRow(query, params...), nil
	} else {
		return cnn.QueryRow(query), nil
	}
}

func (c *MySQLConnection) Execute(query string, params ...interface{}) (sql.Result, error) {
	cnn, err := c.open()

	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer cnn.Close()
	if params != nil {
		return cnn.Exec(query, params...)
	} else {
		return cnn.Exec(query)
	}
}
