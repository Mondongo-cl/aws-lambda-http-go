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
		processGetMethod(r, w)
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := dataaccess.Add(obj.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytesResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesResponse)
}

func processGetMethod(r *http.Request, w http.ResponseWriter) {
	path := r.URL.Path
	segments := strings.Split(path, "/")

	println("Segments", segments)
	if len(segments) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	selectedSegment := segments[2]
	println("segments ok!, selected value is :", selectedSegment)
	id, err := strconv.Atoi(selectedSegment)
	println("conversion result is : ", id)
	if err != nil {
		println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg, err := dataaccess.Find(int32(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	response := datatypes.EchoResponse{
		Id:      int32(id),
		Message: *msg,
	}

	data, e := json.Marshal(&response)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		println("Sending Response..	")
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
