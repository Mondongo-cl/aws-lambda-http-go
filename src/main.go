package main

import (
	"fmt"
	"net/http"

	"github.com/Mondongo-cl/http-rest-echo-go/middleware"
)

func start() {
	http.ListenAndServe(":5001", nil)
}

func main() {
	println("starting hello world service...")
	middleware.RegisterRoutes()
	go start()
	println("Press any key to close ...")
	_, _ = fmt.Scanln()
}
