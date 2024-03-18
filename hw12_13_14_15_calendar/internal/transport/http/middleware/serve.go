package middleware

import (
	"context"
	"encoding/json"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"net/http"
	"time"
)

type errorMessage struct {
	Method string `json:"method"`
	Api    string `json:"api"`
	Error  string `json:"error"`
}

type HandlerFunc func(ctx context.Context, r *http.Request) (interface{}, error)

func Serve(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		response := new(http.Response)

		ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)

		data, err := h(ctx, r)
		if err != nil {
			logger.UseLogger().Error(err)

			response.StatusCode = http.StatusBadRequest
			r.Response = response

			errorMessage := new(errorMessage)
			errorMessage.Method = r.Method
			errorMessage.Api = r.RequestURI
			errorMessage.Error = err.Error()
			b, _ := json.Marshal(errorMessage)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)

			return
		}

		response.StatusCode = http.StatusOK
		r.Response = response
		b, _ := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
