package handlers

import (
	"net/http"
)

/*
Add cors support to "*" as allowed origin:
*/
func CorsHandler(handler http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		rw.Header().Add("Access-Control-Allow-Headers",
			"Accept, Content-Lenght, X-CSRF-Token, Accept-Encoding, Content-Type, Authorization, API-KEY")
		handler.ServeHTTP(rw, r)
	})
}
