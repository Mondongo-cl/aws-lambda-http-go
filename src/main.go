package main

import (
	"encoding/json"
	"net/http"
)

type echoHandler struct {
	Message string
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		response := EchoRequest{
			Message:  "Echo Service, this is a hard coded message!",
			byteSize: 0,
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

func (handler *echoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(handler.Message))
}

func main() {
	println("starting hello world service...")
	http.Handle("/helloworld", &echoHandler{Message: "Hola Mundo"})
	http.HandleFunc("/echo", handleEcho)
	http.ListenAndServe(":5001", nil)
}
