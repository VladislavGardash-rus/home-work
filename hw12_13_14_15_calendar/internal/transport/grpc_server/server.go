package grpc_server

import (
	"context"
	"fmt"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/api/calendar_grpc"
	grpc_handlers "github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/handlers/grpc"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/logger"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/storage"
	"github.com/gardashvs/home-work/hw12_13_14_15_calendar/internal/transport/grpc_server/interceptor"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	eventDataHandler *grpc_handlers.EventDataHandler
	srv              *grpc.Server
	alias            string
	addr             string
}

func NewServer(addr string, storage storage.IStorage, alias string) *Server {
	server := &Server{
		alias:            alias,
		addr:             addr,
		srv:              grpc.NewServer(grpc.UnaryInterceptor(interceptor.Logging())),
		eventDataHandler: grpc_handlers.NewEventDataHandler(storage),
	}

	calendar_grpc.RegisterCalendarServer(server.srv, server.eventDataHandler)

	return server
}

func (s *Server) Start() error {
	logger.UseLogger().Info(fmt.Sprintf("grpc: Server %s started on ", s.alias), s.addr)

	lsn, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	return s.srv.Serve(lsn)
}

func (s *Server) Stop(_ context.Context) error {
	s.srv.GracefulStop()
	logger.UseLogger().Info("grpc: Server closed")
	return nil
}
