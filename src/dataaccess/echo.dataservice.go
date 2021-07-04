package dataaccess

import (
	"errors"
	"log"

	_ "github.com/Mondongo-cl/http-rest-echo-go/datatypes"
)

const cnnStr string = "root:123456@tcp(mysql1:3306)/testdb"

var (
	mySQLConnection MySQLConnection = MySQLConnection{CnnStr: cnnStr}
)

func GetAll() ([]MessageRow, error) {
	log.Println("Get All operation start")
	dbdata, err := mySQLConnection.Select("SELECT ID, Message FROM Messages")
	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("selects statement failed")

	}
	var slice []MessageRow
	for dbdata.Next() {
		var messageValue string
		var id int
		dbdata.Scan(&id, &messageValue)
		item := MessageRow{Id: int32(id), Message: messageValue}
		slice = append(slice, item)
	}
	log.Printf("dbData value is %v \n", dbdata)
	return slice, nil
}

func Add(message string) (int64, error) {
	log.Println("Add operation start")

	result, err := mySQLConnection.Execute("INSERT INTO Messages (Message) VALUES(?);", message)
	if err != nil {
		log.Fatal(err.Error())
		return int64(0), err
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
	log.Println("Add operation start")

	result, err := mySQLConnection.Execute("DELETE FROM Messages WHERE ID = ?;", id)
	if err != nil {
		log.Fatal(err.Error())
		return int64(0), err
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
	log.Println("Find operation start")
	dbdata := mySQLConnection.SelectOne("SELECT ID, Message FROM Messages WHERE ID = ?", id)

	if dbdata != nil {
		var messageValue string
		var id int
		if err := dbdata.Scan(&id, &messageValue); err != nil {
			log.Println("ID ", id, " not found details:: ", err.Error())
			return nil, errors.New(err.Error())
		}
		return &messageValue, nil
	}
	log.Printf("dbData value is %v \n", dbdata)
	return nil, errors.New("no data found")
}
