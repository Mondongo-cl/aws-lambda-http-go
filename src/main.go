package main

import (
	"context"
	_ "fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/go-sql-driver/mysql"
)

func handleRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("processing vent id: %s", r.RequestContext.RequestID)

	switch r.HTTPMethod {
	case http.MethodGet:
		return events.APIGatewayProxyResponse{StatusCode: 200}, nil
	case http.MethodPost:
		return events.APIGatewayProxyResponse{StatusCode: 201}, nil
	default:
		return events.APIGatewayProxyResponse{StatusCode: 504}, nil

	}
}

func main() {
	println("starting hello world service...")
	lambda.Start(handleRequest)

	// middleware.RegisterRoutes()
	// start()
}
