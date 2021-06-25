package business

import (
	"encoding/json"
	"net/http"

	"github.com/Mondongo-cl/http-rest-echo-go/datatypes"
)

func HandleEcho(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		response := datatypes.EchoRequest{
			Message: "Echo Service, this is a hard coded message!",
		}
		data, e := json.Marshal(&response)
		if e != nil {
			println(e)
			w.Write([]byte("Error!!"))
			w.Write([]byte(e.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			println("Sending Response..	")
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)

		}
	}
}
