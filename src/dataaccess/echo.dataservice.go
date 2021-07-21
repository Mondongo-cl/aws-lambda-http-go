package dataaccess

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/common"
	_ "github.com/Mondongo-cl/http-rest-echo-go/datatypes"
)

var (
	mySQLConnection *MySQLConnection = &MySQLConnection{}
)

func GetMySqlVersion() *string {
	version, err := mySQLConnection.CheckVesion()
	if err != nil {
		return nil
	}
	return version
}

func Configure(username *string, password *string, hostname *string, port *int, databasename *string) {
	const cnnStr string = "%s:%s@tcp(%s:%d)/%s"
	mySQLConnection.CnnStr = fmt.Sprintf(cnnStr, *username, *password, *hostname, *port, *databasename)
	version, err := mySQLConnection.CheckVesion()
	if err != nil {
		log.Fatalf("[%s]::can't open connection to mysql instance, error:%s", common.GetHostName(), err.Error())
	}
	log.Printf("Connection Successfull to %s version", *version)
	if !mySQLConnection.CkeckIfTableExists("DelayedHost") {
		err = mySQLConnection.CreateDelayedHostTable()
		if err != nil {
			log.Printf("[%s]::error while create host delay behavior table support error: %s", common.GetHostName(), err.Error())
		}

	}
}

func IsDelayedHost() bool {
	row, err := mySQLConnection.Select("select HostName from DelayedHost order by id desc limit 1;")
	if err != nil {
		log.Printf("[%s]::can't find the current delayed hostname, error: %s", common.GetHostName(), err.Error())
	}
	var hostname string = ""
	if row.Next() {
		row.Scan(&hostname)
	}
	if hostname == common.GetHostName() {
		return true
	}
	return false

}

func GetAll() ([]MessageRow, error) {
	sw := time.Now()
	dbdata, err := mySQLConnection.Select("SELECT ID, Message FROM Messages")

	if err != nil {
		log.Panicf("[%s]::select statement failed -- %s", common.GetHostName(), err.Error())
	}
	defer log.Printf("[%s]::command select executed in %d (ms)", common.GetHostName(), time.Since(sw).Milliseconds())
	var slice []MessageRow

	for dbdata.Next() {
		item := &MessageRow{}
		dbdata.Scan(&item.Id, &item.Message)
		slice = append(slice, *item)
	}
	return slice, nil
}

func Add(message string) (int64, error) {
	sw := time.Now()
	result, err := mySQLConnection.Execute("INSERT INTO Messages (Message) VALUES(?);", message)
	if err != nil {
		log.Panicf("[%s]::insert statement failed -- %s", common.GetHostName(), err.Error())
	}
	defer log.Printf("[%s]::command insert executed in %d (ms)", common.GetHostName(), time.Since(sw).Milliseconds())
	if result != nil {
		if affected, err := result.RowsAffected(); err == nil {
			log.Printf("[%s]::%d rows affected", common.GetHostName(), affected)
			return affected, nil
		} else {
			log.Printf("[%s]::no data found", common.GetHostName())
			return int64(0), err
		}
	}
	return 0, nil
}

func Remove(id int32) (int64, error) {
	sw := time.Now()
	result, err := mySQLConnection.Execute("DELETE FROM Messages WHERE ID = ?;", id)
	if err != nil {
		log.Panicf("[%s]::delete statement failed -- %s", common.GetHostName(), err.Error())
	}
	defer log.Printf("[%s]::command delete executed in %d (ms)", common.GetHostName(), time.Since(sw).Milliseconds())
	if result != nil {
		if affected, err := result.RowsAffected(); err == nil {
			log.Printf("[%s]::%d rows affected", common.GetHostName(), affected)
			return affected, nil
		} else {
			log.Printf("[%s]::no data found", common.GetHostName())
			return int64(0), err
		}
	}
	return 0, nil
}

func Find(id int32) (*string, error) {
	sw := time.Now()
	dbdata, err := mySQLConnection.SelectOne("SELECT ID, Message FROM Messages WHERE ID = ?", id)
	if err != nil {
		log.Panic(err.Error())
	}
	defer log.Printf("[%s]::command select by id executed in %d (ms)", common.GetHostName(), time.Since(sw).Milliseconds())
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
