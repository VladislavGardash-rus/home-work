package http_handlers

import (
	"context"
	"net/http"
)

type HelloWorldHandler struct{}

func NewHelloWorldHandler() *HelloWorldHandler {
	return new(HelloWorldHandler)
}

func (h *HelloWorldHandler) GetHelloWorld(ctx context.Context, r *http.Request) (interface{}, error) {
	return "Hello world", nil
}
