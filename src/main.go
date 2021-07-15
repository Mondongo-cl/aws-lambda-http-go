package main

import (
	"context"
	"encoding/json"
	_ "fmt"
	"log"
	"os"
	"strconv"

	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/datatypes"
	"github.com/Mondongo-cl/http-rest-echo-go/settings"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	_ "github.com/go-sql-driver/mysql"
)

func handleRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	defer log.Printf("event id: %s has benn processed", r.RequestContext.RequestID)
	switch r.HTTPMethod {
	case "GET":
		data, err := dataaccess.GetAll()
		if err != nil {
			panic(err.Error())
		}
		b, err := json.Marshal(data)
		if err != nil {
			panic("Serialization Error:: " + err.Error())
		}

		log.Printf("payload:[%s]", string(b))
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(b)}, nil
	case "POST":
		data := datatypes.EchoRequest{}
		err := json.Unmarshal([]byte(r.Body), &data)
		if err != nil {
			panic("BAD REQUEST")
		}
		dataaccess.Add(data.Message)
		b, err := json.Marshal(data)
		if err != nil {
			panic("Serialization Error:: " + err.Error())
		}
		log.Printf("payload:[%s]", string(b))
		return events.APIGatewayProxyResponse{StatusCode: 201, Body: string(b)}, nil
	default:
		return events.APIGatewayProxyResponse{StatusCode: 504, Body: ""}, nil
	}
}

func main() {

	sess := session.Must(session.NewSession())
	sess.Handlers.Send.PushFront(func(r *request.Request) {
		// Log every request made and its payload
		log.Printf("Request: %s/%s, Params: %s", r.ClientInfo.ServiceName, r.Operation.Name, r.Params)
	})
	s := settings.ConnectionSettings{}
	/*
		host := flag.String("servername", "localhost", "Database Server Name")
		port := flag.Int("serverport", 3306, "Database TCP port")
		username := flag.String("username", "root", "Database User Name")
		password := flag.String("password", "", "Database User Name Password")
		dbname := flag.String("database", "sample", "Database Name")
		flag.Parse()
	*/
	host := os.Getenv("servername")
	port, _ := strconv.Atoi(os.Getenv("serverport"))
	username := os.Getenv("username")
	password := os.Getenv("password")
	dbname := os.Getenv("database")

	s.Host = host
	s.Port = port
	s.Username = username
	s.Password = password
	s.Database = dbname
	dataaccess.Configure(s)

	log.Println("starting hello world service...")
	lambda.Start(handleRequest)
	log.Println("stoping hello worls service...")
	// middleware.RegisterRoutes()
	// start()
}
