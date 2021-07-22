package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mondongo-cl/http-rest-echo-go/business"
	"github.com/Mondongo-cl/http-rest-echo-go/common"
	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/handlers"
)

func ResponseHealthyQuery(w http.ResponseWriter, r *http.Request) {
	version := dataaccess.GetMySqlVersion()
	if version != nil {
		w.WriteHeader(http.StatusOK)
		if dataaccess.IsDelayedHost(false) {
			time.Sleep(time.Second * 1)
			w.Write([]byte(fmt.Sprintf("Hostname:%s\nMy SQL Version:%s\nTime:%v\nstatus:Delayed", common.GetHostName(), *version, time.Now())))
		} else {
			w.Write([]byte(fmt.Sprintf("Hostname:%s\nMy SQL Version:%s\nTime:%v\nstatus:Healthy", common.GetHostName(), *version, time.Now())))
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Hostname:%s\nMy SQL Version:\nTime:%v\nstatus:Unhealthy", common.GetHostName(), time.Now())))
	}

}

func RegisterRoutes(hostname string) {
	http.HandleFunc("/", ResponseHealthyQuery)
	http.HandleFunc("/stats", ResponseHealthyQuery)
	http.Handle("/echo", handlers.CorsHandler(business.HandleEcho))
	http.Handle("/echo/", handlers.CorsHandler(business.HandleEcho))
}
