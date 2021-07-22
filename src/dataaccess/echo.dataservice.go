package dataaccess

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/common"
	_ "github.com/Mondongo-cl/http-rest-echo-go/datatypes"
)

var (
	mySQLConnection *MySQLConnection = &MySQLConnection{}
	lastUpdate      time.Time        = time.Now()
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
		cnn, err := mySQLConnection.Open()
		if err != nil {
			log.Fatalf("[%s]::can obtain a connection error:%s", common.GetHostName(), err.Error())
		}
		err = mySQLConnection.CreateDelayedHostTable(cnn)
		cnn.Close()
		if err != nil {
			log.Printf("[%s]::error while create host delay behavior table support error: %s", common.GetHostName(), err.Error())
		}

	}
}

var hostname string = ""

func IsDelayedHost(useDb bool) bool {
	elapsed := time.Since(lastUpdate)
	thereshold := lastUpdate.Add(time.Second * 1)
	var mycnn *sql.DB = nil
	if useDb {
		log.Printf("[%s]::elapsed time %v, threshold:%v next refresh in %v", common.GetHostName(), elapsed, thereshold, (thereshold.Sub(time.Now())))
		if hostname == "" || time.Since(thereshold).Milliseconds() > 0 {
			log.Printf("[%s]::refreshing the delayed host id %s", common.GetHostName(), hostname)
			lastUpdate = time.Now()
			mycnn, err := mySQLConnection.Open()
			if err != nil {
				log.Fatalf("[%s]::error while open the connection error %s", common.GetHostName(), err.Error())
			}
			row, err := mycnn.Query("select HostName from DelayedHost order by id desc limit 1;")
			if err != nil {
				mycnn.Close()
				log.Printf("[%s]::can't find the current delayed hostname, error: %s", common.GetHostName(), err.Error())
			}
			if row != nil {
				if row.Next() {
					row.Scan(&hostname)
				}
			} else {
				hostname = ""
			}
			log.Printf("[%s]::current delayed host id %s, equals:%v", common.GetHostName(), hostname, common.GetHostName() == hostname)
		}
	}
	// leave a open connection to delayed host
	if hostname == common.GetHostName() {
		log.Printf("[%s]::current host is delayed the host is :%s", common.GetHostName(), hostname)
		return true
	}

	if mycnn != nil {
		log.Printf("[%s]::closing connection on host :%s", common.GetHostName(), common.GetHostName())
		mycnn.Close()
	}
	log.Printf("[%s]::current host is not delayed the delayed host is %s", common.GetHostName(), hostname)
	return false

}

func GetAll() ([]MessageRow, error) {
	sw := time.Now()
	sqlcnn, err := mySQLConnection.Open()
	if err != nil {
		log.Panic(err.Error())
	}
	dbdata, err := sqlcnn.Query("SELECT ID, Message FROM Messages")

	if err != nil {
		sqlcnn.Close()
		log.Panicf("[%s]::select statement failed -- %s", common.GetHostName(), err.Error())
	}
	defer log.Printf("[%s]::command select executed in %d (ms)", common.GetHostName(), time.Since(sw).Milliseconds())
	var slice []MessageRow

	for dbdata.Next() {
		item := &MessageRow{}
		dbdata.Scan(&item.Id, &item.Message)
		slice = append(slice, *item)
	}

	_ = sqlcnn.Close()
	return slice, nil
}

func Add(message string) (int64, error) {
	sw := time.Now()
	sqlcnn, err := mySQLConnection.Open()
	if err != nil {
		log.Panic(err.Error())
	}
	result, err := sqlcnn.Exec("INSERT INTO Messages (Message) VALUES(?);", message)
	if err != nil {
		sqlcnn.Close()
		log.Panicf("[%s]::insert statement failed -- %s", common.GetHostName(), err.Error())
	}
	defer log.Printf("[%s]::command insert executed in %d (ms)", common.GetHostName(), time.Since(sw).Milliseconds())
	var affected int64 = 0
	var currentError *error = nil
	if result != nil {
		if affected, err = result.RowsAffected(); err == nil {
			log.Printf("[%s]::%d rows affected", common.GetHostName(), affected)
		} else {
			*currentError = err
			log.Printf("[%s]::no data found", common.GetHostName())
		}
	}
	sqlcnn.Close()
	return 0, *currentError
}

func Remove(id int32) (int64, error) {
	sw := time.Now()
	sqlcnn, err := mySQLConnection.Open()
	if err != nil {
		log.Panic(err.Error())
	}
	result, err := sqlcnn.Exec("DELETE FROM Messages WHERE ID = ?;", id)
	if err != nil {
		sqlcnn.Close()
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
	sqlcnn.Close()
	return 0, nil
}

func Find(id int32) (*string, error) {
	sw := time.Now()
	sqlcnn, err := mySQLConnection.Open()
	if err != nil {
		log.Panic(err.Error())
	}
	dbdata := sqlcnn.QueryRow("SELECT ID, Message FROM Messages WHERE ID = ?", id)
	defer log.Printf("[%s]::command select by id executed in %d (ms)", common.GetHostName(), time.Since(sw).Milliseconds())
	var messageValue *string = nil
	var currentError *error = nil
	if dbdata != nil {
		var id int
		if err := dbdata.Scan(&id, messageValue); err != nil {
			log.Println("ID ", id, " not found details:: ", err.Error())
			*currentError = errors.New(err.Error())
		}
	}
	sqlcnn.Close()
	return messageValue, *currentError
}
