package main

import (
	"context"
	_ "fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/go-sql-driver/mysql"
)

func handleRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("processing BODY: %s verb: %s path: %s", r.Body, r.HTTPMethod, r.Path)
	defer log.Printf("event id: %s has benn processed", r.RequestContext.RequestID)
	switch r.HTTPMethod {
	case "GET":
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: "{\"Method\":\"GET\"}"}, nil
	case "POST":
		return events.APIGatewayProxyResponse{StatusCode: 201, Body: "{\"Method\":\"POST\"}"}, nil
	default:
		return events.APIGatewayProxyResponse{StatusCode: 504, Body: "{\"Method\":\"OTHER\"}"}, nil
	}
}

func main() {
	log.Println("starting hello world service...")
	lambda.Start(handleRequest)
	log.Println("stoping hello worls service...")
	// middleware.RegisterRoutes()
	// start()
}
