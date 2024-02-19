package http_server

import (
	"context"
	"fmt"
	http_handlers "github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/handlers/http"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	helloWorldHandler *http_handlers.HelloWorldHandler
	eventDataHandler  *http_handlers.EventDataHandler
	srv               *http.Server
	router            *mux.Router
	alias             string
}

func NewServer(addr string, storage storage.IStorage, alias string) *Server {
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	server := &Server{
		alias:             alias,
		srv:               srv,
		helloWorldHandler: http_handlers.NewHelloWorldHandler(),
		eventDataHandler:  http_handlers.NewEventDataHandler(storage),
	}
	server.initRoting()

	return server
}

func (s *Server) initRoting() {
	s.router = mux.NewRouter()
	s.router.Use(loggingMiddleware)
	s.router.HandleFunc("/hello-world", s.helloWorldHandler.GetHelloWorld).Methods(http.MethodGet)

	s.router.HandleFunc("/event", s.eventDataHandler.GetEvents).Methods(http.MethodGet)
	s.router.HandleFunc("/event/create", s.eventDataHandler.PostCreateEvent).Methods(http.MethodPost)
	s.router.HandleFunc("/event/update", s.eventDataHandler.PostUpdateEvent).Methods(http.MethodPost)
	s.router.HandleFunc("/event/delete", s.eventDataHandler.PostDeleteEvent).Methods(http.MethodPost)

	s.srv.Handler = s.router
}

func (s *Server) Start() error {
	logger.UseLogger().Info(fmt.Sprintf("http: Server %s started on", s.alias), s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
