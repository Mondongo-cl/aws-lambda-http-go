package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/business"
	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/handlers"
)

func ResponseHealthyQuery(w http.ResponseWriter, r *http.Request) {
	version := dataaccess.GetMySqlVersion()
	if version != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("My SQL Version:%s\nTime:%v\nstatus:Healthy", *version, time.Now())))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("My SQL Version:\nTime:%v\nstatus:Unhealthy", time.Now())))

	}

}

func RegisterRoutes() {
	http.HandleFunc("/", ResponseHealthyQuery)
	http.HandleFunc("/stats", ResponseHealthyQuery)
	http.Handle("/echo", handlers.CorsHandler(business.HandleEcho))
	http.Handle("/echo/", handlers.CorsHandler(business.HandleEcho))
}
