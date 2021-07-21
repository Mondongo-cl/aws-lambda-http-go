package dataaccess

import (
	"database/sql"
	"errors"
)

func (c *MySQLConnection) CheckVesion() (*string, error) {
	cnn, err := c.open()

	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer cnn.Close()
	versionRow, err := cnn.Query("select Version()")
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
