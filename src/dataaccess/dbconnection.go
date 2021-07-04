package dataaccess

import (
	"database/sql"
	"errors"
)

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

func (c *MySQLConnection) SelectOne(query string, params ...interface{}) *sql.Row {
	cnn, err := c.open()

	if err != nil {
		return nil
	}

	defer cnn.Close()
	if params != nil {
		return cnn.QueryRow(query, params...)
	} else {
		return cnn.QueryRow(query)
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
