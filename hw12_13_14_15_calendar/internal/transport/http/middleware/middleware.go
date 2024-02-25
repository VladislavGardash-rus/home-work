package http_server

import (
	"encoding/json"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"net/http"
)

type errorMessage struct {
	Method string `json:"method"`
	Api    string `json:"api"`
	Error  string `json:"error"`
}

type HandlerFunc func(r *http.Request) (interface{}, error)

func Serve(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		data, err := h(r)
		if err != nil {
			logger.UseLogger().Error(err)

			errorMessage := new(errorMessage)
			errorMessage.Method = r.Method
			errorMessage.Api = r.RequestURI
			errorMessage.Error = err.Error()
			b, _ := json.Marshal(errorMessage)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)
			return
		}

		b, err := json.Marshal(data)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
