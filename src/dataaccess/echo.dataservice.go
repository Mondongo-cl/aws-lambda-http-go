package dataaccess

import (
	"errors"
	"log"

	"github.com/Mondongo-cl/http-rest-echo-go/settings"
)

var mySQLConnection MySQLConnection = MySQLConnection{}

func Configure(settings settings.ConnectionSettings) {
	mySQLConnection.Configure(&settings.Host, &settings.Port, &settings.Username, &settings.Password, &settings.Database)
	log.Printf("starting connection test to %s mysql server", settings.Host)
	c, e := mySQLConnection.Open()
	if e == nil {
		log.Printf("connection successfully to %s, starting ping", settings.Host)
		e = c.Ping()
		if e != nil {
			log.Printf("ERROR:: %v while ping to %s\\%s", e, settings.Host, settings.Database)
			panic(e)
		}
		log.Println("all ok!!, closing connections")
		c.Close()
	} else {
		panic(e)
	}
}

func GetAll() ([]MessageRow, error) {
	log.Println("Get All operation start")
	cnn, err := mySQLConnection.Open()
	if err != nil {
		return nil, err
	}
	dbdata, err := cnn.Query("SELECT ID, Message FROM Messages")
	if err != nil {
		log.Fatal(err)
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
	cnn.Close()
	log.Println("Get all opration end successfully")
	return slice, nil
}

func Add(message string) (int64, error) {
	log.Println("Add operation start")
	cnn, err := mySQLConnection.Open()
	if err != nil {
		return -1, err
	}

	result, err := cnn.Exec("INSERT INTO Messages (Message) VALUES(?);", message)
	cnn.Close()
	if err != nil {
		log.Fatal(err)
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
	cnn, err := mySQLConnection.Open()
	if err != nil {
		return -1, err
	}
	result, err := cnn.Exec("DELETE FROM Messages WHERE ID = ?;", id)
	cnn.Close()
	if err != nil {
		log.Fatal(err)
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
	cnn, err := mySQLConnection.Open()
	if err != nil {
		return nil, err
	}
	dbdata := cnn.QueryRow("SELECT Message FROM Messages WHERE ID = ?", id)
	if dbdata != nil {
		var messageValue string
		var id int
		if err := dbdata.Scan(&messageValue); err != nil {
			log.Println("ID ", id, " not found details:: ", err.Error())
			return nil, err
		}
		return &messageValue, nil
	}
	return nil, errors.New("no data found")
}
