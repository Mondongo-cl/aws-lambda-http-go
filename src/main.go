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

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("processing vent id: %s", request.RequestContext.RequestID)

	return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil

}

func start() {
	http.ListenAndServe(":5001", nil)
}

func main() {
	println("starting hello world service...")
	lambda.Start(handleRequest)

	// middleware.RegisterRoutes()
	// start()
}
