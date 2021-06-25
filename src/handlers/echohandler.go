package handlers

import (
	"fmt"
	"net/http"
)

type EchoHandler struct {
	Message string
}

func (handler *EchoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler pre-write")
	writer.Write([]byte(handler.Message))
	fmt.Println("handler post-write")
}
