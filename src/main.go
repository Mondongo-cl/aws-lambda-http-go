package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/datatypes"
	"github.com/Mondongo-cl/http-rest-echo-go/settings"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

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
		result := string(b)
		log.Printf("payload:[%s]", result)
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: result, Headers: map[string]string{"Content-Type": "application/json"}}, nil
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
		result := string(b)
		log.Printf("payload:[%s]", result)
		return events.APIGatewayProxyResponse{StatusCode: 201, Body: result, Headers: map[string]string{"Content-Type": "application/json"}}, nil
	default:
		return events.APIGatewayProxyResponse{StatusCode: 504, Body: "", Headers: map[string]string{"Content-Type": "application/json"}}, nil
	}

}

func main() {
	s := settings.ConnectionSettings{}
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
}
