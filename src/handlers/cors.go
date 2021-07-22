package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/common"
	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
)

/*
Add cors support to "*" as allowed origin:
*/
func CorsHandler(handler http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if dataaccess.IsDelayedHost(true) {
			log.Printf("[%s]::server is really slow", common.GetHostName())
			time.Sleep(time.Second * 1)
		}
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Header().Add("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		rw.Header().Add("Access-Control-Allow-Headers",
			"Accept, Content-Lenght, X-CSRF-Token, Accept-Encoding, Content-Type, Authorization, API-KEY")
		handler.ServeHTTP(rw, r)
	})
}
