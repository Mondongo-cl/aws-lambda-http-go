package dataaccess

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/Mondongo-cl/http-rest-echo-go/datatypes"
)

var (
	mySQLConnection *MySQLConnection = &MySQLConnection{}
)

func Configure(username *string, password *string, hostname *string, port *int, databasename *string) {
	const cnnStr string = "%s:%s@tcp(%s:%d)/%s"
	mySQLConnection.CnnStr = fmt.Sprintf(cnnStr, *username, *password, *hostname, *port, *databasename)
}

func GetAll() ([]MessageRow, error) {

	dbdata, err := mySQLConnection.Select("SELECT ID, Message FROM Messages")

	if err != nil {
		log.Panic("selects statement failed " + err.Error())
	}

	var slice []MessageRow

	for dbdata.Next() {
		item := &MessageRow{}
		dbdata.Scan(&item.Id, &item.Message)
		slice = append(slice, *item)
	}
	return slice, nil
}

func Add(message string) (int64, error) {
	result, err := mySQLConnection.Execute("INSERT INTO Messages (Message) VALUES(?);", message)
	if err != nil {
		log.Panic(err.Error())
	}
	if result != nil {
		if affected, err := result.RowsAffected(); err == nil {
			log.Println(affected, " rows affected")
			return affected, nil
		} else {
			return int64(0), err
		}
	}
	return 0, nil
}

func Remove(id int32) (int64, error) {
	result, err := mySQLConnection.Execute("DELETE FROM Messages WHERE ID = ?;", id)
	if err != nil {
		log.Panic(err.Error())
	}
	if result != nil {
		if affected, err := result.RowsAffected(); err == nil {
			log.Println(affected, " rows affected")
			return affected, nil
		} else {
			return int64(0), err
		}
	}
	return 0, nil
}

func Find(id int32) (*string, error) {
	dbdata, err := mySQLConnection.SelectOne("SELECT ID, Message FROM Messages WHERE ID = ?", id)
	if err != nil {
		log.Panic(err.Error())
	}
	if dbdata != nil {
		var messageValue string
		var id int
		if err := dbdata.Scan(&id, &messageValue); err != nil {
			log.Println("ID ", id, " not found details:: ", err.Error())
			return nil, errors.New(err.Error())
		}
		return &messageValue, nil
	}
	return nil, errors.New("no data found")
}
