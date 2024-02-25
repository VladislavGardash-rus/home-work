package http_handlers

import (
	"net/http"
)

type HelloWorldHandler struct{}

func NewHelloWorldHandler() *HelloWorldHandler {
	return new(HelloWorldHandler)
}

func (h *HelloWorldHandler) GetHelloWorld(r *http.Request) (interface{}, error) {
	return "Hello world", nil
}
