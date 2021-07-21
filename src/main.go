package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/middleware"
	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/go-sql-driver/mysql"
)

var (
	HostName string
)

func start(port *int) {
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func initAwsSession() *session.Session {
	sess, err := session.NewSession()
	if err != nil {
		log.Println(err.Error())
	}

	cfg := sess.Config
	if cfg != nil && (cfg.Endpoint != nil && cfg.Region != nil) {
		log.Printf("[%s]::starting in Amazon AWS\nEndpoint:[%s]\nRegion[%s]", HostName, *cfg.Endpoint, *cfg.Region)
	}
	return sess
}

func main() {
	currentHostname, err := os.Hostname()
	_ = initAwsSession()

	if err != nil {
		log.Printf("a error occurred while triying to get the underneath os hostname, the error is %s", err.Error())
		HostName = "<<None>>"

	} else {
		HostName = currentHostname
		log.Printf("the hostname was get successfully the current hostname is %s", HostName)
	}

	username := flag.String("dbusername", "root", "database username")
	password := flag.String("dbpassword", "123456", "database password")
	hostname := flag.String("dbhostname", "localhost", "database hostname")
	port := flag.Int("dbport", 3306, "database port number")
	publicPort := flag.Int("httplistenerport", 0, "Http Listener port")
	databasename := flag.String("databasename", "default", "database name")
	flag.Parse()

	if !flag.Parsed() || (*username == "" || *password == "" || *hostname == "" || *port == 0 || *databasename == "" || *publicPort == 0) {
		fmt.Printf("Incorrect Parameters:\n============================\nUSAGE:\n================= ")
		flag.PrintDefaults()
		return
	}
	dataaccess.Configure(username, password, hostname, port, databasename)
	println("starting hello world service...")
	middleware.RegisterRoutes()
	start(publicPort)
}
