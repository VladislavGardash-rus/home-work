package http_handlers

import (
	"net/http"
)

type HelloWorldHandler struct{}

func NewHelloWorldHandler() *HelloWorldHandler {
	return new(HelloWorldHandler)
}

func (h *HelloWorldHandler) GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
