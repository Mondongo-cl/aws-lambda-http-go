package main

import (
	"net/http"

	"github.com/Mondongo-cl/http-rest-echo-go/middleware"
)

func main() {
	println("starting hello world service...")
	middleware.RegisterRoutes()
	http.ListenAndServe(":5001", nil)
}
