package middleware

import (
	"net/http"

	"github.com/Mondongo-cl/http-rest-echo-go/business"
	"github.com/Mondongo-cl/http-rest-echo-go/handlers"
)

func ResponseHealthyQuery(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Healthy"))
}

func RegisterRoutes() {
	http.HandleFunc("/", ResponseHealthyQuery)
	http.HandleFunc("/stats", ResponseHealthyQuery)
	http.Handle("/echo", handlers.CorsHandler(business.HandleEcho))
	http.Handle("/echo/", handlers.CorsHandler(business.HandleEcho))
}
