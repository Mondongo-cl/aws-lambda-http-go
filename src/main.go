package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/datatypes"
	"github.com/Mondongo-cl/http-rest-echo-go/settings"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/go-sql-driver/mysql"
)

func handleRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	p, e := json.Marshal(r)
	if e == nil {
		log.Printf("%v payload received ----> %v", time.Now(), string(p))
	} else {
		log.Printf("cannot marshall request data error %v", e)
	}

	switch r.HTTPMethod {
	case "GET":

		data, err := dataaccess.GetAll()
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 404, Body: err.Error(), Headers: map[string]string{"Content-Type": "text/plain"}}, nil
		}
		b, err := json.Marshal(data)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error(), Headers: map[string]string{"Content-Type": "text/plain"}}, nil
		}
		result := string(b)
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: result, Headers: map[string]string{"Content-Type": "application/json"}}, nil
	case "POST":
		data := datatypes.EchoRequest{}
		err := json.Unmarshal([]byte(r.Body), &data)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 404, Body: err.Error(), Headers: map[string]string{"Content-Type": "text/plain"}}, nil
		}
		dataaccess.Add(data.Message)
		b, err := json.Marshal(data)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error(), Headers: map[string]string{"Content-Type": "text/plain"}}, nil
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
