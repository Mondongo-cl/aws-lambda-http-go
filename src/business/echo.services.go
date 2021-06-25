package business

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Mondongo-cl/http-rest-echo-go/dataaccess"
	"github.com/Mondongo-cl/http-rest-echo-go/datatypes"
)

func HandleEcho(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		failedGet := processGetMethod(r, w)
		if failedGet {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		processPostMethod(r, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
}

func processPostMethod(r *http.Request, w http.ResponseWriter) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var obj datatypes.EchoRequest
	err = json.Unmarshal(bytes, &obj)
	response := dataaccess.Add(obj.Message)

	bytesResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesResponse)
}

func processGetMethod(r *http.Request, w http.ResponseWriter) bool {
	path := r.URL.Path
	segments := strings.Split(path, "/")

	println("Segments", segments)
	if len(segments) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return true
	}
	selectedSegment := segments[2]
	println("segments ok!, selected value is :", selectedSegment)
	id, err := strconv.Atoi(selectedSegment)
	println("conversion result is : ", id)
	if err != nil {
		println(err)
		w.WriteHeader(http.StatusBadRequest)
		return true
	}

	response := datatypes.EchoResponse{
		Id:      int32(id),
		Message: dataaccess.Find(int32(id)),
	}

	data, e := json.Marshal(&response)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return true
	} else {
		println("Sending Response..	")
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	}
	return false
}
