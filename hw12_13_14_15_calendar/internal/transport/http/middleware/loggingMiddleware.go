package http_server

import (
	"fmt"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"net/http"
	"time"
)

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		handler.ServeHTTP(w, r)

		logger.UseLogger().Info(
			fmt.Sprintf("%s %s %s %s %s %d %s",
				r.RemoteAddr,
				startTime.Format(time.DateTime),
				r.Method,
				r.RequestURI,
				w.Header().Get("status"),
				time.Since(startTime).Milliseconds(),
				r.Header["User-Agent"],
			))
	})
}
