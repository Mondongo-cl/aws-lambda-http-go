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
	if len(segments) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(segments) == 3 {

		selectedSegment := segments[2]
		id, err := strconv.Atoi(selectedSegment)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		msg, err := dataaccess.Find(int32(id))
		if err == nil {
			response, err := CreateResponseItem(id, msg)
			if err != nil {
				_ = WriteResponseItem(response, w)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	} else {
		rowList, err := dataaccess.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return

		} else {
			response := CreateResponseItemList(rowList)
			_ = WriteReponseItemList(response, w)
		}
	}
}

func WriteResponseItem(response datatypes.EchoResponse, w http.ResponseWriter) bool {
	data, e := json.Marshal(&response)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return true
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
	return false
}

func WriteReponseItemList(response *[]*datatypes.EchoResponse, w http.ResponseWriter) bool {
	data, e := json.Marshal(response)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return true
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
	return false
}

func CreateResponseItemList(rowList []dataaccess.MessageRow) *[]*datatypes.EchoResponse {
	var response []*datatypes.EchoResponse
	for _, v := range rowList {
		i := datatypes.EchoResponse{
			Message: v.Message,
			Id:      v.Id,
		}
		response = append(response, &i)
	}
	return &response
}

func CreateResponseItem(id int, msg *string) (datatypes.EchoResponse, error) {
	response := datatypes.EchoResponse{
		Id:      int32(id),
		Message: *msg,
	}
	return response, nil
}
