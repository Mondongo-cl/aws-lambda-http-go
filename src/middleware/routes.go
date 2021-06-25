package middleware

import (
	"net/http"

	"github.com/Mondongo-cl/http-rest-echo-go/business"
	"github.com/Mondongo-cl/http-rest-echo-go/handlers"
)

func RegisterRoutes() {
	http.Handle("/helloworld", &handlers.EchoHandler{Message: "Hello World!"})
	http.HandleFunc("/echo", business.HandleEcho)
}
