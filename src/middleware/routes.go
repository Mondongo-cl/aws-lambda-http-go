package middleware

import (
	"net/http"

	"github.com/Mondongo-cl/http-rest-echo-go/business"
	"github.com/Mondongo-cl/http-rest-echo-go/handlers"
)

func RegisterRoutes() {

	// TODO: ADD Healthy check routes
	// get:/
	// get:/stat

	http.Handle("/echo", handlers.CorsHandler(business.HandleEcho))
	http.Handle("/echo/", handlers.CorsHandler(business.HandleEcho))
}
