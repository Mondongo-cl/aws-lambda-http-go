package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/business"
	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/handlers"
)

var (
	hostName string
)

func ResponseHealthyQuery(w http.ResponseWriter, r *http.Request) {
	version := dataaccess.GetMySqlVersion()
	if version != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hostname:%s\nMy SQL Version:%s\nTime:%v\nstatus:Healthy", hostName, *version, time.Now())))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Hostname:%s\nMy SQL Version:\nTime:%v\nstatus:Unhealthy", hostName, time.Now())))

	}

}

func RegisterRoutes(hostname string, DelayedHostname string) {
	hostName = hostname
	http.HandleFunc("/", ResponseHealthyQuery)
	http.HandleFunc("/stats", ResponseHealthyQuery)
	http.Handle("/echo", handlers.CorsHandler(business.HandleEcho, DelayedHostname))
	http.Handle("/echo/", handlers.CorsHandler(business.HandleEcho, DelayedHostname))
}
