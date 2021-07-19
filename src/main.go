package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func start(port *int) {
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func main() {

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
